<template>
  <!-- 路由出口 -->
  <!-- 路由匹配到的组件将渲染在这里 -->
  <router-view></router-view>

  <div v-if="dev" class="builtState">development mode v{{ version }}</div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { printEnv } from './utils'
// import HomePage from './components/HelloWorld.vue'

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
