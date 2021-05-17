<template>
  <div class="m-com">
    <van-nav-bar fixed safe-area-inset-top title="视频社区"></van-nav-bar>
  </div>

  <div class="m-com-main">
    <div class="video-list">
      <div style="border-radius: 10px">
        <iframe
          class="video-item"
          v-for="item in videoSet.dat"
          :src="item.src"
          allowfullscreen="allowfullscreen"
          width="100%"
          height="180"
          scrolling="no"
          frameborder="0"
          sandbox="allow-top-navigation allow-same-origin allow-forms allow-scripts"
        ></iframe>
      </div>
    </div>
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

    const videoSet = reactive({
      dat: [
        {
          src: '//player.bilibili.com/player.html?cid=339650270&aid=418002864&page=1&danmaku=0',
          title: 'wode ',
        },
        {
          src: '//player.bilibili.com/player.html?cid=339605804&aid=760618518&page=1&danmaku=0',
          title: 'wode ',
        },
        {
          src: '//player.bilibili.com/player.html?cid=339615740&aid=888072531&page=1&danmaku=0',
          title: 'wode ',
        },
        {
          src: '//player.bilibili.com/player.html?cid=270447073&aid=500824489&page=1&danmaku=0',
          title: 'wode ',
        },

        {
          src: '//player.bilibili.com/player.html?cid=339610711&aid=545601852&page=1&danmaku=0',
          title: 'wode ',
        },
        {
          src: '//player.bilibili.com/player.html?cid=339468838&aid=715562742&page=1&danmaku=0',
          title: 'wode ',
        },
        {
          src: '//player.bilibili.com/player.html?cid=339646712&aid=673019828&page=1&danmaku=0',
          title: 'wode ',
        },
      ],
    })

    return { videoSet }
  },
})
</script>

<style>
.m-com .van-nav-bar {
  background-color: rgb(30, 40, 49);
  box-shadow: 0 0 6px rgb(0 0 0 / 20%);
}

.m-com .van-nav-bar__title {
  color: #fff;
  font-weight: 600;
  font-size: 18px;
}

.m-com .van-action-sheet__close {
  top: 12px;
  left: 0px;
}
.m-com .van-action-sheet__close::before {
  color: #fff;
}
</style>

<style scoped>
.m-com-main {
  margin-top: 60px;
  padding: 0 16px 10px 16px;
}

.video-item {
  margin-bottom: 10px;
}

h2 {
  display: block;
  font-size: 1.5em;
  margin-block-start: 0em;
  margin-block-end: 0em;
  margin-inline-start: 0px;
  margin-inline-end: 0px;
  font-weight: bold;
}
</style>
