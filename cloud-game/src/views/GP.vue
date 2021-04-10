<template>
  <div class="game-play-window">
    <div class="play-toolbar-container mujin-drak"></div>
    <div class="play-main-container">
      <button style="display: none" @click="exitGameAsync">退出游戏</button>
      <div class="play-inner-content">
        <canvas
          ref="refCanvas"
          @mousemove="onMove"
          @mouseup="onUp"
          @mousedown="onDown"
          class="play-canvas"
          :width="canvasWidth"
          :height="canvasHeight"
          >12</canvas
        >
      </div>
    </div>
    <div class="play-symbol-use"></div>
    <div class="play-other"></div>
    <div class="user-track"></div>
    <div class="extend-adv"></div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, getCurrentInstance, ref } from 'vue'
import { useRouter } from 'vue-router'

import { key } from './store'
import { useStore } from 'vuex'

import { VueCookieNext } from 'vue-cookie-next'

import { IWebSocket } from './ws'

import { exitGame } from './game'

export default defineComponent({
  name: 'App',
  components: {},
  data() {
    return {}
  },
  setup() {
    const router = useRouter()

    // exit game trigger
    const exitGameAsync = async () => {
      exitGame(router)
    }

    const canvasWidth = ref(1280)
    const canvasHeight = ref(720)

    const refCanvas = ref(null)

    function doMouseDown() {
      console.log('as')
    }

    onMounted(() => {
      const context = refCanvas.value.getContext('2d')
      console.log(context)
      // context.canvas.addEventListener('mousedown', doMouseDown, false)
      // context.canvas.addEventListener('mousemove', doMouseDown, false)
    })

    function onMove() {
      console.log('move')
    }
    function onUp() {
      console.log('up')
    }
    function onDown() {
      console.log('down')
    }

    return {
      exitGameAsync,
      canvasWidth,
      canvasHeight,
      refCanvas,
      onMove,
      onUp,
      onDown,
    }
  },
  methods: {},
})
</script>

<style scoped>
/* condition */
.game-play-window {
  margin: 0;
  padding: 0;
}

/* put up the page */
.play-main-container {
  width: 100%;
  height: 100%;
  background: tomato;
  position: absolute;
}

/* inner content define */
.play-inner-content {
  padding: 10px;
}

.play-inner-content canvas {
  background-color: wheat;
  /* height: 100px; */
}
</style>
