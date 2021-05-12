<template>
  <div class="dashboard">
    <div>
      <el-row type="flex" class="row-bg">
        <el-col :span="8"
          ><PlayChart style="margin: 0 auto"></PlayChart
        ></el-col>
        <el-col :span="8"> <Server style="margin: 0 auto"></Server></el-col>
        <el-col :span="8"
          ><PlayChart style="margin: 0 auto"></PlayChart
        ></el-col>
      </el-row>
    </div>

    <div class="da-content"><Content></Content></div>
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

import { getBanner, delBanner, addBanner } from '../api/banner'
import { getHot, delHot, addHot } from '../api/hot'
import { getUpdate, delUpdate, addUpdate } from '../api/update'

import { ElNotification } from 'element-plus'
import { ops } from '../utils'

import { useStore } from 'vuex'
import { key } from '../store'

import PlayChart from '../components/Dash/PlayChart.vue'
import Server from '../components/Dash/Server.vue'
import Content from '../components/Dash/Content.vue'

export default defineComponent({
  data() {
    return {}
  },
  components: {
    PlayChart,
    Server,
    Content,
  },
  setup() {
    let customClass = ref('dark')
    const store = useStore(key)

    // 主题监听
    watch(
      () => store.state.theme,
      (oldValue, newValue) => {
        customClass.value = store.state.theme ? 'dark' : 'light'
      }
    )

    return {}
  },
})
</script>

<style scoped>
.dashboard {
  padding: 30px;
}

.da-content {
  margin-top: 20px;
}
</style>
