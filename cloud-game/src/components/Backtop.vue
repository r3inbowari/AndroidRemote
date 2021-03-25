<template>
  <div class="v-float">
    <div @click="onTop" class="ele-backtop" style="right: 50px; bottom: 50px">
      <i class="el-icon-caret-top"></i>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted } from 'vue'
import { printEnv } from './utils'
import ItemCard from './ItemCard.vue'

// import HomePage from './components/HelloWorld.vue'

export default defineComponent({
  name: 'App',
  data() {
    return {}
  },
  setup() {},
  methods: {
    onTop() {
      ;(function smoothscroll() {
        var currentScroll =
          document.documentElement.scrollTop || document.body.scrollTop
        if (currentScroll > 0) {
          window.requestAnimationFrame(smoothscroll)
          window.scrollTo(0, currentScroll - currentScroll / 5)
        }
      })()
    },
  },
})
</script>

<style>
.ele-backtop {
  position: fixed;
  background-color: rgb(22, 22, 22);
  width: 40px;
  height: 40px;
  border-radius: 50%;
  color: #00a1d5;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  box-shadow: 0 0 6px rgb(0 0 0 / 12%);
  cursor: pointer;
  z-index: 5;
}
</style>
