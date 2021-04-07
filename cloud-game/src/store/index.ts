// import { State } from '@vue/runtime-core'
import { InjectionKey } from 'vue'
import { createStore, StoreOptions, Store } from 'vuex'

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
}

// define injection key
export const key: InjectionKey<Store<State>> = Symbol()

export const store = createStore<State>({
  state: {
    count: 1,
    avatar: 'null',
    token: '',
  },
  mutations: {
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
    }
  },
})
