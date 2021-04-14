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
  watch,
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

    // canvas 导入图片的实际分辨率
    const canvasWidth = ref(1280)
    const canvasHeight = ref(720)
    // canvas 的引用
    const refCanvas = ref(null)
    // transfor ratio (canvasCoord 2 real-imageCanvas)
    let transforRatio = 1

    // cant use watch before the refCanvas create
    // watch(
    //   () => refCanvas.value.clientHeight,
    //   (oldValue, newValue) => {
    //     console.log(newValue)
    //   }
    // )

    // div 宽高获取测试
    // function getWH() {
    //   console.log(refCanvas.value.clientHeight)
    //   console.log(refCanvas.value.clientWidth)
    // }
    // onMounted(getWH)
    // console.log(window)

    // define canvas size(height and width of the canvas layer)
    // @param maxCanvasHeight
    // @param maxCanvasWidth
    const maxCanvasHeight = ref(0)
    const maxCanvasWidth = ref(0)

    // canvas size function
    // and calc the transfor ratio
    // from canvasCoord 2 real-imageCanvas
    // @warning the css of canvas' height and width must be scaled equally
    // use canvasCoord convert way: ratio * canvasCoord
    function updateCanvasSize() {
      maxCanvasHeight.value = refCanvas.value.clientHeight
      maxCanvasWidth.value = refCanvas.value.clientWidth
      console.log(
        '[minicap] resize -> %d, %d',
        maxCanvasHeight.value,
        maxCanvasWidth.value
      )
      // calc ratio
      transforRatio = canvasHeight.value / maxCanvasHeight.value
      console.log('[minicap] cuurrent transfor ratio: %f', transforRatio)
    }

    // add onresize callback function
    // @warning dont overwrite this callback on other place
    window.onresize = () => {
      updateCanvasSize()
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
      // context.moveTo(100, 100) //设置起点状态
      // context.lineTo(700, 700) //设置末端状态
      // context.lineWidth = 5 //设置线宽状态
      // context.strokeStyle = '#222' //设置线的颜色状态
      // context.stroke() //进行绘制
    })

    // exec the first resize
    // @warning this operation must run after got the canvas context
    onMounted(updateCanvasSize)

    // define the request body of touch order
    // @param contact multi-touch contact
    // @param x
    // @param y
    // @param ts timestamp
    // @param type touch type (enum tap/swipe/release)
    const EVENT_TAP = 0
    const EVENT_SWIPE = 1
    const EVENT_RELEASE = 2

    function onMove() {
      // console.log('move')
    }
    function onUp(e) {
      console.log(
        '[minicap] up -> %d, %d -> %d, %d',
        e.layerX,
        e.layerY,
        e.layerX * transforRatio,
        e.layerY * transforRatio
      )
    }
    function onDown(e) {
      console.log(
        '[minicap] down -> %d, %d -> %d, %d',
        e.layerX,
        e.layerY,
        e.layerX * transforRatio,
        e.layerY * transforRatio
      )

      let reqBody = {
        x: parseInt(e.layerX * transforRatio),
        y: parseInt(e.layerY * transforRatio),
        contact: 0,
        type: EVENT_TAP,
        ts: Date.parse(new Date()) / 1000,
      }

      console.log(reqBody)
    }

    // coord transfor from canvasCoord to real-imageCoord
    function coordTransfor(canvasCoordX, canvasCoordY) {}

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
      initWs('ws://127.0.0.1:8080/screen')
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
