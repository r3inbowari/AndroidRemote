<template>
  <!-- 路由出口 -->
  <!-- 路由匹配到的组件将渲染在这里 -->
  <router-view></router-view>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
// import HomePage from './components/HelloWorld.vue'

export default defineComponent({
  name: 'App',
  components: {},
})
</script>
