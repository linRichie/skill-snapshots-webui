<template>
  <div class="min-h-screen">
    <!-- 加载状态 -->
    <div v-if="skillStore.loading" class="flex justify-center items-center py-20">
      <div class="animate-spin rounded-full h-16 w-16 border-4 border-gray-200 border-t-primary-600"></div>
    </div>

    <!-- 技能详情 -->
    <div v-else-if="skill">
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
            <span class="text-gray-900 dark:text-white font-medium">{{ skill.name }}</span>
          </nav>
        </div>
      </div>

      <!-- 技能头部 -->
      <div class="bg-white dark:bg-gray-800 border-b border-gray-200 dark:border-gray-700">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
          <div class="flex flex-col md:flex-row md:items-start md:justify-between gap-4">
            <div class="flex-1">
              <div class="flex items-center space-x-3 mb-3">
                <h1 class="text-3xl font-bold text-gray-900 dark:text-white">{{ skill.name }}</h1>
                <span v-if="skill.latestVersion" class="px-3 py-1 text-sm font-medium bg-primary-100 text-primary-700 rounded-full dark:bg-primary-900 dark:text-primary-300">
                  {{ skill.latestVersion }}
                </span>
              </div>
              <p class="text-lg text-gray-600 dark:text-gray-300">{{ skill.description }}</p>
              <div class="mt-4 flex flex-wrap items-center gap-4 text-sm text-gray-500 dark:text-gray-400">
                <span v-if="skill.category" class="flex items-center">
                  <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z" />
                  </svg>
                  {{ skill.category }}
                </span>
                <span class="flex items-center">
                  <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                  {{ skill.versions?.length || 0 }} 个版本
                </span>
                <span v-if="skill.updatedAt" class="flex items-center">
                  <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                  </svg>
                  更新于 {{ formatDate(skill.updatedAt) }}
                </span>
              </div>
              <!-- 标签 -->
              <div v-if="skill.tags && skill.tags.length" class="mt-3 flex flex-wrap gap-2">
                <span
                  v-for="tag in skill.tags"
                  :key="tag"
                  class="px-3 py-1 text-sm bg-gray-100 text-gray-600 rounded-full dark:bg-gray-700 dark:text-gray-300"
                >
                  {{ tag }}
                </span>
              </div>
            </div>
            <!-- 操作按钮 -->
            <div class="flex items-center space-x-2">
              <router-link :to="`/skills/${skill.name}/versions`" class="btn-secondary">
                <span class="flex items-center">
                  <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                  版本历史
                </span>
              </router-link>
              <router-link to="/admin" class="btn-primary">
                <span class="flex items-center">
                  <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                  </svg>
                  管理
                </span>
              </router-link>
            </div>
          </div>
        </div>
      </div>

      <!-- 内容区域 -->
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div class="grid grid-cols-1 lg:grid-cols-4 gap-8">
          <!-- 左侧主要内容 -->
          <div class="lg:col-span-3">
            <!-- 技能文档 -->
            <div class="card p-6 mb-6">
              <h2 class="text-xl font-semibold text-gray-900 dark:text-white mb-4">技能文档</h2>
              <div v-if="skill.content" class="prose prose-slate max-w-none dark:prose-invert">
                <MarkdownRenderer :content="skill.content" />
              </div>
              <div v-else class="text-center py-8 text-gray-500 dark:text-gray-400">
                <svg class="w-12 h-12 mx-auto mb-3 text-gray-300 dark:text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                </svg>
                <p>暂无文档内容</p>
              </div>
            </div>

            <!-- 脚本列表 -->
            <div v-if="skill.scripts && skill.scripts.length" class="card p-6">
              <h2 class="text-xl font-semibold text-gray-900 dark:text-white mb-4">脚本文件</h2>
              <div class="space-y-2">
                <div
                  v-for="script in skill.scripts"
                  :key="script.name"
                  class="flex items-center justify-between p-3 bg-gray-50 dark:bg-gray-700 rounded-lg"
                >
                  <div class="flex items-center space-x-3">
                    <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                    </svg>
                    <div>
                      <div class="font-medium text-gray-900 dark:text-white">{{ script.name }}</div>
                      <div v-if="script.description" class="text-sm text-gray-500 dark:text-gray-400">{{ script.description }}</div>
                    </div>
                  </div>
                  <code class="text-xs bg-gray-200 dark:bg-gray-600 px-2 py-1 rounded">
                    {{ script.path }}
                  </code>
                </div>
              </div>
            </div>
          </div>

          <!-- 右侧信息栏 -->
          <div class="lg:col-span-1">
            <!-- 版本列表 -->
            <div class="card p-5 mb-6">
              <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-4">版本列表</h3>
              <div v-if="skill.versions && skill.versions.length" class="space-y-2">
                <router-link
                  v-for="version in skill.versions.slice(0, 5)"
                  :key="version.tag"
                  :to="`/skills/${skill.name}/versions`"
                  class="block p-3 rounded-lg hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors"
                >
                  <div class="flex items-center justify-between">
                    <span class="font-medium text-gray-900 dark:text-white">{{ version.version }}</span>
                    <span v-if="version.date" class="text-xs text-gray-500 dark:text-gray-400">
                      {{ formatDateShort(version.date) }}
                    </span>
                  </div>
                  <div v-if="version.message" class="text-sm text-gray-500 dark:text-gray-400 truncate mt-1">
                    {{ version.message }}
                  </div>
                </router-link>
                <router-link
                  v-if="skill.versions.length > 5"
                  :to="`/skills/${skill.name}/versions`"
                  class="block text-center text-sm text-primary-600 hover:text-primary-700 dark:text-primary-400 py-2"
                >
                  查看全部 {{ skill.versions.length }} 个版本
                </router-link>
              </div>
              <div v-else class="text-center py-4 text-gray-500 dark:text-gray-400 text-sm">
                暂无版本
              </div>
            </div>

            <!-- 依赖信息 -->
            <div v-if="skill.dependencies && skill.dependencies.length" class="card p-5">
              <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-4">依赖项</h3>
              <div class="flex flex-wrap gap-2">
                <span
                  v-for="dep in skill.dependencies"
                  :key="dep"
                  class="px-2 py-1 text-sm bg-gray-100 text-gray-700 rounded dark:bg-gray-700 dark:text-gray-300"
                >
                  {{ dep }}
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 错误状态 -->
    <div v-else class="card p-12 text-center">
      <svg class="w-16 h-16 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
      </svg>
      <h3 class="text-lg font-medium text-gray-900 dark:text-white mb-2">技能不存在</h3>
      <p class="text-gray-500 dark:text-gray-400 mb-4">未找到名为 {{ $route.params.name }} 的技能</p>
      <router-link to="/skills" class="btn-primary">返回技能列表</router-link>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useSkillStore } from '@/stores/skills'
import MarkdownRenderer from '@/components/MarkdownRenderer.vue'
import dayjs from 'dayjs'
import relativeTime from 'dayjs/plugin/relativeTime'
import 'dayjs/locale/zh-cn'

dayjs.extend(relativeTime)
dayjs.locale('zh-cn')

const route = useRoute()
const skillStore = useSkillStore()

const skill = skillStore.currentSkill

function formatDate(date: string) {
  return dayjs(date).fromNow()
}

function formatDateShort(date: string) {
  return dayjs(date).format('MM-DD')
}

onMounted(async () => {
  const name = route.params.name as string
  await skillStore.fetchSkillDetail(name)
  if (skillStore.currentSkill) {
    await skillStore.fetchVersions(name)
  }
})
</script>
