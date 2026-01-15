<template>
  <div
    class="markdown-content prose prose-slate max-w-none dark:prose-invert"
    v-html="renderedHtml"
  />
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { marked } from 'marked'
import DOMPurify from 'dompurify'

const props = defineProps<{
  content: string
}>()

// 配置 marked
marked.setOptions({
  breaks: true,
  gfm: true
})

const renderedHtml = computed(() => {
  if (!props.content) return ''
  const html = marked(props.content) as string
  return DOMPurify.sanitize(html)
})
</script>
