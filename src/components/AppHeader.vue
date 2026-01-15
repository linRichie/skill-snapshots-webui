<template>
  <header class="fixed top-0 left-0 right-0 z-50 bg-white dark:bg-gray-800 border-b border-gray-200 dark:border-gray-700">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex items-center justify-between h-16">
        <!-- Logo 和标题 -->
        <div class="flex items-center space-x-3">
          <router-link to="/" class="flex items-center space-x-3">
            <div class="w-10 h-10 bg-gradient-to-br from-primary-500 to-primary-700 rounded-lg flex items-center justify-center">
              <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
              </svg>
            </div>
            <div>
              <h1 class="text-xl font-bold text-gray-900 dark:text-white">Skill Snapshots</h1>
              <p class="text-xs text-gray-500 dark:text-gray-400">{{ $t('app.tagline') }}</p>
            </div>
          </router-link>
        </div>

        <!-- 导航链接 -->
        <nav class="hidden md:flex items-center space-x-1">
          <router-link
            v-for="item in navItems"
            :key="item.name"
            :to="item.to"
            class="px-3 py-2 rounded-lg text-sm font-medium transition-colors"
            :class="$route.name === item.name
              ? 'bg-primary-100 text-primary-700 dark:bg-primary-900 dark:text-primary-300'
              : 'text-gray-600 hover:bg-gray-100 dark:text-gray-300 dark:hover:bg-gray-700'"
          >
            {{ item.label }}
          </router-link>
        </nav>

        <!-- 右侧操作区 -->
        <div class="flex items-center space-x-3">
          <!-- 搜索框 -->
          <div class="hidden sm:block relative">
            <input
              v-model="searchQuery"
              type="text"
              :placeholder="$t('search.placeholder')"
              class="w-48 pl-9 pr-4 py-1.5 text-sm border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white"
              @keyup.enter="handleSearch"
            >
            <svg class="w-4 h-4 text-gray-400 absolute left-3 top-1/2 -translate-y-1/2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
          </div>

          <!-- 主题切换按钮 -->
          <button
            @click="uiStore.toggleTheme"
            class="p-2 rounded-lg hover:bg-gray-100 dark:hover:bg-gray-700 transition-colors"
            :title="uiStore.theme === 'light' ? '切换到深色模式' : '切换到浅色模式'"
          >
            <svg v-if="uiStore.theme === 'light'" class="w-5 h-5 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z" />
            </svg>
            <svg v-else class="w-5 h-5 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z" />
            </svg>
          </button>

          <!-- 移动端菜单按钮 -->
          <button
            @click="uiStore.toggleMobileMenu"
            class="md:hidden p-2 rounded-lg hover:bg-gray-100 dark:hover:bg-gray-700"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
            </svg>
          </button>
        </div>
      </div>
    </div>

    <!-- 移动端菜单 -->
    <div v-if="uiStore.mobileMenuOpen" class="md:hidden border-t border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800">
      <div class="px-4 py-3 space-y-1">
        <router-link
          v-for="item in navItems"
          :key="item.name"
          :to="item.to"
          class="block px-3 py-2 rounded-lg text-sm font-medium"
          :class="$route.name === item.name
            ? 'bg-primary-100 text-primary-700 dark:bg-primary-900 dark:text-primary-300'
            : 'text-gray-600 hover:bg-gray-100 dark:text-gray-300 dark:hover:bg-gray-700'"
          @click="uiStore.closeMobileMenu"
        >
          {{ item.label }}
        </router-link>
      </div>
    </div>
  </header>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useUiStore } from '@/stores/ui'
import { useSkillStore } from '@/stores/skills'

const router = useRouter()
const uiStore = useUiStore()
const skillStore = useSkillStore()

const searchQuery = ref('')

const navItems = [
  { name: 'Skills', label: '技能列表', to: '/skills' },
  { name: 'Versions', label: '版本历史', to: '/versions' },
  { name: 'Admin', label: '管理后台', to: '/admin' }
]

// 国际化支持（简单版）
const $t = (key: string) => {
  const translations: Record<string, string> = {
    'app.tagline': 'Claude Code 技能快照管理',
    'search.placeholder': '搜索技能...'
  }
  return translations[key] || key
}

function handleSearch() {
  if (searchQuery.value.trim()) {
    skillStore.setFilter('search', searchQuery.value)
    if (router.currentRoute.value.name !== 'Skills') {
      router.push('/skills')
    } else {
      skillStore.fetchSkills()
    }
  }
}

// 暴露 $t 给模板使用
defineExpose({
  $t
})
</script>
