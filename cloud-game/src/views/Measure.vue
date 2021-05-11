<template>
  <div class="measure">
    <div class="measure-content">
      <div v-for="(item, index) in MeasureItems.dat">
        <div @click="onOpen(index)" class="measure-item">
          <div class="item-title">{{ item.title }}</div>
        </div>
        <div class="dvi"></div>
      </div>
    </div>
  </div>

  <!-- 查看主页横幅 -->
  <el-dialog
    :custom-class="customClass"
    title="横幅列表"
    v-model="dialogBannerTableVisible"
  >
    <el-table :data="bannerData.dat">
      <el-table-column property="title" label="名称" width="150">
      </el-table-column>
      <el-table-column property="_id" label="ID 识别号" width="150">
      </el-table-column>
      <el-table-column property="alias" label="别名" width="150">
      </el-table-column>
      <el-table-column property="icon" label="App 图标资源" width="150">
      </el-table-column>
      <el-table-column property="cover" label="App 横幅资源" width="150">
      </el-table-column>
      <el-table-column label="操作">
        <template #default="scope">
          <el-button type="success" size="mini">编辑</el-button>
          <el-button
            size="mini"
            type="danger"
            @click="handleDeleteBanner(scope.$index, scope.row)"
            >删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
  </el-dialog>

  <!-- 新增横幅 -->
  <el-dialog
    width="600px"
    :custom-class="customClass"
    title="添加主页横幅"
    v-model="dialogAddBannerVisible"
  >
    <el-form
      label-position="right"
      label-width="80px"
      :model="formBannerData.dat"
      ref="formBannerRef"
    >
      <el-form-item label="游戏名称">
        <el-input
          placeholder="例如: 仙境传说RO: 守护永恒的爱"
          v-model="formBannerData.dat.title"
        ></el-input>
      </el-form-item>
      <el-form-item label="游戏别名">
        <el-input
          placeholder="例如: Ragnarok Mobile"
          v-model="formBannerData.dat.alias"
        ></el-input>
      </el-form-item>
      <el-form-item label="游戏图标">
        <el-input
          placeholder="图标的资源地址, 禁止跨域资源访问"
          v-model="formBannerData.dat.icon"
        ></el-input> </el-form-item
      >.
      <el-form-item label="横幅地址">
        <el-input
          placeholder="横幅的资源地址, 请勿使用重定向资源地址"
          v-model="formBannerData.dat.cover"
        ></el-input>
      </el-form-item>
      <el-button style="width: 560px" type="primary" @click="handleAddBanner()"
        >立即创建</el-button
      >
    </el-form>
  </el-dialog>

  <!-- 查看热门 -->
  <el-dialog
    width="1050px"
    :custom-class="customClass"
    title="热门列表"
    v-model="dialogHotTableVisible"
  >
    <el-table :data="hotData.dat">
      <el-table-column property="title" label="名称" width="100">
      </el-table-column>
      <el-table-column property="_id" label="ID 识别号" width="90">
      </el-table-column>
      <el-table-column property="aid" label="游戏 ID" width="100">
      </el-table-column>
      <el-table-column property="alias" label="别名" width="150">
      </el-table-column>
      <el-table-column property="icon" label="App 图标资源" width="150">
      </el-table-column>
      <el-table-column property="cover" label="App 横幅资源" width="150">
      </el-table-column>
      <el-table-column property="tags[0]" label="标签" width="50">
      </el-table-column>
      <el-table-column property="tags[1]" label="标签" width="50">
      </el-table-column>
      <el-table-column label="操作">
        <template #default="scope">
          <el-button type="success" size="mini">编辑</el-button>
          <el-button
            size="mini"
            type="danger"
            @click="handleDeleteBanner(scope.$index, scope.row)"
            >删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
  </el-dialog>

  <el-dialog
    width="1220px"
    :custom-class="customClass"
    title="最近更新列表"
    v-model="dialogUpdateTableVisible"
  >
    <el-table :data="updateData.dat">
      <el-table-column property="title" label="名称" width="100">
      </el-table-column>
      <el-table-column property="_id" label="ID 识别号" width="90">
      </el-table-column>
      <el-table-column property="aid" label="游戏 ID" width="100">
      </el-table-column>
      <el-table-column property="alias" label="别名" width="150">
      </el-table-column>
      <el-table-column property="icon" label="App 图标资源" width="150">
      </el-table-column>
      <el-table-column property="cover" label="App 横幅资源" width="150">
      </el-table-column>
      <el-table-column label="标签" width="140">
        <template #default="scope">
          <div class="name-wrapper">
            <el-tag v-if="scope.row.new" style="margin-left: 5px" size="medium"
              >最新
            </el-tag>
            <el-tag
              v-if="scope.row.popular"
              style="margin-left: 5px"
              size="medium"
            >
              人气
            </el-tag>
            <el-tag
              v-if="scope.row.recommend"
              style="margin-left: 5px"
              size="medium"
            >
              推荐
            </el-tag>
            <el-tag
              v-if="scope.row.update"
              style="margin-left: 5px"
              size="medium"
            >
              更新
            </el-tag>
          </div>
        </template>
      </el-table-column>
      <el-table-column property="desc" label="描述" width="150">
      </el-table-column>

      <el-table-column label="操作">
        <template #default="scope">
          <el-button type="success" size="mini">编辑</el-button>
          <el-button
            size="mini"
            type="danger"
            @click="handleDeleteBanner(scope.$index, scope.row)"
            >删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
  </el-dialog>
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

import { getBanner, delBanner, addBanner } from '../api/banner'
import { getHot, delHot, addHot } from '../api/hot'
import { getUpdate, delUpdate, addUpdate } from '../api/update'

import { ElNotification } from 'element-plus'
import { ops } from '../utils'

import { useStore } from 'vuex'
import { key } from '../store'

export default defineComponent({
  data() {
    return {}
  },
  setup() {
    const MeasureItems = reactive({
      dat: [
        {
          title: '横幅管理',
        },
        {
          title: '添加横幅',
        },
        {
          title: '热门管理',
        },
        {
          title: '添加热门',
        },
        {
          title: '最近更新管理',
        },
        {
          title: '添加更新内容',
        },
      ],
    })

    let dialogBannerTableVisible = ref(false)
    let dialogAddBannerVisible = ref(false)
    let dialogHotTableVisible = ref(false)
    let dialogAddHotVisible = ref(false)
    let dialogUpdateTableVisible = ref(false)
    let dialogAddUpdateVisible = ref(false)

    const bannerData = reactive({ dat: [] })
    const hotData = reactive({ dat: [] })
    const updateData = reactive({ dat: [] })

    let formBannerRef = ref(null)

    function onOpen(index) {
      console.log(index)
      if (index === 0) {
        // 获取banner
        getBanner()
          .then((res) => {
            bannerData.dat = res.Data
            dialogBannerTableVisible.value = true
          })
          .catch(() => {
            ops(
              'error',
              '消息 😊',
              '网络状态或服务器异常(3003)',
              customClass.value
            )
          })
      } else if (index === 1) {
        dialogAddBannerVisible.value = true
      } else if (index === 2) {
        // 获取hot
        getHot()
          .then((res) => {
            hotData.dat = res.Data
            dialogHotTableVisible.value = true
          })
          .catch(() => {
            ops(
              'error',
              '消息 😊',
              '网络状态或服务器异常(3003)',
              customClass.value
            )
          })
      } else if (index === 3) {
        dialogAddHotVisible.value = true
      } else if (index === 4) {
        // 获取update
        getUpdate()
          .then((res) => {
            updateData.dat = res.Data
            dialogUpdateTableVisible.value = true
          })
          .catch(() => {
            ops(
              'error',
              '消息 😊',
              '网络状态或服务器异常(3003)',
              customClass.value
            )
          })
      } else if (index === 5) {
        dialogAddUpdateVisible.value = true
      }
    }
    // get store using inject store key
    const store = useStore(key)

    // 主题更换 elementui
    let customClass = ref(store.state.theme ? 'dark' : 'light')

    // 主题监听
    watch(
      () => store.state.theme,
      (oldValue, newValue) => {
        customClass.value = store.state.theme ? 'dark' : 'light'
      }
    )

    // 删除banner
    function handleDeleteBanner(index) {
      // console.log(bannerData.dat[index]._id)
      delBanner(bannerData.dat[index]._id)
        .then(() => {
          ops('success', '消息 😊', '删除成功(3001)', customClass.value)
          // 刷新
          bannerData.dat.splice(index, 1)
        })
        .catch(() => {
          ops('error', '消息 😊', '删除失败(3002)', customClass.value)
        })
    }

    // 新增banner
    function handleAddBanner() {
      console.log(formBannerRef.value)
      if (
        formBannerData.dat.title === '' ||
        formBannerData.dat.alias === '' ||
        formBannerData.dat.cover === '' ||
        formBannerData.dat.icon === ''
      ) {
        ops('error', '消息 😊', '添加失败, 参数错误(3003)', customClass.value)

        formBannerRef.value.resetFields()
        return
      }
      addBanner(formBannerData.dat)
        .then(() => {
          ops('success', '消息 😊', '添加成功(3001)', customClass.value)
          formBannerRef.value.resetFields()
        })
        .catch(() => {
          ops('error', '消息 😊', '添加失败(3002)', customClass.value)
        })
    }
    const formBannerData = reactive({
      dat: {
        title: '',
        alias: '',
        cover: '',
        icon: '',
      },
    })

    return {
      handleDeleteBanner,
      customClass,
      MeasureItems,
      onOpen,
      formBannerData,
      handleAddBanner,
      formBannerRef,
      dialogBannerTableVisible,
      dialogAddBannerVisible,
      dialogHotTableVisible,
      dialogAddHotVisible,
      dialogUpdateTableVisible,
      dialogAddUpdateVisible,
      bannerData,
      hotData,
      updateData,
    }
  },
})
</script>

<style scoped>
*,
:after,
:before {
  box-sizing: border-box;
}

.measure-content {
  padding: 30px;
}
</style>

<style scoped>
.measure-item {
  height: 60px;
  width: 100%;
  /* background-color: var(--color-sb-active-row-bg); */
  background-color: var(--color-bg-sidebar);
  border-radius: 10px;
  display: flex;
  align-items: center;
  padding-left: 20px;
}

.measure-item:hover {
  background-color: var(--color-sb-active-row-bg);
}

.dvi {
  height: 1px;
  width: 100%;
  margin-top: 10px;
  margin-bottom: 10px;
  background-color: var(--color-bg-sidebar);
}

.item-title {
  font-size: 18px;
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

<!-- 用户信息框定制 -->
<style>
.el-alert,
.el-notification {
  display: -ms-flexbox;
  background-color: var(--color-bg-sidebar);
  border: 0px solid #ebeef5;
}

.el-notification__title {
  color: var(--color-text-highlight);
}

.el-notification__content p {
  color: var(--color-text-highlight);
}

.el-icon-close::before {
  color: var(--color-text-highlight);
}

/* 输入框定制 */
.el-input__inner {
  background-color: var(--color-input-bg);
  border: 1px solid var(--color-input-border);
  color: var(--color-text);
}
.el-input__inner::placeholder {
  color: #585c89;
}
.el-input__inner:hover {
  border-color: var(--color-toggle-bg);
}
.el-form-item__label {
  color: var(--color-text);
}
</style>
