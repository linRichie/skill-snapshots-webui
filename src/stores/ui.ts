import { defineStore } from 'pinia'
import { ref } from 'vue'

/**
 * UI 状态存储
 */
export const useUiStore = defineStore('ui', () => {
  // 主题
  const theme = ref<'light' | 'dark'>('light')
  const sidebarCollapsed = ref(false)
  const mobileMenuOpen = ref(false)

  // 初始化主题（从 localStorage 读取）
  function initTheme() {
    const saved = localStorage.getItem('theme') as 'light' | 'dark' | null
    if (saved) {
      theme.value = saved
    } else {
      // 根据系统偏好设置
      const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches
      theme.value = prefersDark ? 'dark' : 'light'
    }
    applyTheme()
  }

  // 切换主题
  function toggleTheme() {
    theme.value = theme.value === 'light' ? 'dark' : 'light'
    localStorage.setItem('theme', theme.value)
    applyTheme()
  }

  // 应用主题
  function applyTheme() {
    if (theme.value === 'dark') {
      document.documentElement.classList.add('dark')
    } else {
      document.documentElement.classList.remove('dark')
    }
  }

  // 切换侧边栏
  function toggleSidebar() {
    sidebarCollapsed.value = !sidebarCollapsed.value
  }

  // 切换移动端菜单
  function toggleMobileMenu() {
    mobileMenuOpen.value = !mobileMenuOpen.value
  }

  // 关闭移动端菜单
  function closeMobileMenu() {
    mobileMenuOpen.value = false
  }

  return {
    theme,
    sidebarCollapsed,
    mobileMenuOpen,
    initTheme,
    toggleTheme,
    toggleSidebar,
    toggleMobileMenu,
    closeMobileMenu
  }
})
