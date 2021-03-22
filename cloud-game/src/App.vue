<template>
  <div id="app">
    <!-- 路由出口 -->
    <!-- 路由匹配到的组件将渲染在这里 -->
    <router-view></router-view>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
// import HomePage from './components/HelloWorld.vue'

export default defineComponent({
  name: 'App',
  components: {},
})
</script>

<style>
#app {
  font-family: -apple-system, BlinkMacSystemFont, Helvetica Neue, Helvetica,
    Arial, PingFang SC, Hiragino Sans GB, Microsoft YaHei, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: rgb(28, 32, 34);
  background-color: aqua;
}
::-webkit-scrollbar {
  /*滚动条整体样式*/
  width: 10px; /*高宽分别对应横竖滚动条的尺寸*/
  height: 2px;
}
::-webkit-scrollbar-thumb {
  /*滚动条里面小方块*/
  border-radius: 4px;
  /* box-shadow: inset 0 0 5px rgba(0, 0, 0, 0.2); */
  background: hsla(0, 0%, 100%, 0.1);
}
::-webkit-scrollbar-track {
  /*滚动条里面轨道*/
  /* box-shadow: inset 0 0 5px rgba(0, 0, 0, 0.2);
  border-radius: 0px;
  background: #ededed; */
  background: transparent;
}

body {
  margin: 0px;
  /* bg scroll  */
  background-color: rgb(28, 32, 34);
}
</style>
