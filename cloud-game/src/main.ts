import { createApp } from 'vue'
import App from './App.vue'
// router-next
import { router } from './router'
// vuex
import { store, key } from './store'
// element+
import ElementPlus from 'element-plus'
import 'element-plus/lib/theme-chalk/index.css'

// import global-css
import './css/glob.css'
import './css/btn.css'

// import ali-iconfont
import './assets/icon/iconfont.css'

// import vue-cookies
import { VueCookieNext } from 'vue-cookie-next'

const app = createApp(App)
  .use(router)
  .use(store, key)
  .use(ElementPlus)
  .use(VueCookieNext) // options api mount
app.mount('#app')

// app.$cook = 'Vue3';

// set default config
// VueCookieNext.config({ expire: '7d' })
// set global cookie
// VueCookieNext.setCookie('theme', 'default')
// VueCookieNext.setCookie('hover-time', { expire: '1s' })
