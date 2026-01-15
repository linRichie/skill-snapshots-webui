<template>
  <router-link
    :to="`/skills/${skill.name}`"
    class="card-hover p-5 block"
  >
    <!-- 技能名称和标签 -->
    <div class="flex items-start justify-between mb-3">
      <div class="flex-1">
        <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-1">
          {{ skill.name }}
        </h3>
        <p class="text-sm text-gray-500 dark:text-gray-400 line-clamp-2">
          {{ skill.description }}
        </p>
      </div>
      <span v-if="skill.latestVersion" class="ml-3 px-2 py-1 text-xs font-medium bg-primary-100 text-primary-700 rounded-full dark:bg-primary-900 dark:text-primary-300">
        {{ skill.latestVersion }}
      </span>
    </div>

    <!-- 分类和版本数 -->
    <div class="flex items-center space-x-4 text-sm text-gray-500 dark:text-gray-400">
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
        {{ formatDate(skill.updatedAt) }}
      </span>
    </div>

    <!-- 标签 -->
    <div v-if="skill.tags && skill.tags.length" class="mt-3 flex flex-wrap gap-2">
      <span
        v-for="tag in skill.tags.slice(0, 3)"
        :key="tag"
        class="px-2 py-0.5 text-xs bg-gray-100 text-gray-600 rounded dark:bg-gray-700 dark:text-gray-300"
      >
        {{ tag }}
      </span>
      <span v-if="skill.tags.length > 3" class="px-2 py-0.5 text-xs text-gray-500">
        +{{ skill.tags.length - 3 }}
      </span>
    </div>
  </router-link>
</template>

<script setup lang="ts">
import type { Skill } from '@/types'
import dayjs from 'dayjs'
import relativeTime from 'dayjs/plugin/relativeTime'
import 'dayjs/locale/zh-cn'

dayjs.extend(relativeTime)
dayjs.locale('zh-cn')

defineProps<{
  skill: Skill
}>()

function formatDate(date: string) {
  return dayjs(date).fromNow()
}
</script>
