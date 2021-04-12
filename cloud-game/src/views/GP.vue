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
          :height="canvasHeight"
          :width="canvasWidth"
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
import {
  defineComponent,
  onMounted,
  getCurrentInstance,
  ref,
  reactive,
} from 'vue'
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

    let context = null
    let URL = null

    onMounted(() => {
      context = refCanvas.value.getContext('2d')
      // const context = refCanvas.value.getContext('2d')
      console.log('[minicap] getContext', context)
      // context.canvas.addEventListener('mousedown', doMouseDown, false)
      // context.canvas.addEventListener('mousemove', doMouseDown, false)

      // blob base url
      URL = window.URL || window.webkitURL
      // context.scale(1 / 2, 1 / 2)
      context.moveTo(100, 100) //设置起点状态
      context.lineTo(700, 700) //设置末端状态
      context.lineWidth = 5 //设置线宽状态
      context.strokeStyle = '#222' //设置线的颜色状态
      context.stroke() //进行绘制
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

    // data recv
    function websocketonmessage(data) {
      let blob = new Blob([data.data], { type: 'image/jpeg' })
      let img = new Image()
      img.onload = function () {
        context.drawImage(img, 0, 0)

        // img.onload = null
        // img = null
        // u = null
        // blob = null
      }
      let u = URL.createObjectURL(blob)
      img.src = u
    }

    //数据发送
    function websocketsend(agentData) {
      websock.send(JSON.stringify(agentData))
    }

    //关闭
    function websocketclose(e) {
      console.log('[minicap] connection closed (' + e.code + ')')
    }

    function websocketOpen(e) {
      console.log('[minicap] connected')
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
  padding: 0;
}

.play-inner-content canvas {
  background-color: wheat;
}

/* canvas 水平居中, 垂直居中暂未制作 */
.play-canvas {
  position: fixed;
  height: 100%;
  left: 0;
  right: 0;
  margin: 0 auto;
}
</style>
