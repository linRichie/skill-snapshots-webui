<template>
  <div class="min-h-screen">
    <!-- 页面头部 -->
    <div class="bg-white dark:bg-gray-800 border-b border-gray-200 dark:border-gray-700">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div class="flex flex-col md:flex-row md:items-center md:justify-between gap-4">
          <div>
            <h1 class="text-3xl font-bold text-gray-900 dark:text-white">技能列表</h1>
            <p class="mt-1 text-gray-500 dark:text-gray-400">
              共 {{ skillStore.skills.length }} 个技能，{{ totalVersions }} 个版本快照
            </p>
          </div>
          <div class="flex items-center space-x-3">
            <!-- 分类筛选 -->
            <select
              v-model="selectedCategory"
              class="input-field w-auto"
              @change="handleCategoryChange"
            >
              <option value="">全部分类</option>
              <option v-for="cat in categories" :key="cat" :value="cat">
                {{ cat }} ({{ getCategoryCount(cat) }})
              </option>
            </select>
            <!-- 视图切换 -->
            <div class="flex items-center border border-gray-300 dark:border-gray-600 rounded-lg overflow-hidden">
              <button
                @click="viewMode = 'grid'"
                class="p-2 transition-colors"
                :class="viewMode === 'grid' ? 'bg-primary-500 text-white' : 'hover:bg-gray-100 dark:hover:bg-gray-700'"
              >
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z" />
                </svg>
              </button>
              <button
                @click="viewMode = 'list'"
                class="p-2 transition-colors"
                :class="viewMode === 'list' ? 'bg-primary-500 text-white' : 'hover:bg-gray-100 dark:hover:bg-gray-700'"
              >
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16" />
                </svg>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 统计卡片 -->
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
      <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
        <div class="card p-4">
          <div class="text-sm text-gray-500 dark:text-gray-400">总技能数</div>
          <div class="text-2xl font-bold text-gray-900 dark:text-white">{{ skillStore.skills.length }}</div>
        </div>
        <div class="card p-4">
          <div class="text-sm text-gray-500 dark:text-gray-400">总版本数</div>
          <div class="text-2xl font-bold text-gray-900 dark:text-white">{{ totalVersions }}</div>
        </div>
        <div class="card p-4">
          <div class="text-sm text-gray-500 dark:text-gray-400">分类数</div>
          <div class="text-2xl font-bold text-gray-900 dark:text-white">{{ categories.length }}</div>
        </div>
        <div class="card p-4">
          <div class="text-sm text-gray-500 dark:text-gray-400">最近更新</div>
          <div class="text-lg font-semibold text-gray-900 dark:text-white truncate">{{ lastUpdate }}</div>
        </div>
      </div>
    </div>

    <!-- 技能列表 -->
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 pb-12">
      <!-- 加载状态 -->
      <div v-if="skillStore.loading" class="flex justify-center items-center py-12">
        <div class="animate-spin rounded-full h-12 w-12 border-4 border-gray-200 border-t-primary-600"></div>
      </div>

      <!-- 错误状态 -->
      <div v-else-if="skillStore.error" class="card p-12 text-center">
        <svg class="w-16 h-16 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
        <h3 class="text-lg font-medium text-gray-900 dark:text-white mb-2">加载失败</h3>
        <p class="text-gray-500 dark:text-gray-400 mb-4">{{ skillStore.error }}</p>
        <button @click="loadSkills" class="btn-primary">重试</button>
      </div>

      <!-- 空状态 -->
      <div v-else-if="filteredSkills.length === 0" class="card p-12 text-center">
        <svg class="w-16 h-16 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4" />
        </svg>
        <h3 class="text-lg font-medium text-gray-900 dark:text-white mb-2">暂无技能</h3>
        <p class="text-gray-500 dark:text-gray-400">没有找到匹配的技能</p>
      </div>

      <!-- 网格视图 -->
      <div v-else-if="viewMode === 'grid'" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        <SkillCard
          v-for="skill in filteredSkills"
          :key="skill.name"
          :skill="skill"
        />
      </div>

      <!-- 列表视图 -->
      <div v-else class="card overflow-hidden">
        <table class="w-full">
          <thead class="bg-gray-50 dark:bg-gray-700">
            <tr>
              <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">技能名称</th>
              <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider hidden sm:table-cell">描述</th>
              <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider hidden md:table-cell">分类</th>
              <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">版本</th>
              <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider hidden lg:table-cell">更新时间</th>
              <th class="px-4 py-3 text-right text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">操作</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200 dark:divide-gray-700">
            <tr
              v-for="skill in filteredSkills"
              :key="skill.name"
              class="hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors"
            >
              <td class="px-4 py-3">
                <router-link :to="`/skills/${skill.name}`" class="font-medium text-primary-600 hover:text-primary-700 dark:text-primary-400">
                  {{ skill.name }}
                </router-link>
              </td>
              <td class="px-4 py-3 text-sm text-gray-500 dark:text-gray-400 hidden sm:table-cell max-w-xs truncate">
                {{ skill.description }}
              </td>
              <td class="px-4 py-3 hidden md:table-cell">
                <span v-if="skill.category" class="px-2 py-1 text-xs bg-gray-100 text-gray-600 rounded dark:bg-gray-600 dark:text-gray-300">
                  {{ skill.category }}
                </span>
              </td>
              <td class="px-4 py-3 text-sm">
                <span class="px-2 py-1 text-xs bg-primary-100 text-primary-700 rounded-full dark:bg-primary-900 dark:text-primary-300">
                  {{ skill.latestVersion || '-' }}
                </span>
              </td>
              <td class="px-4 py-3 text-sm text-gray-500 dark:text-gray-400 hidden lg:table-cell">
                {{ skill.updatedAt ? formatDate(skill.updatedAt) : '-' }}
              </td>
              <td class="px-4 py-3 text-right">
                <router-link :to="`/skills/${skill.name}`" class="text-primary-600 hover:text-primary-700 dark:text-primary-400 text-sm">
                  查看
                </router-link>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useSkillStore } from '@/stores/skills'
import SkillCard from '@/components/SkillCard.vue'
import dayjs from 'dayjs'
import relativeTime from 'dayjs/plugin/relativeTime'
import 'dayjs/locale/zh-cn'

dayjs.extend(relativeTime)
dayjs.locale('zh-cn')

const skillStore = useSkillStore()

const viewMode = ref<'grid' | 'list'>('grid')
const selectedCategory = ref('')

// 计算属性
const filteredSkills = computed(() => {
  let result = skillStore.skills
  if (selectedCategory.value) {
    result = result.filter(s => s.category === selectedCategory.value)
  }
  return result
})

const totalVersions = computed(() => {
  return skillStore.skills.reduce((sum, skill) => sum + (skill.versions?.length || 0), 0)
})

const categories = computed(() => {
  const cats = new Set(skillStore.skills.map(s => s.category).filter(Boolean) as string[])
  return Array.from(cats).sort()
})

const lastUpdate = computed(() => {
  const dates = skillStore.skills
    .map(s => s.updatedAt)
    .filter(Boolean)
    .sort()
    .reverse()
  return dates.length ? dayjs(dates[0]).fromNow() : '无'
})

// 方法
function loadSkills() {
  skillStore.fetchSkills()
  skillStore.fetchCategories()
}

function handleCategoryChange() {
  skillStore.setFilter('category', selectedCategory.value)
}

function getCategoryCount(cat: string) {
  return skillStore.skills.filter(s => s.category === cat).length
}

function formatDate(date: string) {
  return dayjs(date).format('YYYY-MM-DD')
}

// 生命周期
onMounted(() => {
  loadSkills()
})
</script>
