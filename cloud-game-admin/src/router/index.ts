import {
  createRouter,
  createWebHashHistory,
  createWebHistory,
} from 'vue-router'
import { importView } from '../utils'

export const routerMap = [
  {
    path: '/',
    component: importView('Login'),
  },
]

export const asyncRouterMap = []

export const router = createRouter({
  history: createWebHistory(),
  routes: routerMap,
})
