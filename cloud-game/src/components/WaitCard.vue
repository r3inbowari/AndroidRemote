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
        <span class="dialog-title"> 启动中 </span>
      </template>
      <div class="dialog-content">
        <div><span></span></div>
        <div><span>服务器拥挤, 正在排队</span></div>
        <div>
          <span>前方还有 {{ personQueueSum }} 人, 请耐心等待</span>
        </div>
        <div><span>游戏搭载中</span></div>
        <div><span>正在下载游戏依赖环境...</span></div>
        <div><span>游戏环境搭建中...</span></div>
        <div><span>正在做最后的准备, 请稍后...</span></div>
      </div>
      <template #footer>
        <span class="dialog-footer">
          <button @click="onRunOpen" class="action-btn">打开游戏</button>
        </span>
        <el-progress :percentage="openPercentage"> </el-progress>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import { defineComponent, onMounted, ref } from 'vue'

export default defineComponent({
  name: 'App',
  data() {
    return {}
  },
  setup(props) {
    // 前方等待人数
    const personQueueSum = ref(0)

    // 加载小点点
    const loadingDotStr = ref('.')

    // 卡片开关
    const waitCardVisible = ref(false)

    // 游戏启动角度
    const openPercentage = ref(0)

    // 卡片close处理
    function handleClose() {
      console.log('close')
      waitCardVisible.value = false
    }

    // 加载过程模拟
    function handleRun() {
      let tag = setInterval(() => {
        openPercentage.value++
        if (openPercentage.value === 100) {
          clearInterval(tag)
        }
      }, 30)
    }

    onMounted(handleRun)

    return {
      handleClose,
      waitCardVisible,
      openPercentage,
      personQueueSum,
    }
  },
  methods: {
    openWait() {
      console.log('hello')
      this.waitCardVisible = true
    },
    onRunOpen() {},
  },
})
</script>

<style>
.dialog-content {
  text-align: center;
}
/* 卡片全局 */
.wait-queue-card .el-dialog {
  background-color: rgba(28, 32, 34, 0.9);
  border-radius: 10px;
}

/* 标题颜色 */
.wait-queue-card .dialog-title {
  color: aliceblue;
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
</style>
