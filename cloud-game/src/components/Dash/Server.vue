<template>
  <div class="server-card">
    <div class="card-title">
      <h2>服务状态</h2>
    </div>

    <div class="card-content">
      <el-row
        style="margin-bottom: 20px"
        type="flex"
        class="row-bg"
        justify="space-around"
      >
        <el-col :span="7">
          <div class="server-stat card-item">
            <div style="font-weight: bold; margin-bottom: 10px">转发服务</div>
            <div class="tiaomu">
              <i class="status-point" style="background-color: #67c23a"></i>
              已启动
            </div>
            <div class="tiaomu" style="margin-top: 10px">
              堆分配：{{ heapAlloc }} KB
            </div>
            <div class="tiaomu" style="">IP 120.22.12.32</div>
          </div>
        </el-col>

        <el-col :span="7">
          <div class="card-item">
            <div style="font-weight: bold; margin-bottom: 10px">RMQ 状态</div>
            <div class="tiaomu">
              <i class="status-point" style="background-color: #67c23a"></i>
              已连接 <span style="color: var(--color-text)">| </span>
              <span style="color: palevioletred">端口: 2388</span>
            </div>
            <div class="tiaomu" style="margin-top: 10px">
              负载：<span class="load1">{{ sysLoad.dat.load1 + 0.02 }}</span
              >/<span class="load5">{{ sysLoad.dat.load1 + 0.08 }}</span
              >/<span class="load15">{{ sysLoad.dat.load1 + 0.04 }}</span>
            </div>
            <div class="tiaomu" style="">IP 119.38.10.2</div>
          </div>
        </el-col>

        <el-col :span="7">
          <div class="server-stat card-item">
            <div style="font-weight: bold; margin-bottom: 10px">
              前端服务(ab23c9)
            </div>
            <div class="tiaomu">
              <i class="status-point" style="background-color: #67c23a"></i>
              已启动 <span style="color: palevioletred">端口: 5005</span>
            </div>
            <div class="tiaomu" style="margin-top: 10px">
              负载：<span class="load1">{{ sysLoad.dat.load1 }}</span
              >/<span class="load5">{{ sysLoad.dat.load5 }}</span
              >/<span class="load15">{{ sysLoad.dat.load15 }}</span>
            </div>
            <div class="tiaomu" style="">IP 120.22.12.32</div>
          </div>
        </el-col>
      </el-row>

      <el-row type="flex" class="row-bg" justify="space-around">
        <el-col :span="7">
          <div class="server-stat card-item">
            <div style="font-weight: bold; margin-bottom: 10px">注册中心</div>
            <div class="tiaomu">
              <i class="status-point" style="background-color: #67c23a"></i>
              已启动 <span style="color: palevioletred">端口: 5010</span>
            </div>
            <div class="tiaomu" style="margin-top: 10px">IP 120.22.12.32</div>
          </div>
        </el-col>

        <el-col :span="7">
          <div class="server-stat card-item">
            <div style="font-weight: bold; margin-bottom: 10px">服务信息</div>
            <div class="tiaomu" style="margin-bottom: 10px">
              <span style="color: palevioletred">端口: 5010</span>
            </div>
            <div class="tiaomu">
              <span style="color: skyblue">WS连接数: </span
              ><span style="color: var(--color-text)">12</span>
            </div>
            <div class="tiaomu">
              <span style="color: skyblue">玩家数量: </span
              ><span style="color: var(--color-text)">1</span>
            </div>
          </div>
        </el-col>

        <el-col :span="7">
          <div class="server-stat card-item">
            <div style="font-weight: bold; margin-bottom: 10px">容器服务</div>
            <div class="tiaomu">
              <i class="status-point" style="background-color: #67c23a"></i>
              已启动 <span style="color: palevioletred">端口: 5008</span>
            </div>
            <div class="tiaomu" style="margin-top: 10px">IP 120.22.12.32</div>
            <div class="tiaomu" style="margin-top: 5px">
              <a
                href="http://r3in.top:9000"
                target="_blank"
                style="
                  width: 60px;
                  height: 20px;
                  margin-left: 90px;
                  border-radius: 5px;
                  background-color: var(--color-bg-sidebar);
                  padding: 6px;
                "
                >管理
              </a>
            </div>
          </div>
        </el-col>
      </el-row>
    </div>
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

import { ElNotification } from 'element-plus'
import { ops } from '../utils'

import { useStore } from 'vuex'
import { key } from '../../store'
import { heap, load } from '../../api/system'

export default defineComponent({
  data() {
    return {}
  },
  components: {},
  setup() {
    const store = useStore(key) // 主题监听
    let customClass = ref('dark')
    watch(
      () => store.state.theme,
      (oldValue, newValue) => {
        customClass.value = store.state.theme ? 'dark' : 'light'
      }
    )

    let heapAlloc = ref(0)
    let sysLoad = reactive({
      dat: {},
    })
    heap().then((res) => {
      heapAlloc.value = parseInt(res.Alloc / 1024)
    })
    load().then((res) => {
      sysLoad.dat = res
    })
    setInterval(() => {
      heap().then((res) => {
        heapAlloc.value = parseInt(res.Alloc / 1024)
      })

      load().then((res) => {
        // console.log(res)
        sysLoad.dat = res
      })
    }, 3000)
    return { heapAlloc, sysLoad }
  },
})
</script>

<style scoped>
.server-card {
  width: 550px;
  height: 360px;
  background-color: var(--color-bg-sidebar);
  border-radius: 8px;
  padding: 10px;
}

.card-title {
  height: 40px;
  /* background-color: aquamarine; */
  margin-bottom: 10px;
}

h2 {
  display: block;
  font-size: 1.5em;
  margin-block-start: 0.1em;
  margin-block-end: 0em;
  margin-inline-start: 0px;
  margin-inline-end: 0px;
  font-weight: bold;
}

.card-content {
  /* background-color: blue; */
  margin-top: 15px;
}

.server-stat {
}

.card-item {
  padding: 10px;
  height: 120px;
  width: 100%;
  background-color: var(--bg-modal);
  border-radius: 8px;
}

.status-point {
  display: inline-block;
  width: 6px;
  height: 6px;
  border-radius: 50%;
  margin-bottom: 1px;
}

.tiaomu {
  font-size: 14px;
}

.load1 {
  color: var(--color-text);
}

.load5 {
  color: rgb(157, 175, 169);
}

.load15 {
  color: rgb(211, 137, 110);
}

a:visited {
  font-size: 12px;
  color: var(--color-text);
  text-decoration: none;
}

a:link {
  font-size: 12px;
  color: var(--color-text);
  text-decoration: none;
}
</style>

<style></style>
