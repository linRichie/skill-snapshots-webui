import axios from 'axios'
import type { ApiResponse, Skill, SkillDetail, Version, ListResponse } from '@/types'

// 创建 axios 实例
const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || '/api',
  timeout: 30000,
  headers: {
    'Content-Type': 'application/json'
  }
})

// 请求拦截器
api.interceptors.request.use(
  (config) => {
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// 响应拦截器
api.interceptors.response.use(
  (response) => {
    return response.data
  },
  (error) => {
    const message = error.response?.data?.message || error.message || '请求失败'
    return Promise.reject(new Error(message))
  }
)

/**
 * 技能相关 API
 */
export const skillApi = {
  /**
   * 获取所有技能列表
   */
  getList: (params?: { page?: number; pageSize?: number; search?: string; category?: string }) => {
    return api.get<ListResponse<Skill>>('/skills', { params })
  },

  /**
   * 获取技能详情
   */
  getDetail: (name: string, version?: string) => {
    return api.get<SkillDetail>(`/skills/${name}`, { params: { version } })
  },

  /**
   * 创建技能
   */
  create: (data: Partial<Skill>) => {
    return api.post<ApiResponse<Skill>>('/skills', data)
  },

  /**
   * 更新技能
   */
  update: (name: string, data: Partial<Skill>) => {
    return api.put<ApiResponse<Skill>>(`/skills/${name}`, data)
  },

  /**
   * 删除技能
   */
  delete: (name: string) => {
    return api.delete<ApiResponse>(`/skills/${name}`)
  },

  /**
   * 获取技能版本列表
   */
  getVersions: (name: string) => {
    return api.get<Version[]>(`/skills/${name}/versions`)
  },

  /**
   * 获取所有版本
   */
  getAllVersions: () => {
    return api.get<Version[]>('/versions')
  },

  /**
   * 获取分类列表
   */
  getCategories: () => {
    return api.get<string[]>('/categories')
  }
}

/**
 * 系统相关 API
 */
export const systemApi = {
  /**
   * 获取系统信息
   */
  getInfo: () => {
    return api.get<any>('/system/info')
  },

  /**
   * 获取统计数据
   */
  getStats: () => {
    return api.get<any>('/system/stats')
  }
}

export default api
