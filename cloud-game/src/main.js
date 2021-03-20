import { createApp } from 'vue'
import App from './App.vue'
import { createStore } from 'vuex'
import { router } from "./router"

const store = createStore({
	state() {
		return {
			count: 0
		}
	},
	mutations: {
		increment(state) {
			state.count++
		}
	}
})

const app = createApp(App)
	.use(store)
	.use(router)
app.mount('#app')
