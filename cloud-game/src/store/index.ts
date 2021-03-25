import { State } from '@vue/runtime-core'
import { InjectionKey } from 'vue'
import { createStore, StoreOptions, Store } from 'vuex'

declare module '@vue/runtime-core' {
	// define your typings for the store state
	export interface State {
		count: number
	}

	interface ComponentCustomProperties {
		$store: Store<State>
	}	
}

// define injection key
export const key: InjectionKey<Store<State>> = Symbol()

export const store = createStore<State>({
  state: {
    count: 1
  },
  mutations: {
    increment (state) {
      // mutate state
      state.count++
    }
  }
})