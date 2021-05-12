// import { State } from '@vue/runtime-core'
// @ts-ignore
import { InjectionKey, App } from 'vue'

import { createStore, StoreOptions, Store } from 'vuex'

// import app from "../main"
import { getCurrentInstance } from 'vue'
import { IWebSocket } from '../ws'

// declare module '@vue/runtime-core' {
// 	// define your typings for the store state
// 	export interface State {
// 		count: number
// 	}

// 	interface ComponentCustomProperties {
// 		$store: Store<State>
// 	}
// }

export interface State {
  count: number
  avatar: String
  token: String
  ws: IWebSocket | null
  theme: boolean
  session: String
}

// define injection key
export const key: InjectionKey<Store<State>> = Symbol()

export const store = createStore<State>({
  state: {
    count: 1,
    avatar: 'null',
    token: '',
    // ws global
    ws: null,
    theme: true,
    session: '',
  },
  mutations: {
    setTheme(state, t) {
      state.theme = t
      console.log('[store] setTheme ' + state.theme)
    },
    increment(state) {
      // mutate state
      state.count++
    },
    setToken(state, t) {
      state.token = t
      console.log('[store] setToken ' + state.token)
    },
    setAvatar(state, t) {
      state.avatar = t
      console.log('[store] setAvatar ' + state.avatar)
    },
    setWs(state, t) {
      state.ws = t
      console.log('[store] setWs ')
    },
    setSession(state, t) {
      state.session = t
      console.log('[store] setSession ' + state.session)
    },
  },
})
