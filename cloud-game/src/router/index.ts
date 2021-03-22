import { createRouter, createWebHashHistory } from 'vue-router'
import { importView } from '../utils'

export const routerMap = [
  {
    path: '/',
    component: importView('HomePage'),
  },
  {
    path: '/home',
    component: importView('HomePage'),
  },
]

export const asyncRouterMap = []

export const router = createRouter({
  history: createWebHashHistory(),
  routes: routerMap,
})
