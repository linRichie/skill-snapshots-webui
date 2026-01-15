<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50 dark:bg-gray-900 py-12 px-4">
    <div class="max-w-md w-full">
      <div class="card p-8 text-center">
        <!-- 404 图标 -->
        <div class="mb-6">
          <div class="inline-flex items-center justify-center w-24 h-24 rounded-full bg-primary-100 dark:bg-primary-900">
            <svg class="w-12 h-12 text-primary-600 dark:text-primary-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
        </div>

        <!-- 标题 -->
        <h1 class="text-4xl font-bold text-gray-900 dark:text-white mb-2">404</h1>
        <h2 class="text-xl font-semibold text-gray-700 dark:text-gray-200 mb-4">页面未找到</h2>

        <!-- 描述 -->
        <p class="text-gray-500 dark:text-gray-400 mb-8">
          抱歉，您访问的页面不存在。可能是链接错误或页面已被删除。
        </p>

        <!-- 操作按钮 -->
        <div class="flex flex-col sm:flex-row items-center justify-center space-y-3 sm:space-y-0 sm:space-x-4">
          <router-link to="/skills" class="btn-primary w-full sm:w-auto">
            <span class="flex items-center justify-center">
              <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16" />
              </svg>
              技能列表
            </span>
          </router-link>
          <button @click="goBack" class="btn-secondary w-full sm:w-auto">
            <span class="flex items-center justify-center">
              <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18" />
              </svg>
              返回上页
            </span>
          </button>
        </div>
      </div>

      <!-- 快速链接 -->
      <div class="mt-6">
        <p class="text-center text-sm text-gray-500 dark:text-gray-400 mb-4">快速访问</p>
        <div class="grid grid-cols-3 gap-3">
          <router-link to="/skills" class="card p-3 text-center hover:shadow-md transition-shadow">
            <svg class="w-6 h-6 mx-auto mb-1 text-primary-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
            </svg>
            <div class="text-xs text-gray-600 dark:text-gray-300">技能</div>
          </router-link>
          <router-link to="/versions" class="card p-3 text-center hover:shadow-md transition-shadow">
            <svg class="w-6 h-6 mx-auto mb-1 text-green-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <div class="text-xs text-gray-600 dark:text-gray-300">版本</div>
          </router-link>
          <router-link to="/admin" class="card p-3 text-center hover:shadow-md transition-shadow">
            <svg class="w-6 h-6 mx-auto mb-1 text-orange-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
            </svg>
            <div class="text-xs text-gray-600 dark:text-gray-300">管理</div>
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()

function goBack() {
  if (window.history.length > 1) {
    router.back()
  } else {
    router.push('/skills')
  }
}
</script>
