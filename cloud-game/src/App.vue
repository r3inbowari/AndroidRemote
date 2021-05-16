<template>
  <!-- 路由出口 -->
  <!-- 路由匹配到的组件将渲染在这里 -->
  <router-view></router-view>

  <div v-if="dev" class="builtState">development mode v{{ version }}</div>
</template>

<script lang="ts">
import { defineComponent, onMounted, getCurrentInstance, ref } from 'vue'
import { printEnv } from './utils'
// import HomePage from './components/HelloWorld.vue'

import { key } from './store'
import { useStore } from 'vuex'

import { VueCookieNext } from 'vue-cookie-next'

import { userInfo } from './api/user'

import { IWebSocket } from './ws'
import { useRouter } from 'vue-router'

export default defineComponent({
  name: 'App',
  components: {},
  data() {
    return {
      dev: import.meta.env.DEV,
      version: import.meta.env['VITE_WEB_VERSION'],
    }
  },
  setup() {
    printEnv()

    const store = useStore(key)

    // load user info using localstorage or cookies
    function initUserInfo() {
      // 直接加载本地cookies map
      if (VueCookieNext.isCookieAvailable('token')) {
        let t = VueCookieNext.getCookie('token')
        store.commit('setToken', t)

        // pull user info online, no using localstorage
        userInfo()
          .then((res) => {
            // console.log(res)
            // save foreach
            // undefined avatar param
            store.commit('setInfo', res.data)
          })
          .catch(() => {
            console.log('[app] get info failed. see ↓')
          })
      }
    }

    // app root instance
    // this ws class has some error not be handled
    // but it can use
    // only one connection limited
    const currentInstance = getCurrentInstance()
    // ws instance
    let ws = new IWebSocket(currentInstance, null)
    store.commit('setWs', ws)
    const router = useRouter()

    function initMsgWebsocket() {
      ws.initMsgWebsocket()
    }
    onMounted(initUserInfo)
    // init ws
    onMounted(initMsgWebsocket)

    // 设备类型判断
    function isMobile() {
      let flag = navigator.userAgent.match(
        /(phone|pad|pod|iPhone|iPod|ios|iPad|Android|Mobile|BlackBerry|IEMobile|MQQBrowser|JUC|Fennec|wOSBrowser|BrowserNG|WebOS|Symbian|Windows Phone)/i
      )
      return flag
    }
    // 执行判断
    onMounted(() => {
      if (isMobile()) {
        // alert('移动端')
        // 跳转到移动端
        router.replace({
          name: 'MIndex',
        })
      } else {
        // router.replace({
        //   name: 'PCPage',
        // })
      }
    })
  },
})
</script>

<style scoped>
.builtState {
  position: fixed;
  float: left;
  bottom: 4px;
  right: 16px;

  font-size: 18px;
  color: rgb(255, 255, 255);
  font-weight: 600;
}
</style>
