<template>
  <div class="min-h-screen">
    <!-- 面包屑导航 -->
    <div class="bg-white dark:bg-gray-800 border-b border-gray-200 dark:border-gray-700">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4">
        <nav class="flex items-center space-x-2 text-sm">
          <router-link to="/" class="text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-300">
            首页
          </router-link>
          <svg class="w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
          </svg>
          <router-link to="/skills" class="text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-300">
            技能列表
          </router-link>
          <svg class="w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
          </svg>
          <router-link :to="`/skills/${skillName}`" class="text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-300">
            {{ skillName }}
          </router-link>
          <svg class="w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
          </svg>
          <span class="text-gray-900 dark:text-white font-medium">版本历史</span>
        </nav>
      </div>
    </div>

    <!-- 页面头部 -->
    <div class="bg-white dark:bg-gray-800 border-b border-gray-200 dark:border-gray-700">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
        <div class="flex items-center justify-between">
          <div>
            <h1 class="text-2xl font-bold text-gray-900 dark:text-white">{{ skillName }} - 版本历史</h1>
            <p class="mt-1 text-gray-500 dark:text-gray-400">共 {{ versions.length }} 个版本</p>
          </div>
          <router-link :to="`/skills/${skillName}`" class="btn-secondary">
            返回详情
          </router-link>
        </div>
      </div>
    </div>

    <!-- 版本列表 -->
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <!-- 加载状态 -->
      <div v-if="loading" class="flex justify-center items-center py-12">
        <div class="animate-spin rounded-full h-12 w-12 border-4 border-gray-200 border-t-primary-600"></div>
      </div>

      <!-- 版本时间线 -->
      <div v-else class="card p-6">
        <div v-if="versions.length === 0" class="text-center py-12 text-gray-500 dark:text-gray-400">
          <svg class="w-16 h-16 mx-auto mb-4 text-gray-300 dark:text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          <p>暂无版本记录</p>
        </div>

        <div v-else class="relative">
          <!-- 时间线 -->
          <div class="absolute left-4 top-0 bottom-0 w-0.5 bg-gray-200 dark:bg-gray-700"></div>

          <div class="space-y-6">
            <div
              v-for="(version, index) in versions"
              :key="version.tag"
              class="relative pl-10"
            >
              <!-- 时间线节点 -->
              <div class="absolute left-0 w-8 h-8 bg-primary-500 rounded-full flex items-center justify-center">
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                </svg>
              </div>

              <!-- 版本卡片 -->
              <div class="card p-5 hover:shadow-md transition-shadow">
                <div class="flex items-start justify-between mb-2">
                  <div>
                    <h3 class="text-lg font-semibold text-gray-900 dark:text-white">
                      {{ version.version }}
                    </h3>
                    <p class="text-sm text-gray-500 dark:text-gray-400">
                      Tag: <code class="px-1.5 py-0.5 bg-gray-100 dark:bg-gray-700 rounded text-xs">{{ version.tag }}</code>
                    </p>
                  </div>
                  <div class="flex items-center space-x-2">
                    <span v-if="version.date" class="text-sm text-gray-500 dark:text-gray-400">
                      {{ formatDate(version.date) }}
                    </span>
                  </div>
                </div>
                <p v-if="version.message" class="text-gray-600 dark:text-gray-300">
                  {{ version.message }}
                </p>
                <div v-if="version.author" class="mt-2 flex items-center space-x-2 text-sm text-gray-500 dark:text-gray-400">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                  </svg>
                  <span>{{ version.author }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useRoute } from 'vue-router'
import { useSkillStore } from '@/stores/skills'
import dayjs from 'dayjs'
import relativeTime from 'dayjs/plugin/relativeTime'
import 'dayjs/locale/zh-cn'

dayjs.extend(relativeTime)
dayjs.locale('zh-cn')

const route = useRoute()
const skillStore = useSkillStore()

const loading = ref(false)
const versions = ref<any[]>([])

const skillName = computed(() => route.params.name as string)

function formatDate(date: string) {
  return dayjs(date).fromNow()
}

onMounted(async () => {
  loading.value = true
  try {
    const data = await skillStore.fetchVersions(skillName.value)
    versions.value = data || []
  } catch (error) {
    console.error('获取版本列表失败:', error)
  } finally {
    loading.value = false
  }
})
</script>
