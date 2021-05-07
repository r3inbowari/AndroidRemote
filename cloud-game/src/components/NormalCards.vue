<template>
  <ul class="list">
    <NormalCard :info="slideItems[0]"></NormalCard>
    <NormalCard></NormalCard>
  </ul>
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

import NormalCard from './NormalCard.vue'

export default defineComponent({
  name: 'NormalCards',
  components: { NormalCard },
  setup() {
    let slideItems = reactive([
      {
        title: '内核活动0',
        alias: 'kernel event',
        cover:
          'https://image-glb.qpyou.cn/hubweb/hive_img/web/banner/20200317/41a9e65126c0463ad975ac444027bfa3_1200x490.jpg',
        icon:
          'https://image-glb.qpyou.cn/hubweb/gmnotice/appcenter/1606972197901.jpg',
        type: 'RPG',
        tags: ['app', '最受喜爱'],
        desc: '精彩对决',
        update: true,
        aid: '6c1de797-ce90-43db-a51d-c77c440792f6',
      },
      {
        title: '内核活动1',
        alias: 'kernel event',
        cover:
          'https://image-glb.qpyou.cn/hubweb/hive_img/web/banner/20201203/00d2a88aeb6037a2ada2ffd1ad703e61_1200x490.jpg',
        icon:
          'https://image-glb.qpyou.cn/hubweb/gmnotice/appcenter/1560919617216.jpg',
        type: 'RPG',
        tags: ['app', '最受喜爱'],
      },
      {
        title: '内核活动2',
        alias: 'kernel event',
        cover:
          'https://image-glb.qpyou.cn/hubweb/hive_img/web/banner/20210204/0aa3bbb693b4d1b9cb6fc61ce9d9ac8e_1200x490.jpg',
        icon:
          'https://image-glb.qpyou.cn/hubweb/gmnotice/appcenter/1612412918869.jpg',
        type: 'RPG',
        tags: ['app', '最受喜爱'],
      },
    ])
    return { slideItems }
  },
})
</script>

<style scoped>
.list {
  margin-top: 30px;
}

* {
  margin: 0;
  padding: 0;
  vertical-align: baseline;
  border: 0;
  font-family: 'Noto Sans', sans-serif;
}
</style>
