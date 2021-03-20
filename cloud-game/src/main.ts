import { createApp } from 'vue'
import App from './App.vue'
import { key, store } from './store'
import { router } from './router'

const app = createApp(App)

  .use(router)
  .use(store, key)
app.mount('#app')
