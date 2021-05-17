import {
  createRouter,
  createWebHashHistory,
  createWebHistory,
} from 'vue-router'
import { importView } from '../utils'

export const routerMap = [
  {
    name: 'MIndex',
    path: '/m',
    component: importView('MIndex'),
    redirect: '/m/game',
    children: [
      {
        name: 'MMe',
        path: 'me',
        component: importView('MMe'),
      },
      {
        name: 'MCom',
        path: 'com',
        component: importView('MCom'),
      },
      {
        name: 'MGame',
        path: 'game',
        component: importView('MGame'),
      },
      {
        name: 'MLogin',
        path: 'login',
        component: importView('MLogin'),
      },
      {
        name: 'MAbout',
        path: 'about',
        component: importView('MAbout'),
      },
      {
        name: 'MSetting',
        path: 'setting',
        component: importView('MSetting'),
      },
      {
        name: 'MCharge',
        path: 'charge',
        component: importView('MCharge'),
      },
      {
        name: 'MPlay',
        path: 'play',
        component: importView('MPlay'),
      },
    ],
  },
  {
    name: 'PCPage',
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
        name: 'About',
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
  // history: createWebHistory(),
  history: createWebHashHistory(),
  routes: routerMap,
})
