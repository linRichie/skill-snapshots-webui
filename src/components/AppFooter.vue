<template>
  <footer class="mt-auto py-6 border-t border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex flex-col md:flex-row items-center justify-between space-y-4 md:space-y-0">
        <!-- 版权信息 -->
        <div class="text-sm text-gray-500 dark:text-gray-400">
          <p>&copy; {{ currentYear }} Skill Snapshots WebUI. Released under MIT License.</p>
        </div>

        <!-- 链接 -->
        <div class="flex items-center space-x-6 text-sm">
          <a href="https://github.com/your-username/skill-snapshots-webui" target="_blank" rel="noopener noreferrer" class="text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-300">
            GitHub
          </a>
          <a href="https://docs.anthropic.com" target="_blank" rel="noopener noreferrer" class="text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-300">
            Claude Docs
          </a>
        </div>
      </div>
    </div>
  </footer>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const currentYear = computed(() => new Date().getFullYear())
</script>
