import { key } from '../store'
import { useStore } from 'vuex'

import { VueCookieNext } from 'vue-cookie-next'

export const importView = (viewName: string) => () =>
  import('../views/' + viewName + '.vue')

export const importCom = (viewName: string) => () =>
  import('../components/' + viewName + '.vue')

// 打印环境
/**
 *
 */
export function printEnv() {
  if (import.meta.env.DEV) print('mode: development')
  else print('mode: production')
}

export function print(o: string) {
  console.log('🌸 %c' + o, 'color:pink')
}

// not allowed
// export function getToken(): String {
//   const store = useStore(key)
//   return store.state.token
// }

export function getToken(): String {
  let t = VueCookieNext.getCookie('token')
  return t
}