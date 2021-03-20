import { createRouter, createWebHashHistory } from 'vue-router'
import { _import_view } from '@/utils'

export const routerMap = [
  {
    path: '/',
    component: _import_view('Home')
  }
]

export const asyncRouterMap = []

export const router = createRouter({
  history: createWebHashHistory(),
  routes: routerMap
})
