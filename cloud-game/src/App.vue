<template>
  <!-- 路由出口 -->
  <!-- 路由匹配到的组件将渲染在这里 -->
  <router-view></router-view>

  <div v-if="dev" class="builtState">development mode v{{ version }}</div>
</template>

<script lang="ts">
import { defineComponent, onMounted } from 'vue'
import { printEnv } from './utils'
// import HomePage from './components/HelloWorld.vue'

import { key } from './store'
import { useStore } from 'vuex'

import { VueCookieNext } from 'vue-cookie-next'

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
      }
    }

    onMounted(initUserInfo)
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
