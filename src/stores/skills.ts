import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { Skill, SkillDetail, Version } from '@/types'
import { skillApi } from '@/api'

/**
 * 技能数据存储
 */
export const useSkillStore = defineStore('skills', () => {
  // 状态
  const skills = ref<Skill[]>([])
  const currentSkill = ref<SkillDetail | null>(null)
  const versions = ref<Version[]>([])
  const categories = ref<string[]>([])
  const loading = ref(false)
  const error = ref<string | null>(null)

  // 分页
  const pagination = ref({
    page: 1,
    pageSize: 20,
    total: 0,
    totalPages: 0
  })

  // 筛选
  const filters = ref({
    search: '',
    category: ''
  })

  // 计算属性
  const filteredSkills = computed(() => {
    let result = skills.value

    if (filters.value.search) {
      const search = filters.value.search.toLowerCase()
      result = result.filter(skill =>
        skill.name.toLowerCase().includes(search) ||
        skill.description.toLowerCase().includes(search)
      )
    }

    if (filters.value.category) {
      result = result.filter(skill => skill.category === filters.value.category)
    }

    return result
  })

  const stats = computed(() => ({
    total: skills.value.length,
    byCategory: skills.value.reduce((acc, skill) => {
      const cat = skill.category || '未分类'
      acc[cat] = (acc[cat] || 0) + 1
      return acc
    }, {} as Record<string, number>)
  }))

  // 方法
  async function fetchSkills(page = 1, pageSize = 20) {
    loading.value = true
    error.value = null
    try {
      const response = await skillApi.getList({
        page,
        pageSize,
        search: filters.value.search || undefined,
        category: filters.value.category || undefined
      })
      skills.value = response.items
      pagination.value = response.pagination
    } catch (err) {
      error.value = err instanceof Error ? err.message : '获取技能列表失败'
      console.error('获取技能列表失败:', err)
    } finally {
      loading.value = false
    }
  }

  async function fetchSkillDetail(name: string, version?: string) {
    loading.value = true
    error.value = null
    try {
      const response = await skillApi.getDetail(name, version)
      currentSkill.value = response
      return response
    } catch (err) {
      error.value = err instanceof Error ? err.message : '获取技能详情失败'
      console.error('获取技能详情失败:', err)
      return null
    } finally {
      loading.value = false
    }
  }

  async function fetchVersions(name: string) {
    try {
      const response = await skillApi.getVersions(name)
      versions.value = response
      return response
    } catch (err) {
      console.error('获取版本列表失败:', err)
      return []
    }
  }

  async function fetchAllVersions() {
    try {
      const response = await skillApi.getAllVersions()
      versions.value = response
      return response
    } catch (err) {
      console.error('获取所有版本失败:', err)
      return []
    }
  }

  async function fetchCategories() {
    try {
      const response = await skillApi.getCategories()
      categories.value = response
      return response
    } catch (err) {
      console.error('获取分类失败:', err)
      return []
    }
  }

  function setFilter(key: 'search' | 'category', value: string) {
    filters.value[key] = value
  }

  function clearCurrentSkill() {
    currentSkill.value = null
  }

  return {
    // 状态
    skills,
    currentSkill,
    versions,
    categories,
    loading,
    error,
    pagination,
    filters,
    // 计算属性
    filteredSkills,
    stats,
    // 方法
    fetchSkills,
    fetchSkillDetail,
    fetchVersions,
    fetchAllVersions,
    fetchCategories,
    setFilter,
    clearCurrentSkill
  }
})
