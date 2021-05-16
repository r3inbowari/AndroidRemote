<template>
  <div class="m-about">
    <van-nav-bar fixed safe-area-inset-top title="关于">
      <template #left>
        <i
          @click="backMe"
          class="
            van-badge__wrapper
            van-icon van-icon-cross
            van-action-sheet__close
          "
        ></i>
      </template>
    </van-nav-bar>

    <!-- <div class="title"><h1>关于</h1></div> -->
    <div class="about-logo">
      <svg
        class="octicon octicon-mark-github v-align-middle"
        height="100"
        viewBox="0 0 16 16"
        version="1.1"
        width="100"
        aria-hidden="true"
      >
        <path
          fill-rule="evenodd"
          d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.013 8.013 0 0016 8c0-4.42-3.58-8-8-8z"
        ></path>
      </svg>
    </div>

    <div class="version">build on ab231c...c64232</div>
    <div class="version-num">v0.0.1</div>
    <div class="project-author">r3inbowari @2021</div>
    <a
      href="https://github.com/r3inbowari/AndroidRemote"
      class="project-source"
    >
      源代码
    </a>
    <a href="mailto:r3inbowari@gmail.com"> Contact </a>
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
import { getBanner } from '../api/banner'
import { getHot } from '../api/hot'
import { getUpdate } from '../api/update'

import { useStore } from 'vuex'
import { key } from '../store'
import { useRouter } from 'vue-router'

import { userLogin, userInfo } from '../api/user'

import { Toast } from 'vant'

import { VueCookieNext } from 'vue-cookie-next'

export default defineComponent({
  components: {},
  setup() {
    const store = useStore(key)
    const router = useRouter()

    function backMe() {
      router.push({
        name: 'MMe',
      })
    }

    return { backMe }
  },
})
</script>

<style scoped>
.m-about {
  text-align: center;
  color: white;
}

.project-author {
  margin-top: 20px;
  font-size: 20px;
}
.about-logo {
  margin-top: 60px;
}
.version {
  margin-top: 12px;
}
</style>

<style>
.m-about .van-nav-bar {
  background-color: rgb(30, 40, 49);
  box-shadow: 0 0 6px rgb(0 0 0 / 20%);
}

.m-about .van-nav-bar__title {
  color: #fff;
  font-weight: 600;
  font-size: 18px;
}

.m-about .van-action-sheet__close {
  top: 12px;
  left: 0px;
}
.m-about .van-action-sheet__close::before {
  color: #fff;
}
</style>
