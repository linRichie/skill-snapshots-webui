import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'Home',
    redirect: '/skills'
  },
  {
    path: '/skills',
    name: 'Skills',
    component: () => import('@/views/SkillList.vue'),
    meta: { title: '技能列表' }
  },
  {
    path: '/skills/:name',
    name: 'SkillDetail',
    component: () => import('@/views/SkillDetail.vue'),
    meta: { title: '技能详情' }
  },
  {
    path: '/skills/:name/versions',
    name: 'SkillVersions',
    component: () => import('@/views/SkillVersions.vue'),
    meta: { title: '版本历史' }
  },
  {
    path: '/versions',
    name: 'Versions',
    component: () => import('@/views/VersionList.vue'),
    meta: { title: '所有版本' }
  },
  {
    path: '/admin',
    name: 'Admin',
    component: () => import('@/views/AdminPanel.vue'),
    meta: { title: '管理后台' }
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: () => import('@/views/NotFound.vue'),
    meta: { title: '页面未找到' }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// 设置页面标题
router.beforeEach((to, _from, next) => {
  const title = to.meta.title as string
  document.title = title ? `${title} - Skill Snapshots` : 'Skill Snapshots'
  next()
})

export default router
