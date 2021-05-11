import {
  createRouter,
  createWebHashHistory,
  createWebHistory,
} from 'vue-router'
import { importView } from '../utils'

export const routerMap = [
  {
    path: '/',
    component: importView('Home'),
  },
  {
    name: 'Home',
    path: '/home',
    component: importView('HomePage'),
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
  {
    name: 'Login',
    path: '/login',
    component: importView('LoginPage'),
  },
  {
    name: 'Center',
    path: '/center',
    component: importView('Manager'),
    children: [
      {
        path: 'about',
        component: importView('About'),
      },
      {
        path: 'log',
        component: importView('Log'),
      },
      {
        path: 'setting',
        component: importView('Setting'),
      },
      {
        path: 'measure',
        component: importView('Measure'),
      },
      {
        path: 'global',
        component: importView('Global'),
      },
      {
        path: 'dashboard',
        component: importView('Dashboard'),
      },
      {
        path: 'connection',
        component: importView('Connection'),
      },
    ],
  },
]

export const asyncRouterMap = []

export const router = createRouter({
  history: createWebHistory(),
  routes: routerMap,
})
