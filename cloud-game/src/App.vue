<template>
  <!-- 路由出口 -->
  <!-- 路由匹配到的组件将渲染在这里 -->
  <router-view></router-view>

  <div v-if="dev" class="builtState">development mode v{{ version }}</div>
</template>

<script lang="ts">
import { defineComponent, onMounted, getCurrentInstance } from 'vue'
import { printEnv } from './utils'
// import HomePage from './components/HelloWorld.vue'

import { key } from './store'
import { useStore } from 'vuex'

import { VueCookieNext } from 'vue-cookie-next'

import { userInfo } from './api/user'

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
            store.commit('setAvatar', res.data.avatar)
          })
          .catch((e) => {
            console.log(e)
          })
      }
    }

    function initMsgWebsocket() {
      const currentInstance = getCurrentInstance();
      // currentInstance?.appContext.config.globalProperties.$socket.send('ping');
      setTimeout(() => {
        currentInstance?.appContext.config.globalProperties.$socket.send('ping');
      }, 4000);
      console.log('12',currentInstance);
      
      (currentInstance?.appContext.config.globalProperties.sockets).onmessage = (res: {
        data: string
      }) => {
        console.log(res.data)
      }
    }

    onMounted(initUserInfo)
    onMounted(initMsgWebsocket)
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
