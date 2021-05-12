<template>
  <div class="da-content-header">
    <el-row type="flex" class="header-row" justify="space-between">
      <el-col :span="6"><h2>用户状态</h2></el-col>
      <el-col :span="6">
        <div
          style="
            float: right;
            margin-right: 30px;
            margin-top: 10px;
            color: steelblue;
          "
        >
          更新时间 {{ freshTime }}
        </div>
      </el-col>
    </el-row>
  </div>
  <div class="da-content-main">
    <el-table
      empty-text="无数据"
      :data="wsSessionsData.dat"
      style="width: 100%"
    >
      <el-table-column prop="stub" label="存根" width="150"> </el-table-column>
      <el-table-column prop="uid" label="用户名" width="100"> </el-table-column>
      <el-table-column prop="created" label="创建时间" width="100">
      </el-table-column>
      <el-table-column prop="term" label="持续时间（分钟）" width="140">
      </el-table-column>
      <el-table-column prop="state" label="状态" width="100"> </el-table-column>
      <el-table-column prop="did" label="容器ID" width="200"> </el-table-column>
      <el-table-column prop="aid" label="游戏ID" width="200"> </el-table-column>
      <el-table-column prop="remote_address" label="远程地址">
      </el-table-column>
      <el-table-column label="操作">
        <template #default="scope">
          <el-button
            size="mini"
            type="danger"
            @click="handlekillSession(scope.$index, scope.row)"
            >强制关闭会话
          </el-button>
        </template>
      </el-table-column>
    </el-table>
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
import { ops, formatDate } from '../../utils'
import { useStore } from 'vuex'
import { key } from '../../store'

import { getSessions, killSession } from '../../api/session'

export default defineComponent({
  data() {
    return {}
  },
  components: {},
  setup() {
    const store = useStore(key) // 主题监听
    let customClass = ref('dark')

    // 响应式 ws sessions
    const wsSessionsData = reactive({
      dat: [],
      ts: '',
    })

    watch(
      () => store.state.theme,
      (oldValue, newValue) => {
        customClass.value = store.state.theme ? 'dark' : 'light'
      }
    )

    let freshTime = ref('')

    onMounted(() => {
      getSessions().then((res) => {
        wsSessionsData.dat = res.Data
        freshTime.value = formatDate()
      })

      setInterval(() => {
        getSessions().then((res) => {
          wsSessionsData.dat = res.Data
          freshTime.value = formatDate()
        })
      }, 5000)
    })

    let tableData = [
      {
        stub: '2016-05-02',
        uid: '王小虎',
        created: '早上',
        term: '58分钟',
        state: '未绑定',
        did: 'a23bd2ccfa23bd2ccf',
        aid: '未绑定',
        address: '192.168.5.233:23889',
      },
    ]

    function handlekillSession(index, row) {
      killSession(row.stub)
      wsSessionsData.dat.splice(index, 1)
    }
    return { freshTime, tableData, wsSessionsData, handlekillSession }
  },
})
</script>

<style>
.da-content-header {
  background-color: var(--color-bg-sidebar);
  padding: 10px;
  border-radius: 8px;
  margin-bottom: 10px;
}

.da-content-main {
  background-color: var(--color-bg-sidebar);
  padding: 10px;
  border-radius: 8px;
}

h2 {
  display: block;
  font-size: 1.5em;
  margin-block-start: 0.23em;
  margin-block-end: 0.23em;
  margin-inline-start: 0px;
  margin-inline-end: 0px;
  font-weight: bold;
}

.header-row {
  display: flex;
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
