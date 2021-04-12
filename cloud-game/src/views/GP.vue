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

    // canvas 实际的分辨率
    const canvasWidth = ref(1280)
    const canvasHeight = ref(720)
    // canvas 的引用
    const refCanvas = ref(null)

    function doMouseDown() {
      console.log('as')
    }

    onMounted(() => {
      const context = refCanvas.value.getContext('2d')
      // const context = refCanvas.value.getContext('2d')
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

    let websock = null
    let global_callback = null

    function initWs(wsuri) {
      //ws地址
      // var wsuri = 'ws://' + getWebIP() + ':' + serverPort
      // var wsuri = 'ws://127.0.0.1:8080/ws'
      websock = new WebSocket(wsuri)
      websock.onmessage = function (e) {
        websocketonmessage(e)
      }
      websock.onclose = function (e) {
        websocketclose(e)
      }
      websock.onopen = function () {
        websocketOpen()
      }

      //连接发生错误的回调方法
      websock.onerror = function () {
        console.log('WebSocket连接发生错误')
      }
    }

    let recCnt = 0
    // data recv
    function websocketonmessage(data) {
      // global_callback(JSON.parse(e.data))
      // recCnt++
      // console.log(recCnt)
      const context = refCanvas.value.getContext('2d')
      let blob = new Blob([data.data], { type: 'image/jpeg' })
      let URL = window.URL || window.webkitURL
      let img = new Image()
      console.log(data)
      img.onload = function () {
        context.drawImage(img, 0, 0)

        // img.onload = null
        // img = null
        // u = null
        // blob = null
      }

      let u = URL.createObjectURL(blob)
      img.src = u

      // console.log(context)
    }

    //数据发送
    function websocketsend(agentData) {
      websock.send(JSON.stringify(agentData))
    }

    //关闭
    function websocketclose(e) {
      console.log('connection closed (' + e.code + ')')
    }

    function websocketOpen(e) {
      console.log('连接成功')
    }

    function getWebIP() {
      var curIP = window.location.hostname
      return curIP
    }

    onMounted(() => {
      initWs('ws://127.0.0.1:8080/ws')
    })

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
