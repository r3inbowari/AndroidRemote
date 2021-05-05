import { createRouter, createWebHashHistory } from 'vue-router'
import { importView } from '../utils'

export const routerMap = [
  {
    path: '/',
    component: importView('HomePage'),
  },
  {
    name: 'Home',
    path: '/home',
    component: importView('Home'),
  },
  {
    name: 'Play',
    path: '/play',
    component: importView('GP'),
  },
]

export const asyncRouterMap = []

export const router = createRouter({
  history: createWebHashHistory(),
  routes: routerMap,
})
