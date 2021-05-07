import {
  createRouter,
  createWebHashHistory,
  createWebHistory,
} from 'vue-router'
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
  {
    name: 'Game',
    path: '/game/:aid',
    component: importView('Game'),
  },
]

export const asyncRouterMap = []

export const router = createRouter({
  history: createWebHistory(),
  routes: routerMap,
})
