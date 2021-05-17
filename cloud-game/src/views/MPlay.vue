<template>
  <div class="m-play">
    <div class="play-main" ref="refMain">
      <canvas
        ref="refCanvas"
        class="play-canvas"
        :height="canvasHeight"
        :width="canvasWidth"
        @touchmove="onMove"
        @touchstart="onStart"
        @touchend="onEnd"
      ></canvas>
    </div>
  </div>

  <van-loading
    v-if="loading"
    class="global-loading"
    type="spinner"
    color="#1989fa"
  />
</template>

<script>
import {
  defineComponent,
  onMounted,
  getCurrentInstance,
  ref,
  reactive,
  watch,
} from 'vue'

import { useStore } from 'vuex'
import { key } from '../store'
import { useRouter, useRoute } from 'vue-router'

export default defineComponent({
  components: {},
  setup() {
    const store = useStore(key)
    const router = useRouter()
    const route = useRoute()
    const currentInstance = getCurrentInstance()
    let refCanvas = ref(null)
    let refMain = ref(null)

    onMounted(() => {
      console.log('1', refCanvas.value.offsetWidth)
      console.log('2', refCanvas.value.offsetHeight)
    })

    let canvasContext = null
    const canvasWidth = ref(720)
    const canvasHeight = ref(1280)

    // define canvas size(height and width of the canvas layer)
    // @param maxCanvasHeight
    // @param maxCanvasWidth
    const maxCanvasHeight = ref(0)
    const maxCanvasWidth = ref(0)

    // transfor ratio (canvasCoord 2 real-imageCanvas)
    let transforRatio = 1

    // canvas size function
    // and calc the transfor ratio
    // from canvasCoord 2 real-imageCanvas
    // @warning the css of canvas' height and width must be scaled equally
    // use canvasCoord convert way: ratio * canvasCoord
    function updateCanvasSize() {
      console.log(refCanvas.value.clientHeight)
      maxCanvasHeight.value = refCanvas.value.clientHeight
      maxCanvasWidth.value = refCanvas.value.clientWidth
      console.log(
        '[minicap] resize -> %d, %d',
        maxCanvasHeight.value,
        maxCanvasWidth.value
      )
      // calc ratio
      transforRatio = canvasHeight.value / maxCanvasHeight.value
      console.log('[minicap] current transfor ratio: %f', transforRatio)
    }

    // add onresize callback function
    // @warning dont overwrite this callback on other place
    window.onresize = () => {
      updateCanvasSize()
    }

    // 隐藏nav tab
    store.commit('setMenuEnable', false)

    // 加载
    let loading = ref(true)

    function onMessage(msg) {
      let blob = new Blob([msg.data], { type: 'image/jpeg' })
      let img = new Image()
      img.onload = function () {
        canvasContext.drawImage(img, 0, -720)
      }
      let u = URL.createObjectURL(blob)
      img.src = u
    }
    let websock = null
    // 拉流 ws
    function initWs(wsuri) {
      websock = new WebSocket(wsuri)
      websock.onmessage = function (msg) {
        onMessage(msg)
      }
      websock.onclose = function (e) {
        console.log('[minicap] connection closed (' + e.code + ')')
      }
      websock.onopen = function () {
        console.log('[minicap] connected')
      }
      websock.onerror = function () {
        console.log('[minicap] error')
      }
    }
    function sendMessage(agentData) {
      websock.send(agentData)
    }

    onMounted(() => {
      // 获取上下文
      canvasContext = refCanvas.value.getContext('2d')
      // canvasContext.translate(200, 150)
      canvasContext.rotate((90 * Math.PI) / 180)

      // init ratio
      updateCanvasSize()

      // 设置监听函数
      currentInstance.appContext.config.globalProperties.sockets.onmessage = (
        res
      ) => {
        let result = JSON.parse(res.data)
        if (result.op === 6) {
          // 成功申请
          store.state.ws.send({ op: 5, data: route.params.aid })
          store.commit('setSession', result.stub)
          let wsi =
            import.meta.env['VITE_STREAMING_BASE_URL'] +
            'screen?session=' +
            result.stub
          initWs(wsi)
          setInterval(() => {
            loading.value = false
          }, 5000)
        }
      }

      // 请求机器
      currentInstance.appContext.config.globalProperties.$socket.send(
        JSON.stringify({
          op: 4,
        })
      )
    })

    function exit() {
      store.commit('setMenuEnable', false)
      router.back()
    }

    function ShortToUInt16(s) {
      var targets = []
      targets[0] = (s >> 8) & 0xff
      targets[1] = s & 0xff
      return targets
    }

    function U64(i) {
      var targets = []
      targets[3] = i & 0xff
      targets[2] = (i >> 8) & 0xff
      targets[1] = (i >> 16) & 0xff
      targets[0] = (i >> 24) & 0xff
      return targets
    }

    function onMove(e) {
      // console.log('move', event.touches[0].pageX, event.touches[0].pageY)
      // console.log('convert', convertCoordFromCanvas2Device(e))

      SendMove(convertCoordFromCanvas2Device(e))
    }
    function onStart(e) {
      // console.log('touch', event.touches[0].pageX, event.touches[0].pageY)
      // console.log('convert', convertCoordFromCanvas2Device(e))
      SendStart(convertCoordFromCanvas2Device(e))
    }
    function onEnd(e) {
      SendEnd({})
      // console.log('end')
    }

    function convertCoordFromCanvas2Device(event) {
      return {
        PointX: parseInt(event.touches[0].pageY * transforRatio),
        PointY: parseInt(
          (maxCanvasWidth.value - event.touches[0].pageX) * transforRatio
        ),
      }
    }

    function SendMove(ed) {
      let tag = [0x45, 0x56, 0x33]
      let contact = 1
      let type = 1
      let combine = (contact << 4) | type

      let x = ShortToUInt16(ed.PointX)
      let y = ShortToUInt16(ed.PointY)
      let ts = U64(1621189705)

      tag.push.apply(tag, [combine])
      tag.push.apply(tag, x)
      tag.push.apply(tag, y)
      tag.push.apply(tag, ts)

      let arrayBuffer = new Int8Array(tag).buffer
      sendMessage(arrayBuffer)
    }

    function SendStart(ed) {
      let tag = [0x45, 0x56, 0x33]
      let contact = 1
      let type = 0
      let combine = (contact << 4) | type

      let x = ShortToUInt16(ed.PointX)
      let y = ShortToUInt16(ed.PointY)
      let ts = U64(1621189705)

      tag.push.apply(tag, [combine])
      tag.push.apply(tag, x)
      tag.push.apply(tag, y)
      tag.push.apply(tag, ts)

      let arrayBuffer = new Int8Array(tag).buffer
      sendMessage(arrayBuffer)
    }

    function SendEnd(ed) {
      let tag = [0x45, 0x56, 0x33]
      let contact = 1
      let type = 2
      let combine = (contact << 4) | type

      let x = ShortToUInt16(66)
      let y = ShortToUInt16(66)
      let ts = U64(1621189705)

      tag.push.apply(tag, [combine])
      tag.push.apply(tag, x)
      tag.push.apply(tag, y)
      tag.push.apply(tag, ts)

      let arrayBuffer = new Int8Array(tag).buffer
      sendMessage(arrayBuffer)
    }
    return {
      exit,
      loading,
      refCanvas,
      canvasWidth,
      canvasHeight,
      onMove,
      onStart,
      onEnd,
      refMain,
    }
  },
})
</script>

<style>
body {
  /* 全局污染 白色主题 */
  background-color: orange !important;
}
#app {
  background-color: orange;
}

/* 上下左右居中 */
.global-loading {
  position: absolute;
  left: 50%;
  top: 50%;
  /* border: 1px solid #000; */
  transform: translate(-50%, -50%); /* 50%为自身尺寸的一半 */
}

.play-canvas {
  width: 100%;
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
}
</style>
