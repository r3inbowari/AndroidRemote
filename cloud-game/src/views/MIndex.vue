<template>
  <div style="padding-bottom: 50px">
    <router-view></router-view>
  </div>
  <van-tabbar @change="onChangeMenu" v-model="menuActive">
    <van-tabbar-item v-for="(item, index) in menuList.dat">
      <span>{{ item.name }}</span>
      <template #icon="props">
        <img :src="props.active ? item.active : item.inactive" />
      </template>
    </van-tabbar-item>
  </van-tabbar>
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
import { key } from '../../store'
import { useRouter } from 'vue-router'

import imgGameActive from '../assets/menu_game_active.png'
import imgGameInactive from '../assets/menu_game_inactive.png'
import imgComActive from '../assets/menu_com_active.png'
import imgComInactive from '../assets/menu_com_inactive.png'
import imgMeActive from '../assets/menu_me_active.png'
import imgMeInactive from '../assets/menu_me_inactive.png'

export default defineComponent({
  components: {},
  setup() {
    const router = useRouter()

    let menuActive = ref(0)
    let menuList = reactive({
      dat: [
        {
          active: imgGameActive,
          inactive: imgGameInactive,
          name: '游戏',
          route: 'MGame',
        },
        {
          active: imgComActive,
          inactive: imgComInactive,
          name: '社区',
          route: 'MCom',
        },
        {
          active: imgMeActive,
          inactive: imgMeInactive,
          name: '我的',
          route: 'MMe',
        },
      ],
    })

    function onChangeMenu(index) {
      router.push({
        name: menuList.dat[index].route,
      })
    }
    return {
      menuList,
      menuActive,
      onChangeMenu,
    }
  },
})
</script>

<style>
body {
  /* 全局污染 白色主题 */
  background-color: rgb(30, 40, 49) !important;
}
#app {
  background-color: rgb(30, 40, 49);
}
</style>

<style>
/* 主背景 */
.van-tabbar {
  background-color: rgb(30, 40, 49);
  box-shadow: 0 0 6px rgb(0 0 0 / 20%);
}
/* 活动背景 */
.van-tabbar-item--active {
  background-color: rgb(30, 40, 49);
}
/* 隐藏边框 */
[class*='van-hairline']::after {
  border-bottom: 0 solid #ebedf0;
  border-top: 0 solid #ebedf0;
}
</style>
