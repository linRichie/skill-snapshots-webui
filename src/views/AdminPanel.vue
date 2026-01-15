<template>
  <div class="min-h-screen">
    <!-- 页面头部 -->
    <div class="bg-gradient-to-r from-primary-600 to-primary-800 dark:from-primary-800 dark:to-primary-900">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div class="flex items-center justify-between">
          <div>
            <h1 class="text-3xl font-bold text-white">管理后台</h1>
            <p class="mt-1 text-primary-100">技能快照管理与配置</p>
          </div>
          <div class="flex items-center space-x-3">
            <button @click="refreshData" class="px-4 py-2 bg-white/20 hover:bg-white/30 text-white rounded-lg transition-colors flex items-center">
              <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
              </svg>
              刷新数据
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- 统计卡片 -->
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 -mt-4">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-4">
        <div class="card p-5 border-l-4 border-primary-500">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-500 dark:text-gray-400">技能总数</p>
              <p class="text-2xl font-bold text-gray-900 dark:text-white">{{ stats.totalSkills }}</p>
            </div>
            <div class="w-12 h-12 bg-primary-100 dark:bg-primary-900 rounded-lg flex items-center justify-center">
              <svg class="w-6 h-6 text-primary-600 dark:text-primary-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
              </svg>
            </div>
          </div>
        </div>
        <div class="card p-5 border-l-4 border-green-500">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-500 dark:text-gray-400">版本总数</p>
              <p class="text-2xl font-bold text-gray-900 dark:text-white">{{ stats.totalVersions }}</p>
            </div>
            <div class="w-12 h-12 bg-green-100 dark:bg-green-900 rounded-lg flex items-center justify-center">
              <svg class="w-6 h-6 text-green-600 dark:text-green-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </div>
          </div>
        </div>
        <div class="card p-5 border-l-4 border-purple-500">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-500 dark:text-gray-400">分类数量</p>
              <p class="text-2xl font-bold text-gray-900 dark:text-white">{{ stats.totalCategories }}</p>
            </div>
            <div class="w-12 h-12 bg-purple-100 dark:bg-purple-900 rounded-lg flex items-center justify-center">
              <svg class="w-6 h-6 text-purple-600 dark:text-purple-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z" />
              </svg>
            </div>
          </div>
        </div>
        <div class="card p-5 border-l-4 border-orange-500">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-500 dark:text-gray-400">最近更新</p>
              <p class="text-lg font-semibold text-gray-900 dark:text-white truncate">{{ stats.lastUpdate }}</p>
            </div>
            <div class="w-12 h-12 bg-orange-100 dark:bg-orange-900 rounded-lg flex items-center justify-center">
              <svg class="w-6 h-6 text-orange-600 dark:text-orange-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
              </svg>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 主内容区域 -->
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <!-- 左侧：技能管理 -->
        <div class="lg:col-span-2 space-y-6">
          <!-- 技能列表管理 -->
          <div class="card">
            <div class="p-5 border-b border-gray-200 dark:border-gray-700">
              <h2 class="text-lg font-semibold text-gray-900 dark:text-white">技能管理</h2>
              <p class="text-sm text-gray-500 dark:text-gray-400 mt-1">管理所有技能快照</p>
            </div>
            <div class="p-5">
              <!-- 操作栏 -->
              <div class="flex flex-col sm:flex-row items-center justify-between gap-4 mb-4">
                <div class="relative w-full sm:w-64">
                  <input
                    v-model="searchQuery"
                    type="text"
                    placeholder="搜索技能..."
                    class="input-field w-full pl-10"
                  >
                  <svg class="w-4 h-4 text-gray-400 absolute left-3 top-1/2 -translate-y-1/2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                  </svg>
                </div>
                <div class="flex items-center space-x-2">
                  <select v-model="categoryFilter" class="input-field w-auto">
                    <option value="">全部分类</option>
                    <option v-for="cat in categories" :key="cat" :value="cat">{{ cat }}</option>
                  </select>
                </div>
              </div>

              <!-- 技能表格 -->
              <div class="overflow-x-auto">
                <table class="w-full">
                  <thead>
                    <tr class="text-left text-sm text-gray-500 dark:text-gray-400 border-b border-gray-200 dark:border-gray-700">
                      <th class="pb-3 font-medium">技能名称</th>
                      <th class="pb-3 font-medium hidden sm:table-cell">描述</th>
                      <th class="pb-3 font-medium">分类</th>
                      <th class="pb-3 font-medium text-right">操作</th>
                    </tr>
                  </thead>
                  <tbody class="divide-y divide-gray-100 dark:divide-gray-700">
                    <tr v-for="skill in filteredSkills" :key="skill.name" class="hover:bg-gray-50 dark:hover:bg-gray-700/50">
                      <td class="py-3">
                        <router-link :to="`/skills/${skill.name}`" class="font-medium text-primary-600 hover:text-primary-700 dark:text-primary-400">
                          {{ skill.name }}
                        </router-link>
                        <div class="text-xs text-gray-400 mt-0.5">{{ skill.latestVersion || '-' }}</div>
                      </td>
                      <td class="py-3 text-sm text-gray-600 dark:text-gray-300 hidden sm:table-cell max-w-xs truncate">
                        {{ skill.description }}
                      </td>
                      <td class="py-3">
                        <span v-if="skill.category" class="px-2 py-1 text-xs bg-gray-100 text-gray-600 rounded dark:bg-gray-700 dark:text-gray-300">
                          {{ skill.category }}
                        </span>
                      </td>
                      <td class="py-3 text-right">
                        <div class="flex items-center justify-end space-x-2">
                          <button @click="viewSkill(skill)" class="p-1.5 text-gray-400 hover:text-primary-600 transition-colors" title="查看">
                            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                            </svg>
                          </button>
                        </div>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
          </div>

          <!-- 快速操作 -->
          <div class="card">
            <div class="p-5 border-b border-gray-200 dark:border-gray-700">
              <h2 class="text-lg font-semibold text-gray-900 dark:text-white">快速操作</h2>
            </div>
            <div class="p-5">
              <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                <button @click="exportSkills" class="flex items-center p-4 border border-dashed border-gray-300 dark:border-gray-600 rounded-lg hover:border-purple-500 hover:bg-purple-50 dark:hover:bg-purple-900/20 transition-colors">
                  <div class="w-10 h-10 bg-purple-100 dark:bg-purple-900 rounded-lg flex items-center justify-center mr-3">
                    <svg class="w-5 h-5 text-purple-600 dark:text-purple-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4" />
                    </svg>
                  </div>
                  <div class="text-left">
                    <div class="font-medium text-gray-900 dark:text-white">导出数据</div>
                    <div class="text-sm text-gray-500 dark:text-gray-400">导出所有技能</div>
                  </div>
                </button>
                <button @click="showSettingsModal = true" class="flex items-center p-4 border border-dashed border-gray-300 dark:border-gray-600 rounded-lg hover:border-orange-500 hover:bg-orange-50 dark:hover:bg-orange-900/20 transition-colors">
                  <div class="w-10 h-10 bg-orange-100 dark:bg-orange-900 rounded-lg flex items-center justify-center mr-3">
                    <svg class="w-5 h-5 text-orange-600 dark:text-orange-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                    </svg>
                  </div>
                  <div class="text-left">
                    <div class="font-medium text-gray-900 dark:text-white">系统设置</div>
                    <div class="text-sm text-gray-500 dark:text-gray-400">配置管理</div>
                  </div>
                </button>
              </div>
            </div>
          </div>
        </div>

        <!-- 右侧：信息面板 -->
        <div class="space-y-6">
          <!-- 分类统计 -->
          <div class="card">
            <div class="p-5 border-b border-gray-200 dark:border-gray-700">
              <h2 class="text-lg font-semibold text-gray-900 dark:text-white">分类统计</h2>
            </div>
            <div class="p-5">
              <div v-if="Object.keys(stats.byCategory).length === 0" class="text-center py-4 text-gray-500 dark:text-gray-400 text-sm">
                暂无分类数据
              </div>
              <div v-else class="space-y-3">
                <div v-for="(count, category) in stats.byCategory" :key="category" class="flex items-center justify-between">
                  <span class="text-sm text-gray-600 dark:text-gray-300">{{ category }}</span>
                  <div class="flex items-center space-x-2">
                    <div class="w-24 h-2 bg-gray-200 dark:bg-gray-700 rounded-full overflow-hidden">
                      <div class="h-full bg-primary-500 rounded-full" :style="{ width: `${(count / stats.totalSkills) * 100}%` }"></div>
                    </div>
                    <span class="text-sm font-medium text-gray-900 dark:text-white w-6">{{ count }}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- 最近更新 -->
          <div class="card">
            <div class="p-5 border-b border-gray-200 dark:border-gray-700">
              <h2 class="text-lg font-semibold text-gray-900 dark:text-white">最近更新</h2>
            </div>
            <div class="p-5">
              <div v-if="recentSkills.length === 0" class="text-center py-4 text-gray-500 dark:text-gray-400 text-sm">
                暂无更新记录
              </div>
              <div v-else class="space-y-3">
                <router-link
                  v-for="skill in recentSkills"
                  :key="skill.name"
                  :to="`/skills/${skill.name}`"
                  class="flex items-center space-x-3 p-2 rounded-lg hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors"
                >
                  <div class="w-8 h-8 bg-primary-100 dark:bg-primary-900 rounded-lg flex items-center justify-center">
                    <svg class="w-4 h-4 text-primary-600 dark:text-primary-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                    </svg>
                  </div>
                  <div class="flex-1 min-w-0">
                    <div class="text-sm font-medium text-gray-900 dark:text-white truncate">{{ skill.name }}</div>
                    <div class="text-xs text-gray-500 dark:text-gray-400">{{ skill.latestVersion }}</div>
                  </div>
                  <div class="text-xs text-gray-400">{{ formatDate(skill.updatedAt) }}</div>
                </router-link>
              </div>
            </div>
          </div>

          <!-- 系统信息 -->
          <div class="card">
            <div class="p-5 border-b border-gray-200 dark:border-gray-700">
              <h2 class="text-lg font-semibold text-gray-900 dark:text-white">系统信息</h2>
            </div>
            <div class="p-5 space-y-3 text-sm">
              <div class="flex justify-between">
                <span class="text-gray-500 dark:text-gray-400">版本</span>
                <span class="text-gray-900 dark:text-white font-medium">v1.0.0</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-500 dark:text-gray-400">运行时间</span>
                <span class="text-gray-900 dark:text-white font-medium">{{ uptime }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-500 dark:text-gray-400">环境</span>
                <span class="px-2 py-0.5 text-xs bg-green-100 text-green-700 rounded dark:bg-green-900 dark:text-green-300">Production</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 设置弹窗 -->
    <div v-if="showSettingsModal" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/50" @click.self="showSettingsModal = false">
      <div class="card max-w-md w-full p-6">
        <h2 class="text-xl font-semibold text-gray-900 dark:text-white mb-4">系统设置</h2>
        <div class="space-y-4">
          <div class="flex items-center justify-between">
            <div>
              <div class="font-medium text-gray-900 dark:text-white">深色模式</div>
              <div class="text-sm text-gray-500 dark:text-gray-400">使用深色主题</div>
            </div>
            <button @click="toggleTheme" class="relative w-12 h-6 rounded-full transition-colors" :class="uiStore.theme === 'dark' ? 'bg-primary-600' : 'bg-gray-300 dark:bg-gray-600'">
              <span class="absolute left-1 top-1 w-4 h-4 bg-white rounded-full transition-transform" :class="uiStore.theme === 'dark' ? 'translate-x-6' : ''"></span>
            </button>
          </div>
        </div>
        <div class="flex justify-end mt-6">
          <button @click="showSettingsModal = false" class="btn-primary">关闭</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { useSkillStore } from '@/stores/skills'
import { useUiStore } from '@/stores/ui'
import dayjs from 'dayjs'
import relativeTime from 'dayjs/plugin/relativeTime'
import 'dayjs/locale/zh-cn'

dayjs.extend(relativeTime)
dayjs.locale('zh-cn')

const router = useRouter()
const skillStore = useSkillStore()
const uiStore = useUiStore()

// 状态
const searchQuery = ref('')
const categoryFilter = ref('')
const showSettingsModal = ref(false)
const uptime = ref('0天 0小时')

// 计算属性
const skills = computed(() => skillStore.skills)

const categories = computed(() => {
  const cats = new Set(skills.value.map(s => s.category).filter(Boolean) as string[])
  return Array.from(cats).sort()
})

const filteredSkills = computed(() => {
  let result = skills.value
  if (searchQuery.value) {
    const search = searchQuery.value.toLowerCase()
    result = result.filter(s =>
      s.name.toLowerCase().includes(search) ||
      s.description.toLowerCase().includes(search)
    )
  }
  if (categoryFilter.value) {
    result = result.filter(s => s.category === categoryFilter.value)
  }
  return result
})

const stats = computed(() => ({
  totalSkills: skills.value.length,
  totalVersions: skills.value.reduce((sum, s) => sum + (s.versions?.length || 0), 0),
  totalCategories: categories.value.length,
  byCategory: skills.value.reduce((acc, skill) => {
    const cat = skill.category || '未分类'
    acc[cat] = (acc[cat] || 0) + 1
    return acc
  }, {} as Record<string, number>),
  lastUpdate: skills.value.length
    ? dayjs(
        skills.value
          .map(s => s.updatedAt)
          .filter(Boolean)
          .sort()
          .reverse()[0] as string
      ).fromNow()
    : '无'
}))

const recentSkills = computed(() => {
  return [...skills.value]
    .sort((a, b) => {
      const dateA = a.updatedAt ? new Date(a.updatedAt).getTime() : 0
      const dateB = b.updatedAt ? new Date(b.updatedAt).getTime() : 0
      return dateB - dateA
    })
    .slice(0, 5)
})

// 方法
function formatDate(date: string | undefined) {
  return date ? dayjs(date).fromNow() : '-'
}

function refreshData() {
  skillStore.fetchSkills()
  skillStore.fetchCategories()
}

function viewSkill(skill: any) {
  router.push(`/skills/${skill.name}`)
}

function exportSkills() {
  const data = JSON.stringify(skills.value, null, 2)
  const blob = new Blob([data], { type: 'application/json' })
  const url = URL.createObjectURL(blob)
  const a = document.createElement('a')
  a.href = url
  a.download = `skills-${dayjs().format('YYYY-MM-DD')}.json`
  a.click()
  URL.revokeObjectURL(url)
}

function toggleTheme() {
  uiStore.toggleTheme()
}

// 运行时间计算
let uptimeInterval: ReturnType<typeof setInterval>
function updateUptime() {
  const startTime = new Date()
  startTime.setHours(startTime.getHours() - 1)

  const now = new Date()
  const diff = now.getTime() - startTime.getTime()
  const days = Math.floor(diff / (1000 * 60 * 60 * 24))
  const hours = Math.floor((diff % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60))

  uptime.value = `${days}天 ${hours}小时`
}

onMounted(() => {
  refreshData()
  updateUptime()
  uptimeInterval = setInterval(updateUptime, 60000)
})

onUnmounted(() => {
  if (uptimeInterval) clearInterval(uptimeInterval)
})
</script>
