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

// import vue-native-websocket-next
import VueNativeSock from 'vue-native-websocket-vue3'
// import store ext to vuesock
// import store from './store'

// import { Fabric } from 'vue-fabric';

import Vant from 'vant'
import 'vant/lib/index.css'
import { Toast } from 'vant'
import { Dialog } from 'vant'

// 桌面端适配
// import '@vant/touch-emulator'

const app = createApp(App)
app
  .use(router)
  .use(store, key)
  .use(ElementPlus)
  .use(Toast) // vant toast
  .use(Dialog)
  .use(Vant) // vant
  .use(VueCookieNext) // options api mount
  .use(VueNativeSock, import.meta.env['VITE_PUSH_BASE_URL'], {
    reconnection: false, // 自动重连
    reconnectionAttempts: 20, // 重连次数
    reconnectionDelay: 10 * 1000, // 重连间隔
    connectManually: true, // 启动手动连接模式
  })

app.mount('#app')
//
// console.log(app);
// export default app

// app.$cook = 'Vue3';

// set default config
// VueCookieNext.config({ expire: '7d' })
// set global cookie
// VueCookieNext.setCookie('theme', 'default')
// VueCookieNext.setCookie('hover-time', { expire: '1s' })
