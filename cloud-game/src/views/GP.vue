<template>
  <div
    class="game-play-window keyboard-body"
    ref="kbRef"
    @key-event="keyboardEvnetHandler"
  >
    <div class="play-toolbar-container mujin-drak"></div>
    <div class="play-main-container">
      <button style="display: none" @click="exitGameAsync">退出游戏</button>
      <div class="play-inner-content">
        <canvas
          @contextmenu.prevent=""
          disable-scroll="true"
          ref="refCanvas"
          @touchmove="onMove"
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

import { key } from '../store'
import { useStore } from 'vuex'

import { VueCookieNext } from 'vue-cookie-next'

import { IWebSocket } from './ws'

import { exitGame } from './game'
import { keyCodeMap } from '../utils'

export default defineComponent({
  name: 'App',
  components: {},
  methods: {
    show(event) {
      console.log(event)
    },
  },
  data() {
    return {}
  },
  setup() {
    const store = useStore(key)
    function stobs(str) {
      var ch,
        st,
        re = []
      for (var i = 0; i < str.length; i++) {
        ch = str.charCodeAt(i) // get char
        st = [] // set up "stack"
        do {
          st.push(ch & 0xff) // push byte to stack
          ch = ch >> 8 // shift value down by 1 byte
        } while (ch)
        // add stack contents to result
        // done because chars have "wrong" endianness
        re = re.concat(st.reverse())
      }
      // return an array of bytes
      return re
    }

    function ShortToUInt16(s) {
      var targets = []
      targets[0] = (s >> 8) & 0xff
      targets[1] = s & 0xff
      return targets
    }

    function U64(i) {
      // console.log(i)
      var targets = []
      // targets[0] = 0
      // targets[1] = 0
      // targets[2] = 0
      // targets[3] = 0
      // targets[7] = i & 0xff
      // targets[6] = (i >> 8) & 0xff
      // targets[5] = (i >> 16) & 0xff
      // targets[4] = (i >> 24) & 0xff

      targets[3] = i & 0xff
      targets[2] = (i >> 8) & 0xff
      targets[1] = (i >> 16) & 0xff
      targets[0] = (i >> 24) & 0xff
      return targets
    }

    onMounted(() => {
      // 0x45, 0x56, 0x45, 0x4E, 0X54,
      // 0x33,
      // 0x10,       // 0001 0000 -> Contact 1, Type,
      // 0x01, 0xF4, // 500 X
      // 0x01, 0xF4, // 500 Y
      // 0x00, 0x00, 0x00, 0x00, 0x60, 0x79, 0x86, 0xA6, // 1618577062 Ts
      // console.log('test')
      // let tag = [0x45, 0x56, 0x45, 0x4e, 0x54, 0x33]
      // let contact = 1
      // let type = 0
      // let combine = contact << 4 | type
      // let x = ShortToUInt16(500)
      // let y = ShortToUInt16(500)
      // let ts = U64(Date.parse(new Date()) / 1000)
      // let sb = stobs('a8f5f167f44f4964e6c998dee827110c')
      // tag.push.apply(tag, [combine])
      // tag.push.apply(tag, x)
      // tag.push.apply(tag, y)
      // tag.push.apply(tag, ts)
      // tag.push.apply(tag, sb)
      // console.log(tag)
      // --------------------
      // setTimeout(() => {
      //   console.log('move')
      //   let tag = [0x45, 0x56, 0x45, 0x4e, 0x54, 0x33]
      //   let contact = 1
      //   let type = 1
      //   let combine = (contact << 4) | type
      //   let x = ShortToUInt16(parseInt(500))
      //   let y = ShortToUInt16(parseInt(500))
      //   let ts = U64(Date.parse(new Date()) / 1000)
      //   let sb = stobs('a8f5f167f44f4964e6c998dee827110c')
      //   tag.push.apply(tag, [combine])
      //   tag.push.apply(tag, x)
      //   tag.push.apply(tag, y)
      //   tag.push.apply(tag, ts)
      //   tag.push.apply(tag, sb)
      //   let arrayBuffer = new Int8Array(tag).buffer
      //   websocketsend(arrayBuffer)
      // }, 5000)

      setTimeout(() => {
        console.log('keyboard event')
      }, 5000)
    })

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

    let isPress = false
    let cnt = 0
    function onMove(e) {
      cnt++
      if (cnt % 2 === 1 && isPress) {
        // console.log(
        //   '[minicap] move -> %d, %d -> %d, %d',
        //   e.layerX,
        //   e.layerY,
        //   e.layerX * transforRatio,
        //   e.layerY * transforRatio
        // )
        console.log('move')
        let tag = [0x45, 0x56, 0x33]
        let contact = 1
        let type = 1
        let combine = (contact << 4) | type
        let x = ShortToUInt16(parseInt(e.layerX * transforRatio))
        let y = ShortToUInt16(parseInt(e.layerY * transforRatio))
        let ts = U64(Date.parse(new Date()) / 1000)
        // let sb = stobs('a8f5f167f44f4964e6c998dee827110c')

        tag.push.apply(tag, [combine])
        tag.push.apply(tag, x)
        tag.push.apply(tag, y)
        tag.push.apply(tag, ts)
        // tag.push.apply(tag, sb)

        let arrayBuffer = new Int8Array(tag).buffer
        websocketsend(arrayBuffer)
      }
    }
    function onUp(e) {
      isPress = false
      console.log(
        '[minicap] up -> %d, %d -> %d, %d',
        e.layerX,
        e.layerY,
        e.layerX * transforRatio,
        e.layerY * transforRatio
      )

      // TODO: test1
      // websocketsend(
      //   JSON.stringify({
      //     id: 'a8f5f167f44f4964e6c998dee827110c',
      //     contact: 2,
      //     x: parseInt(e.layerX * transforRatio),
      //     y: parseInt(e.layerY * transforRatio),
      //     ts: Date.parse(new Date()) / 1000,
      //   })
      // )

      let tag = [0x45, 0x56, 0x33]
      let contact = 1
      let type = 2
      let combine = (contact << 4) | type
      let x = ShortToUInt16(parseInt(e.layerX * transforRatio))
      let y = ShortToUInt16(parseInt(e.layerY * transforRatio))
      let ts = U64(Date.parse(new Date()) / 1000)
      // let sb = stobs('a8f5f167f44f4964e6c998dee827110c')

      tag.push.apply(tag, [combine])
      tag.push.apply(tag, x)
      tag.push.apply(tag, y)
      tag.push.apply(tag, ts)
      // tag.push.apply(tag, sb)

      let arrayBuffer = new Int8Array(tag).buffer
      websocketsend(arrayBuffer)
    }
    function onDown(e) {
      isPress = true
      console.log(
        '[minicap] down -> %d, %d -> %d, %d',
        e.layerX,
        e.layerY,
        e.layerX * transforRatio,
        e.layerY * transforRatio
      )

      // TODO: test2
      // let p = JSON.stringify({
      //   id: 'a8f5f167f44f4964e6c998dee827110c',
      //   contact: 0,
      //   x: parseInt(e.layerX * transforRatio),
      //   y: parseInt(e.layerY * transforRatio),
      //   ts: Date.parse(new Date()) / 1000,
      // })

      let tag = [0x45, 0x56, 0x33]
      let contact = 1
      let type = 0
      let combine = (contact << 4) | type
      let x = ShortToUInt16(parseInt(e.layerX * transforRatio))
      let y = ShortToUInt16(parseInt(e.layerY * transforRatio))
      let ts = U64(Date.parse(new Date()) / 1000)
      // let sb = stobs('a8f5f167f44f4964e6c998dee827110c')

      tag.push.apply(tag, [combine])
      tag.push.apply(tag, x)
      tag.push.apply(tag, y)
      tag.push.apply(tag, ts)
      // tag.push.apply(tag, sb)

      let arrayBuffer = new Int8Array(tag).buffer
      websocketsend(arrayBuffer)
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
        // console.log(new Date().getTime())
      }
      let u = URL.createObjectURL(blob)
      img.src = u
    }

    //数据发送
    function websocketsend(agentData) {
      // websock.send(JSON.stringify(agentData))
      websock.send(agentData)
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
      let stub = store.state.session

      let wsi =
        import.meta.env['VITE_STREAMING_BASE_URL'] + 'screen?session=' + stub
      initWs(wsi)
    })

    // 转换与映射关系
    // 键盘事件 -> 点击事件
    // @param key    int           键位 keyCode 值
    // @param redef  bool          是否重定义
    // @param def    Coord [x, y]  重定义坐标对
    // 键盘事件 -> 本身
    const kbRef = ref(null)
    const eventConfigReactive = reactive({
      eventConfig: {},
    })
    function keyboardEvnetHandler(event) {
      console.log(event.keyCode)
      if (keyCodeMap.has(event.keyCode)) {
        // console.log(keyCodeMap.get(event.keyCode))
        // 发送测试
        let tag = [0x45, 0x56, 0x33]
        let contact = 0
        let type = 3 // 3 表示键盘事件
        let combine = (contact << 4) | type
        let x = ShortToUInt16(keyCodeMap.get(event.keyCode))
        let y = ShortToUInt16(0)
        let ts = U64(Date.parse(new Date()) / 1000)
        // let sb = stobs('a8f5f167f44f4964e6c998dee827110c')

        tag.push.apply(tag, [combine])
        tag.push.apply(tag, x)
        tag.push.apply(tag, y)
        tag.push.apply(tag, ts)
        // tag.push.apply(tag, sb)

        let arrayBuffer = new Int8Array(tag).buffer
        websocketsend(arrayBuffer)
      }
    }

    function keyboardEvnetInit() {
      console.log(keyCodeMap)
      // 配置加载
      eventConfigReactive.eventConfig = {
        aid: '6c1de797-ce90-43db-a51d-c77c440792f6', // 游戏id
        uid: 'dd33b246-ce90-43db-a51d-c77c440792f6', // 游戏用户id
        index: 0, // 配置索引(0-2)
        tables: new Map([[46, [300, 300]]]), // 拦截 46 并重定向到Touch 44
      }

      // 初始化
      // Vue.config.keyCodes 修改 code 映射参数
      console.log(kbRef)
      window.addEventListener('keydown', function (event) {
        // 使用标准事件插件自定义事件
        // Vue div中绑定的@key-event就会收到这个自定义事件
        var myEvent = new Event('key-event')
        myEvent.keyCode = event.keyCode
        // dispatch this event
        kbRef.value.dispatchEvent(myEvent)
      })
    }

    onMounted(keyboardEvnetInit)

    return {
      exitGameAsync,
      canvasWidth,
      canvasHeight,
      refCanvas,
      kbRef,
      onMove,
      onUp,
      onDown,
      keyboardEvnetHandler,
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
  /* background: tomato; */
  background: rgb(30, 30, 30);
  position: absolute;
}

/* inner content define */
.play-inner-content {
  padding: 0;
}

.play-inner-content canvas {
  /* background-color: wheat; */
  background-color: rgb(30, 30, 30);
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
