import { createApp } from 'vue'
import App from './App.vue'
// 路由
import { router } from './router'
// vuex
import { key, store } from './store'
// element+
import ElementPlus from 'element-plus'
import 'element-plus/lib/theme-chalk/index.css'

// import global css
import './css/glob.css'
import './css/btn.css'

// ali iconfont
import './assets/icon/iconfont.css'

const app = createApp(App).use(router).use(store, key).use(ElementPlus)
app.mount('#app')
