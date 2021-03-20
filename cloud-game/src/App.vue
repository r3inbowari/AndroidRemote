<template>
  <div>
    <img alt="Vue logo" src="./assets/logo.png" />
    <div>12</div>
  </div>
  <!-- <HelloWorld msg="Welcome to Your Vue.js + TypeScript App"/> -->
</template>

<script lang="ts">
// import { onMounted } from "vue";
import { Options, Vue } from 'vue-class-component'
// import HelloWorld from './components/HelloWorld.vue';

@Options({
  // components: {
  //   HelloWorld,
  // },
  // props: {
  //   title: String
  // }
})
export default class App extends Vue {
  // Class properties will be component data
  author: string = 'r3inbowari'
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
