<template>
  <div class="wait-queue-card">
    <el-dialog
      top="30vh"
      v-model="waitCardVisible"
      width="25%"
      :show-close="false"
      center
      :close-on-click-modal="false"
    >
      <template #title>
        <span class="dialog-title"> 游戏启动中 </span>
      </template>
      <div class="dialog-content">
        <div class="dialog-content-tip">
          <div><span></span></div>
          <div v-if="processValue === 0">
            <div>
              <span>服务器拥挤, 正在排队 {{ loadingDotStr }}</span>
            </div>
            <div>
              <span>前方还有 {{ personQueueSum }} 人, 请耐心等待</span>
            </div>
          </div>
          <div v-if="processValue === 1">
            <span>正在请求服务器资源 {{ loadingDotStr }}</span>
          </div>
          <div v-if="processValue === 2">
            <span>正在下载游戏依赖环境 {{ loadingDotStr }}</span>
          </div>
          <div v-if="processValue === 3">
            <span>游戏环境搭建中 {{ loadingDotStr }}</span>
          </div>
          <div v-if="processValue === 4">
            <span>正在做最后的准备, 请稍后 {{ loadingDotStr }}</span>
          </div>
          <div v-if="processValue === 5">
            <span>服务器拥挤, 正在排队 {{ loadingDotStr }}</span>
          </div>
        </div>
      </div>
      <template #footer>
        <span class="dialog-footer">
          <button @click="onRunClose" class="action-btn">取消</button>
        </span>
        <el-progress :percentage="openPercentage"> </el-progress>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import {
  defineComponent,
  onMounted,
  reactive,
  ref,
  getCurrentInstance,
} from 'vue'
import { useRouter } from 'vue-router'

import { key } from '../store'
import { useStore } from 'vuex'
import { startTimer } from '../utils'

export default defineComponent({
  name: 'App',
  data() {
    return {}
  },
  setup(props) {
    const router = useRouter()

    // 前方等待人数
    const personQueueSum = ref(0)

    // 加载小点点
    const loadingDotStr = ref('.')

    // 卡片开关
    const waitCardVisible = ref(false)

    // 游戏启动角度
    const openPercentage = ref(0)

    // 启动 tag 序号
    const processValue = ref(0)

    // 卡片close处理
    function handleClose() {
      console.log('close')
      waitCardVisible.value = false
    }

    const raceClk = ref(false)

    const store = useStore(key)

    const currentInstance = getCurrentInstance()

    // 加载过程模拟
    function handleRun() {
      currentInstance.appContext.config.globalProperties.sockets.onmessage = (
        res
      ) => {
        // msg case
        let result = JSON.parse(res.data)
        // console.log('[ws] ' + result.op)
        if (result.op === 6) {
          // open
          store.state.ws.send({ op: 5 })

          store.commit('setSession', result.stub)
          setInterval(() => {
            router.replace({
              name: 'Play',
            })
          }, 5000)
        }
      }

      // 请求机器
      currentInstance.appContext.config.globalProperties.$socket.send(
        JSON.stringify({
          op: 4,
        })
      )

      let timeout = 30

      let cnt = 10
      startTimer(cnt, () => {
        // console.log(cnt)
        cnt++
        openPercentage.value++
        dotdotdot()
        // console.log(raceClk.value)
        if (raceClk.value) {
          waitCardVisible.value = false
          openPercentage.value = 0
          raceClk.value = false
          return -1
        }
        if (openPercentage.value === 20) {
          processValue.value = 1
        }

        if (openPercentage.value === 40) {
          processValue.value = 2
        }

        if (openPercentage.value === 60) {
          processValue.value = 3
        }

        if (openPercentage.value === 80) {
          processValue.value = 4
        }

        if (openPercentage.value === 90) {
          processValue.value = 5
        }

        if (openPercentage.value === 100) {
          // router.replace({
          //   name: 'Play',
          // })
          return -1
        } else {
          return cnt
        }
      })

      function dotdotdot() {
        if (loadingDotStr.value === '.') {
          loadingDotStr.value = '..'
        } else if (loadingDotStr.value === '..') {
          loadingDotStr.value = '...'
        } else if (loadingDotStr.value === '...') {
          loadingDotStr.value = '.'
        }
      }

      // setTimeout(() => {
      //   store.state.ws.send({ msg: 'hello' })
      // }, 4000)
    }

    // onMounted(handleRun)

    function onRunClose() {
      console.log('[wait] close waiting window')
      raceClk.value = true
    }

    return {
      handleClose,
      waitCardVisible,
      openPercentage,
      personQueueSum,
      processValue,
      loadingDotStr,
      onRunClose,
      handleRun,
    }
  },
  methods: {
    openWait() {
      this.waitCardVisible = true
      this.handleRun()
    },
  },
})
</script>

<style>
.dialog-content {
  text-align: center;
}
/* 卡片全局 */
.wait-queue-card .el-dialog {
  background-color: rgba(7, 65, 99, 0.9);
  /* background-color: rgba(28, 32, 34, 0.9); */
  border-radius: 10px;
}

/* 标题颜色 */
.wait-queue-card .dialog-title {
  color: rgb(235, 239, 243);
  font-weight: 700;
  line-height: 20px;
  font-size: 18px;
}

.wait-queue-card .el-dialog__footer {
  /* text-align: center; */
}

/* 隐藏进度条text的占位 */
.wait-queue-card .el-progress__text {
  display: none;
}

/* 框fotter的padding */
.wait-queue-card .el-dialog__footer {
  padding: 0px;
}

/* 进度条bar的背景白色替换 */
.wait-queue-card .el-progress-bar__outer {
  background-color: rgba(28, 32, 40, 0.9);
}

/* 进度条 top margin */
.wait-queue-card .el-progress {
  margin-top: 20px;
}

.dialog-content-tip {
  color: rgb(156, 230, 20);
  font-weight: 600;
  line-height: 20px;
}
</style>
