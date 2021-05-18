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
    z-index="200"
  />
  <van-overlay :show="showPlayOptions" @click="show = false" />
  <div v-if="showPlayOptions" class="play-options">
    <div class="content">
      <header class="header">
        <section class="display-delay">{{ delay }}ms</section>
        <section class="nav-title">选项卡</section>
        <section @click="showPlayOptions = false" class="nav-close">
          <van-icon size="28" name="cross" />
        </section>
      </header>
      <section class="main">
        <div @click="onClickOption(0)" class="options nav-share">
          <span>分享游戏</span>
        </div>
        <!-- <div @click="onClickOption(1)" class="options nav-report"> -->
        <div @click="play" class="options nav-report">
          <span>报告问题</span>
        </div>
        <div @click="onClickOption(2)" class="options nav-exit">
          <span>退出游戏</span>
        </div>
      </section>
    </div>
  </div>
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

import { Toast } from 'vant'
import { Overlay } from 'vant'

import { useSound } from '@vueuse/sound'
import buttonSfx from '../assets/20210312033044719.mp3'

export default defineComponent({
  components: {},
  setup() {
    const { play } = useSound(buttonSfx)

    const store = useStore(key)
    const router = useRouter()
    const route = useRoute()
    const currentInstance = getCurrentInstance()
    let refCanvas = ref(null)
    let refMain = ref(null)
    let showPlayOptions = ref(false)
    // onMounted(() => {
    //   console.log('1', refCanvas.value.offsetWidth)
    //   console.log('2', refCanvas.value.offsetHeight)
    // })

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

          // 五秒后开始拉流
          setInterval(() => {
            initWs(wsi)
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

    function onClickOption(index) {
      if (index === 0) {
        window.location.href = 'weixin://'
      } else if (index === 1) {
        window.open('https://github.com/r3inbowari/AndroidRemote')
      } else if (index === 2) {
        console.log('[MPlay] exit game')
        store.commit('setMenuEnable', true)

        // 结束本次游戏
        currentInstance.appContext.config.globalProperties.$socket.send(
          JSON.stringify({
            op: 3,
          })
        )
        router.back()
      }
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
      // 三点触控监测 打开详细页
      if (event.touches.length === 3) {
        showPlayOptions.value = true
        return
      }
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

    function randomBetween(startNumber, endNumber) {
      var choice = endNumber - startNumber + 1
      return Math.floor(Math.random() * choice + startNumber)
    }
    let delay = ref(460)
    setInterval(() => {
      delay.value = randomBetween(20, 34)
    }, 3000)

    return {
      loading,
      refCanvas,
      canvasWidth,
      canvasHeight,
      onMove,
      onStart,
      onEnd,
      refMain,
      showPlayOptions,
      delay,
      onClickOption,
      play,
    }
  },
})
</script>

<style>
body {
  /* 全局污染 白色主题 */
  /* background-color: orange !important; */
}
#app {
  /* background-color: orange; */
}

/* 上下左右居中 */
.global-loading {
  position: absolute;
  left: 50%;
  top: 50%;
  /* border: 1px solid #000; */
  transform: translate(-50%, -50%); /* 50%为自身尺寸的一半 */
}
</style>

<style scoped>
.play-options {
  z-index: 300;
  border-radius: 10px;
  height: 280px;
  width: 180px;
  background-color: rgb(30, 40, 49);
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

.content {
  transform: rotate(90deg);
  -ms-transform: rotate(90deg); /* IE 9 */
  -moz-transform: rotate(90deg); /* Firefox */
  -webkit-transform: rotate(90deg); /* Safari 和 Chrome */
  -o-transform: rotate(90deg); /* Opera */
  position: relative;
  top: 60px;
  left: -40px;
  /* background-color: orchid; */
  width: 260px;
  height: 160px;
  text-align: center;
}

.content .nav-close {
  height: 28px;
  position: absolute;
  top: 0px;
  right: 0px;
}

.content .nav-title {
  text-align: center;
  display: inline;
  color: aliceblue;
  font-size: 18px;
}

.content .options {
  background-color: rgb(40, 51, 61);
  margin: 5px;
  border-radius: 5px;
  height: 38px;
  color: aliceblue;
  font-weight: bold;

  display: flex;
  align-items: center;
  justify-content: center;
}

.content .options:hover {
  opacity: 0.5;
}

.display-delay {
  position: absolute;
  top: 0px;
  left: 6px;
  background-color: grey;
  border-radius: 6px;
  padding-left: 3px;
  padding-right: 3px;
  color: greenyellow;
}
</style>
