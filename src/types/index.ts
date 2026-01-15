/**
 * 技能类型定义
 */
export interface Skill {
  name: string
  description: string
  path: string
  versions: Version[]
  latestVersion?: string
  category?: string
  tags?: string[]
  createdAt?: string
  updatedAt?: string
  content?: string
  dependencies?: string[]
  scripts?: Script[]
}

/**
 * 版本类型定义
 */
export interface Version {
  tag: string
  skillName: string
  version: string
  message?: string
  date?: string
  author?: string
}

/**
 * 技能详情类型
 */
export interface SkillDetail extends Skill {
  content?: string
  dependencies?: string[]
  scripts?: Script[]
}

/**
 * 脚本类型定义
 */
export interface Script {
  name: string
  path: string
  description?: string
}

/**
 * API 响应类型
 */
export interface ApiResponse<T = any> {
  success: boolean
  data?: T
  message?: string
  error?: string
}

/**
 * 分页类型
 */
export interface Pagination {
  page: number
  pageSize: number
  total: number
  totalPages: number
}

/**
 * 列表响应类型
 */
export interface ListResponse<T> {
  items: T[]
  pagination: Pagination
}
