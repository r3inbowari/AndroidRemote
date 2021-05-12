<template>
  <div class="container-management">
    <header class="grid container-header">
      <div class="title"><h2>容器数据列表</h2></div>
    </header>
    <div style="margin-top: 10px" class="container-main grid">
      <el-table
        empty-text="暂无容器数据"
        :data="DeviceData.dat"
        style="width: 100%"
      >
        <el-table-column prop="deviceID" label="设备 ID" width="200">
        </el-table-column>
        <el-table-column prop="deviceBrand" label="类型" width="100">
        </el-table-column>
        <el-table-column prop="deviceContainerID" label="容器 ID" width="150">
        </el-table-column>
        <el-table-column prop="deviceIMEI" label="IMEI" width="150">
        </el-table-column>
        <el-table-column prop="deviceMAC" label="MAC" width="150">
        </el-table-column>
        <el-table-column prop="deviceRelease" label="ABI" width="80">
        </el-table-column>
        <el-table-column prop="deviceSDK" label="Level" width="80">
        </el-table-column>
        <el-table-column prop="deviceWidth" label="X轴" width="80">
        </el-table-column>
        <el-table-column prop="deviceHeight" label="Y轴" width="80">
        </el-table-column>
        <el-table-column prop="deviceTotalMem" label="虚拟内存(MB)" width="120">
        </el-table-column>
        <el-table-column prop="deviceAvailMem" label="可用内存(MB)" width="120">
        </el-table-column>
      </el-table>
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
import { key } from '../store'
import { getDeviceCount, getDeviceAll } from '../api/device'

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

    const DeviceData = reactive({ dat: [] })

    getDeviceAll().then((res) => {
      res.forEach((element, index) => {
        res[index].deviceAvailMem = parseInt(
          element.deviceAvailMem / 1024 / 1024
        )
        res[index].deviceTotalMem = parseInt(
          element.deviceTotalMem / 1024 / 1024
        )
      })
      DeviceData.dat = res
    })

    return { DeviceData }
  },
})
</script>

<style scoped>
.container-management {
  padding: 30px;
}

.container-management .grid {
  background-color: var(--color-bg-sidebar);
  border-radius: 8px;
  padding: 5px 15px 5px 15px;
}
</style>

<!-- 对话框与列表的主题定制 -->
<style>
/* 对话框颜色 */
.el-dialog {
  background-color: var(--color-bg-sidebar);
}

/* 表格底色 */
.el-table th,
.el-table tr {
  background-color: var(--color-bg-sidebar);
}

/* 表格底色扩展 */
.el-table,
.el-table__expanded-cell {
  background-color: var(--color-bg-sidebar);
}

/* 表格项悬浮 */
.el-table--enable-row-hover .el-table__body tr:hover > td {
  background-color: var(--color-input-border);
}

/* 表格头颜色 */
.el-table th,
.el-table tr:hover {
  background-color: var(--color-sb-active-row-bg);
}

/* 表格间隔分离 */
.el-table td,
.el-table th.is-leaf {
  border-bottom: 1px solid var(--color-separator);
}

/* 表格底部分离 */
.el-table--border::after,
.el-table--group::after,
.el-table::before {
  background-color: var(--color-separator);
}

/* 字体 */
.el-table .cell,
.el-dialog__title {
  color: var(--color-text);
}

/* 弹窗标题字重 */
.el-dialog__title {
  font-weight: bold;
}
</style>
