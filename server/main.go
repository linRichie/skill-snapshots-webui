// Skill Snapshots WebUI - Backend API Server
// A RESTful API for managing Claude Code skill snapshots
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

// Config represents the application configuration
type Config struct {
	SkillsDir  string            `yaml:"skills_dir"`
	Categories map[string]string `yaml:"categories"`
	Descriptions map[string]string `yaml:"descriptions"`
	Server    ServerConfig      `yaml:"server"`
}

// ServerConfig represents server-specific configuration
type ServerConfig struct {
	Port string `yaml:"port"`
	Mode string `yaml:"mode"`
}

// Skill represents a Claude Code skill with version information
type Skill struct {
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Path          string    `json:"path"`
	Versions      []Version `json:"versions,omitempty"`
	LatestVersion string    `json:"latestVersion,omitempty"`
	Category      string    `json:"category,omitempty"`
	Tags          []string  `json:"tags,omitempty"`
	CreatedAt     string    `json:"createdAt,omitempty"`
	UpdatedAt     string    `json:"updatedAt,omitempty"`
	Content       string    `json:"content,omitempty"`
	Dependencies  []string  `json:"dependencies,omitempty"`
	Scripts       []Script  `json:"scripts,omitempty"`
}

// Version represents a version snapshot of a skill
type Version struct {
	Tag       string `json:"tag"`
	SkillName string `json:"skillName"`
	Version   string `json:"version"`
	Message   string `json:"message,omitempty"`
	Date      string `json:"date,omitempty"`
	Author    string `json:"author,omitempty"`
}

// Script represents a script file within a skill
type Script struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Description string `json:"description,omitempty"`
}

// ListResponse represents a paginated list response
type ListResponse struct {
	Items      []Skill    `json:"items"`
	Pagination Pagination `json:"pagination"`
}

// Pagination represents pagination information
type Pagination struct {
	Page       int `json:"page"`
	PageSize   int `json:"pageSize"`
	Total      int `json:"total"`
	TotalPages int `json:"totalPages"`
}

// APIResponse represents a standard API response
type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Global variables
var (
	config Config
	skills []Skill
)

func init() {
	// Load configuration
	if err := loadConfig(); err != nil {
		log.Printf("加载配置文件失败: %v，使用默认配置", err)
		// Use default config
		config = Config{
			SkillsDir: "../skill-snapshots",
			Categories: make(map[string]string),
			Descriptions: make(map[string]string),
			Server: ServerConfig{
				Port: "8000",
				Mode: "debug",
			},
		}
	}

	// Resolve skills directory to absolute path
	skillsDir := expandPath(config.SkillsDir)
	if !filepath.IsAbs(skillsDir) {
		// Relative to current directory
		abs, err := filepath.Abs(skillsDir)
		if err != nil {
			log.Fatal("无法解析技能目录路径:", err)
		}
		skillsDir = abs
	}
	config.SkillsDir = skillsDir

	// Load skills data
	loadSkills()
}

// expandPath 展开路径中的环境变量和波浪号
// 支持: ~, ${HOME}, $HOME, ${VAR}
func expandPath(path string) string {
	// 展开波浪号 ~
	if strings.HasPrefix(path, "~/") || path == "~" {
		homeDir, err := os.UserHomeDir()
		if err == nil {
			if path == "~" {
				return homeDir
			}
			path = filepath.Join(homeDir, path[2:])
		}
	}

	// 展开环境变量 ${VAR} 和 $VAR
	path = os.ExpandEnv(path)

	return path
}

// loadConfig loads configuration from config.yaml
func loadConfig() error {
	configPath := "config.yaml"

	// Check if config.yaml exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return fmt.Errorf("配置文件不存在: %s", configPath)
	}

	// Read config file
	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}

	// Parse YAML
	if err := yaml.Unmarshal(data, &config); err != nil {
		return fmt.Errorf("解析配置文件失败: %w", err)
	}

	// Validate required fields
	if config.SkillsDir == "" {
		return fmt.Errorf("配置文件中缺少 skills_dir 字段")
	}

	log.Printf("配置加载成功: skills_dir=%s", config.SkillsDir)
	return nil
}

// loadSkills loads skill data from the configured directory
func loadSkills() {
	// Check if skills directory exists
	if _, err := os.Stat(config.SkillsDir); os.IsNotExist(err) {
		log.Printf("技能目录不存在: %s", config.SkillsDir)
		return
	}

	entries, err := os.ReadDir(config.SkillsDir)
	if err != nil {
		log.Printf("读取目录失败: %v", err)
		return
	}

	skills = []Skill{}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		skillName := entry.Name()

		// Skip system directories
		if strings.HasPrefix(skillName, ".") || skillName == "node_modules" ||
		   skillName == "dist" || skillName == "build" || skillName == "webui" {
			continue
		}

		skillPath := filepath.Join(config.SkillsDir, skillName)
		skillFile := filepath.Join(skillPath, "SKILL.md")

		// Read SKILL.md content
		var content string
		if data, err := os.ReadFile(skillFile); err == nil {
			content = string(data)
		}

		// Get description from config
		description := config.Descriptions[skillName]
		if description == "" {
			description = "暂无描述"
		}

		// Get category from config
		category := config.Categories[skillName]

		// Parse version information from git tags
		versions := parseVersions(skillName)

		// Get latest version
		latestVersion := ""
		if len(versions) > 0 {
			latestVersion = versions[0].Version
		}

		// Get file modification time
		var updatedAt string
		if info, err := entry.Info(); err == nil {
			updatedAt = info.ModTime().Format(time.RFC3339)
		}

		skill := Skill{
			Name:          skillName,
			Description:   description,
			Path:          skillPath,
			Versions:      versions,
			LatestVersion: latestVersion,
			Category:      category,
			UpdatedAt:     updatedAt,
			Content:       content,
		}

		skills = append(skills, skill)
	}

	// Sort by update time (newest first)
	sort.Slice(skills, func(i, j int) bool {
		return skills[i].UpdatedAt > skills[j].UpdatedAt
	})

	log.Printf("已加载 %d 个技能", len(skills))
}

// parseVersions parses version information from git tags
func parseVersions(skillName string) []Version {
	versions := []Version{}

	// Get git repository root (search upward for .git directory)
	gitRoot := findGitRoot(config.SkillsDir)
	if gitRoot == "" {
		// Not in a git repository
		return versions
	}

	// Run git tag command to get all tags for this skill
	cmd := exec.Command("git", "tag", "-l", "--sort=-v:refname", fmt.Sprintf("%s/v*", skillName))
	cmd.Dir = gitRoot
	output, err := cmd.Output()
	if err != nil {
		// No tags found for this skill
		return versions
	}

	tags := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(tags) == 0 || (len(tags) == 1 && tags[0] == "") {
		return versions
	}

	// Get detailed information for each tag
	for _, tag := range tags {
		if tag == "" {
			continue
		}

		// Extract version number (e.g., "baidu-hot/v7" -> "v7")
		version := tag
		if idx := strings.LastIndex(tag, "/"); idx != -1 {
			version = tag[idx+1:]
		}

		// Get tag date
		tagDate := getTagDate(gitRoot, tag)

		// Get tag message
		tagMessage := getTagMessage(gitRoot, tag)

		versions = append(versions, Version{
			Tag:       tag,
			SkillName: skillName,
			Version:   version,
			Message:   tagMessage,
			Date:      tagDate,
		})
	}

	return versions
}

// findGitRoot searches upward for the .git directory
func findGitRoot(startPath string) string {
	path := startPath
	for {
		gitDir := filepath.Join(path, ".git")
		if _, err := os.Stat(gitDir); err == nil {
			return path
		}

		parent := filepath.Dir(path)
		if parent == path {
			// Reached root without finding .git
			return ""
		}
		path = parent
	}
}

// extractVersionNumber extracts the numeric version from a tag
func extractVersionNumber(tag string) int {
	idx := strings.LastIndex(tag, "/v")
	if idx == -1 {
		idx = strings.Index(tag, "/v")
	}
	if idx == -1 {
		return 0
	}

	versionStr := tag[idx+2:]
	num, err := strconv.Atoi(versionStr)
	if err != nil {
		return 0
	}
	return num
}

// getTagDate gets the creation date of a git tag
func getTagDate(repoDir, tag string) string {
	cmd := exec.Command("git", "log", "-1", "--format=%ci", tag)
	cmd.Dir = repoDir
	output, err := cmd.Output()
	if err != nil {
		return time.Now().Format(time.RFC3339)
	}

	dateStr := strings.TrimSpace(string(output))
	if dateStr == "" {
		return time.Now().Format(time.RFC3339)
	}

	t, err := time.Parse("2006-01-02 15:04:05 -0700", dateStr)
	if err != nil {
		return time.Now().Format(time.RFC3339)
	}
	return t.Format(time.RFC3339)
}

// getTagMessage gets the commit message for a git tag
func getTagMessage(repoDir, tag string) string {
	cmd := exec.Command("git", "log", "-1", "--format=%s", tag)
	cmd.Dir = repoDir
	output, err := cmd.Output()
	if err != nil {
		return "版本更新"
	}

	msg := strings.TrimSpace(string(output))
	if msg == "" {
		return "版本更新"
	}
	return msg
}

// setupRouter configures and returns the Gin router
func setupRouter() *gin.Engine {
	// Set Gin mode
	gin.SetMode(config.Server.Mode)

	router := gin.Default()

	// Configure CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	router.Use(cors.New(corsConfig))

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"time":   time.Now().Format(time.RFC3339),
			"config": gin.H{
				"skills_dir": config.SkillsDir,
			},
		})
	})

	// API routes
	api := router.Group("/api")
	{
		// Skill endpoints
		api.GET("/skills", getSkills)
		api.GET("/skills/:name", getSkillDetail)
		api.GET("/skills/:name/versions", getSkillVersions)
		api.GET("/versions", getAllVersions)
		api.GET("/categories", getCategories)
		api.GET("/config", getConfig) // 新增：获取当前配置

		// System endpoints
		api.GET("/system/info", getSystemInfo)
		api.GET("/system/stats", getSystemStats)
	}

	// Serve static files (frontend build output)
	router.Static("/assets", "./dist/assets")
	router.NoRoute(func(c *gin.Context) {
		c.File("./dist/index.html")
	})

	return router
}

// getSkills returns a paginated list of skills
func getSkills(c *gin.Context) {
	page := 1
	pageSize := 20

	if p := c.Query("page"); p != "" {
		fmt.Sscanf(p, "%d", &page)
	}
	if ps := c.Query("pageSize"); ps != "" {
		fmt.Sscanf(ps, "%d", &pageSize)
	}

	search := c.Query("search")
	category := c.Query("category")

	// Filter skills
	filtered := []Skill{}
	for _, skill := range skills {
		if search != "" && !strings.Contains(strings.ToLower(skill.Name), strings.ToLower(search)) &&
			!strings.Contains(strings.ToLower(skill.Description), strings.ToLower(search)) {
			continue
		}
		if category != "" && skill.Category != category {
			continue
		}
		filtered = append(filtered, skill)
	}

	// Paginate
	total := len(filtered)
	totalPages := (total + pageSize - 1) / pageSize
	start := (page - 1) * pageSize
	end := start + pageSize

	if start >= total {
		start = total
	}
	if end > total {
		end = total
	}

	c.JSON(http.StatusOK, ListResponse{
		Items: filtered[start:end],
		Pagination: Pagination{
			Page:       page,
			PageSize:   pageSize,
			Total:      total,
			TotalPages: totalPages,
		},
	})
}

// getSkillDetail returns detailed information about a specific skill
func getSkillDetail(c *gin.Context) {
	name := c.Param("name")

	for _, skill := range skills {
		if skill.Name == name {
			c.JSON(http.StatusOK, skill)
			return
		}
	}

	c.JSON(http.StatusNotFound, APIResponse{
		Success: false,
		Error:   "技能不存在",
	})
}

// getSkillVersions returns all versions for a specific skill
func getSkillVersions(c *gin.Context) {
	name := c.Param("name")

	for _, skill := range skills {
		if skill.Name == name {
			c.JSON(http.StatusOK, skill.Versions)
			return
		}
	}

	c.JSON(http.StatusNotFound, APIResponse{
		Success: false,
		Error:   "技能不存在",
	})
}

// getAllVersions returns all versions across all skills
func getAllVersions(c *gin.Context) {
	allVersions := []Version{}
	for _, skill := range skills {
		allVersions = append(allVersions, skill.Versions...)
	}

	// Sort by date
	sort.Slice(allVersions, func(i, j int) bool {
		return allVersions[i].Date > allVersions[j].Date
	})

	c.JSON(http.StatusOK, allVersions)
}

// getCategories returns a list of all skill categories
func getCategories(c *gin.Context) {
	categories := map[string]bool{}
	for _, skill := range skills {
		if skill.Category != "" {
			categories[skill.Category] = true
		}
	}

	result := make([]string, 0, len(categories))
	for cat := range categories {
		result = append(result, cat)
	}
	sort.Strings(result)

	c.JSON(http.StatusOK, result)
}

// getConfig returns the current configuration
func getConfig(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"skills_dir": config.SkillsDir,
		"server": gin.H{
			"port": config.Server.Port,
			"mode": config.Server.Mode,
		},
	})
}

// getSystemInfo returns system information
func getSystemInfo(c *gin.Context) {
	hostname, _ := os.Hostname()
	c.JSON(http.StatusOK, gin.H{
		"version":   "1.0.0",
		"hostname":  hostname,
		"startTime": time.Now().Add(-time.Hour).Format(time.RFC3339),
		"goVersion": "1.21",
	})
}

// getSystemStats returns system statistics
func getSystemStats(c *gin.Context) {
	totalVersions := 0
	byCategory := map[string]int{}

	for _, skill := range skills {
		totalVersions += len(skill.Versions)
		if skill.Category != "" {
			byCategory[skill.Category]++
		}
	}

	// Get last update time
	var lastUpdate string
	if len(skills) > 0 && skills[0].UpdatedAt != "" {
		lastUpdate = skills[0].UpdatedAt
	}

	c.JSON(http.StatusOK, gin.H{
		"totalSkills":   len(skills),
		"totalVersions": totalVersions,
		"byCategory":    byCategory,
		"lastUpdate":    lastUpdate,
	})
}

func main() {
	// Check if config.yaml exists, if not, print help message
	if _, err := os.Stat("config.yaml"); os.IsNotExist(err) {
		fmt.Println("╔═══════════════════════════════════════════════════════════════╗")
		fmt.Println("║           Skill Snapshots WebUI - 首次运行配置向导              ║")
		fmt.Println("╠═══════════════════════════════════════════════════════════════╣")
		fmt.Println("║")
		fmt.Println("║  欢迎使用 Skill Snapshots WebUI！")
		fmt.Println("║")
		fmt.Println("║  第一步：创建配置文件")
		fmt.Println("║")
		fmt.Println("║  请复制 config.example.yaml 并重命名为 config.yaml：")
		fmt.Println("║")
		fmt.Println("║    cp config.example.yaml config.yaml")
		fmt.Println("║")
		fmt.Println("║  第二步：编辑配置文件")
		fmt.Println("║")
		fmt.Println("║  修改 config.yaml 中的 skills_dir 参数，指向你的技能目录：")
		fmt.Println("║")
		fmt.Println("║    skills_dir: \"/path/to/your/skills\"")
		fmt.Println("║")
		fmt.Println("║  第三步：启动服务")
		fmt.Println("║")
		fmt.Println("║    go run main.go")
		fmt.Println("║")
		fmt.Println("║  配置文件说明：")
		fmt.Println("║")
		fmt.Println("║  - skills_dir: 你的技能仓库根目录（必需）")
		fmt.Println("║  - categories: 技能分类映射（可选）")
		fmt.Println("║  - descriptions: 技能描述（可选）")
		fmt.Println("║  - server.port: API 服务端口（默认：8000）")
		fmt.Println("║  - server.mode: Gin 模式（debug/release，默认：release）")
		fmt.Println("║")
		fmt.Println("╚═══════════════════════════════════════════════════════════════╝")
		fmt.Println()
		os.Exit(0)
	}

	router := setupRouter()

	port := config.Server.Port
	if port == "" {
		port = "8000"
	}

	log.Printf("服务器启动在 http://localhost:%s", port)
	log.Printf("技能目录: %s", config.SkillsDir)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}
