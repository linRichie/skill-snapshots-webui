<template>
  <div class="min-h-screen">
    <!-- 页面头部 -->
    <div class="bg-white dark:bg-gray-800 border-b border-gray-200 dark:border-gray-700">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div class="flex items-center justify-between">
          <div>
            <h1 class="text-3xl font-bold text-gray-900 dark:text-white">版本历史</h1>
            <p class="mt-1 text-gray-500 dark:text-gray-400">
              所有技能的版本快照记录
            </p>
          </div>
        </div>
      </div>
    </div>

    <!-- 筛选栏 -->
    <div class="bg-white dark:bg-gray-800 border-b border-gray-200 dark:border-gray-700">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4">
        <div class="flex flex-col md:flex-row md:items-center md:justify-between gap-4">
          <div class="flex items-center space-x-3">
            <select
              v-model="filterSkill"
              class="input-field w-auto md:w-48"
            >
              <option value="">全部技能</option>
              <option v-for="skill in skills" :key="skill" :value="skill">
                {{ skill }}
              </option>
            </select>
          </div>
          <div class="text-sm text-gray-500 dark:text-gray-400">
            共 {{ filteredVersions.length }} 个版本
          </div>
        </div>
      </div>
    </div>

    <!-- 版本列表 -->
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <!-- 加载状态 -->
      <div v-if="loading" class="flex justify-center items-center py-12">
        <div class="animate-spin rounded-full h-12 w-12 border-4 border-gray-200 border-t-primary-600"></div>
      </div>

      <!-- 版本卡片网格 -->
      <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        <div
          v-for="version in paginatedVersions"
          :key="version.tag"
          class="card p-5 hover:shadow-md transition-shadow"
        >
          <div class="flex items-start justify-between mb-3">
            <div>
              <router-link
                :to="`/skills/${version.skillName}`"
                class="font-semibold text-gray-900 dark:text-white hover:text-primary-600 dark:hover:text-primary-400"
              >
                {{ version.skillName }}
              </router-link>
              <span class="ml-2 px-2 py-0.5 text-xs bg-primary-100 text-primary-700 rounded-full dark:bg-primary-900 dark:text-primary-300">
                {{ version.version }}
              </span>
            </div>
          </div>
          <p v-if="version.message" class="text-sm text-gray-600 dark:text-gray-300 line-clamp-2 mb-3">
            {{ version.message }}
          </p>
          <div class="flex items-center justify-between text-xs text-gray-500 dark:text-gray-400">
            <span v-if="version.date">{{ formatDate(version.date) }}</span>
            <code class="px-1.5 py-0.5 bg-gray-100 dark:bg-gray-700 rounded">
              {{ version.tag }}
            </code>
          </div>
        </div>
      </div>

      <!-- 分页 -->
      <div v-if="totalPages > 1" class="mt-8 flex justify-center items-center space-x-2">
        <button
          @click="currentPage > 1 && currentPage--"
          :disabled="currentPage <= 1"
          class="px-4 py-2 rounded-lg border border-gray-300 dark:border-gray-600 disabled:opacity-50 disabled:cursor-not-allowed hover:bg-gray-50 dark:hover:bg-gray-700"
        >
          上一页
        </button>
        <div class="flex items-center space-x-1">
          <button
            v-for="page in visiblePages"
            :key="page"
            @click="currentPage = page"
            class="w-10 h-10 rounded-lg"
            :class="page === currentPage
              ? 'bg-primary-600 text-white'
              : 'border border-gray-300 dark:border-gray-600 hover:bg-gray-50 dark:hover:bg-gray-700'"
          >
            {{ page }}
          </button>
        </div>
        <button
          @click="currentPage < totalPages && currentPage++"
          :disabled="currentPage >= totalPages"
          class="px-4 py-2 rounded-lg border border-gray-300 dark:border-gray-600 disabled:opacity-50 disabled:cursor-not-allowed hover:bg-gray-50 dark:hover:bg-gray-700"
        >
          下一页
        </button>
      </div>

      <!-- 空状态 -->
      <div v-if="!loading && filteredVersions.length === 0" class="card p-12 text-center">
        <svg class="w-16 h-16 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4" />
        </svg>
        <h3 class="text-lg font-medium text-gray-900 dark:text-white mb-2">暂无版本记录</h3>
        <p class="text-gray-500 dark:text-gray-400">没有找到任何版本快照</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import { useSkillStore } from '@/stores/skills'
import dayjs from 'dayjs'
import relativeTime from 'dayjs/plugin/relativeTime'
import 'dayjs/locale/zh-cn'

dayjs.extend(relativeTime)
dayjs.locale('zh-cn')

const skillStore = useSkillStore()

const loading = ref(false)
const filterSkill = ref('')
const currentPage = ref(1)
const pageSize = 12

const skills = computed(() => {
  const skillNames = new Set(skillStore.versions.map(v => v.skillName))
  return Array.from(skillNames).sort()
})

const filteredVersions = computed(() => {
  let result = [...skillStore.versions]
  if (filterSkill.value) {
    result = result.filter(v => v.skillName === filterSkill.value)
  }
  // 按日期排序
  return result.sort((a, b) => {
    const dateA = a.date ? new Date(a.date).getTime() : 0
    const dateB = b.date ? new Date(b.date).getTime() : 0
    return dateB - dateA
  })
})

const totalPages = computed(() => Math.ceil(filteredVersions.value.length / pageSize))

const paginatedVersions = computed(() => {
  const start = (currentPage.value - 1) * pageSize
  const end = start + pageSize
  return filteredVersions.value.slice(start, end)
})

const visiblePages = computed(() => {
  const pages = []
  const maxVisible = 5
  let start = Math.max(1, currentPage.value - Math.floor(maxVisible / 2))
  let end = Math.min(totalPages.value, start + maxVisible - 1)

  if (end - start < maxVisible - 1) {
    start = Math.max(1, end - maxVisible + 1)
  }

  for (let i = start; i <= end; i++) {
    pages.push(i)
  }
  return pages
})

function formatDate(date: string) {
  return dayjs(date).fromNow()
}

watch(filterSkill, () => {
  currentPage.value = 1
})

onMounted(async () => {
  loading.value = true
  try {
    await skillStore.fetchAllVersions()
  } finally {
    loading.value = false
  }
})
</script>
