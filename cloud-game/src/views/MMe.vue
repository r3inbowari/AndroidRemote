<template>
  <div class="me">
    <div v-if="!isLogin" class="not-login">
      <div @click="onLogin" class="user-detail user-detail2">
        <div class="user-avatar">
          <img
            src="https://avatars.githubusercontent.com/u/30739857?s=60&v=4"
            class="avatar-img"
          />
        </div>
        <div>
          <h3>未登录</h3>
          <p>登陆后才可以免费试玩哦</p>
        </div>
      </div>
    </div>

    <div v-if="isLogin" class="not-login">
      <div @click="onCharge" class="user-detail user-detail2">
        <div class="user-avatar">
          <img :src="store.state.info.avatar" class="avatar-img" />
        </div>
        <div>
          <h3>{{ store.state.info.username }}</h3>
          <p>{{ store.state.info.mobile }}</p>
          <div class="user-info">
            <span>LV.{{ store.state.info.level }}</span>
            <span>小可爱</span>
          </div>
          <div class="user-info point">
            <span>剩余时长 {{ store.state.info.point }} 分钟</span>
          </div>
        </div>
      </div>
    </div>

    <!-- <header style="padding: 10px">
      <h2>我的</h2>
    </header> -->

    <ul class="menulist center-block">
      <li @click="help" class="">
        <span>帮助与反馈</span>
        <i class="icon icon-me-arrow"></i>
      </li>
      <li @click="showShare = true">
        <span>分享</span>
        <i class="icon icon-me-arrow"></i>
      </li>
      <li @click="onSetting" class="">
        <span>设置</span>
        <i class="icon icon-me-arrow"></i>
      </li>
      <li v-if="isLogin" @click="onExit" class="">
        <span>退出</span>
        <i class="icon icon-me-arrow"></i>
      </li>
      <li @click="onAbout" class="">
        <span>关于</span>
        <i class="icon icon-me-arrow"></i>
      </li>
    </ul>
  </div>

  <van-share-sheet
    v-model:show="showShare"
    title="立即分享给好友"
    :options="shareOptions"
    @select="onSelect"
  />
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
import { useRouter, useRoute } from 'vue-router'

import { Dialog } from 'vant'

import { VueCookieNext } from 'vue-cookie-next'

export default defineComponent({
  components: {},
  setup() {
    const router = useRouter()
    const route = useRoute()
    let isLogin = ref(false)
    const store = useStore(key)

    if (store.state.token == '') {
      isLogin.value = false
    } else {
      isLogin.value = true
    }

    let showShare = ref(false)
    const shareOptions = [
      [
        { name: '微信', icon: 'wechat' },
        { name: '朋友圈', icon: 'wechat-moments' },
        { name: '微博', icon: 'weibo' },
        { name: 'QQ', icon: 'qq' },
      ],
      [
        { name: '复制链接', icon: 'link' },
        { name: '分享海报', icon: 'poster' },
        { name: '二维码', icon: 'qrcode' },
        { name: '小程序码', icon: 'weapp-qrcode' },
      ],
    ]

    function onLogin() {
      router.push({
        name: 'MLogin',
      })
    }

    function onSetting() {
      router.push({
        name: 'MSetting',
      })
    }

    function help() {
      window.open('https://github.com/r3inbowari/AndroidRemote')
    }

    function onAbout() {
      router.push({
        name: 'MAbout',
      })
    }

    function onSelect(selected) {
      if (selected.icon === 'wechat' || selected.icon === 'wechat-moments') {
        window.location.href = 'weixin://dl/moments'
      } else {
        window.location.href = 'tencent://message/?uin=10987654321'
      }
    }

    function onExit() {
      Dialog.confirm({
        title: '提示',
        message: '确定要退出吗',
      })
        .then(() => {
          VueCookieNext.removeCookie('token')
          store.commit('setToken', '')
          isLogin.value = false
        })
        .catch(() => {
          // on cancel
        })
    }

    function onCharge() {
      router.push({
        name: 'MCharge',
      })
    }
    return {
      onSetting,
      store,
      shareOptions,
      showShare,
      onSelect,
      help,
      isLogin,
      onLogin,
      onExit,
      onAbout,
      onCharge,
    }
  },
})
</script>

<style>
/* 弹出确认 */
.van-dialog__message--has-title {
  padding-top: 8px;
  color: rgb(243, 237, 237);
}
.van-button--default {
  background-color: rgb(38, 49, 58);
}
.van-dialog__cancel {
  color: #fff;
}
</style>

<style>
body {
  /* 全局污染 白色主题 */
  background-color: rgb(30, 40, 49) !important;
  color: #fff;
  background-color: #1e2831;
  min-width: 320px;
  line-height: 1.5;
  background-position: top;
  background-repeat: no-repeat;
  -webkit-user-select: none;
  -ms-user-select: none;
  user-select: none;
}
#app {
  background-color: rgb(30, 40, 49);
}
</style>

<style scoped>
/* 线性背景 */
.me {
  background-image: -webkit-gradient(
    linear,
    left top,
    left bottom,
    from(#40515f),
    to(#1e2831)
  );
  background-image: linear-gradient(180deg, #40515f, #1e2831);
  background-size: 100% 248px;
  background-repeat: no-repeat;
}

h2 {
  display: block;
  font-size: 1.5em;
  margin-block-start: 0em;
  margin-block-end: 0em;
  margin-inline-start: 0px;
  margin-inline-end: 0px;
  font-weight: bold;
  color: rgb(165, 157, 157);
}

h1,
h2,
h3,
h4,
h5,
li,
p,
ul {
  padding: 0;
  margin: 0;
}
h3 {
  color: #fff;
  display: -webkit-box;
  display: -webkit-flex;
  display: flex;
  -webkit-box-align: center;
  -webkit-align-items: center;
  align-items: center;
  -webkit-box-pack: start !important;
  -webkit-justify-content: flex-start !important;
  justify-content: flex-start !important;
}

p {
  display: block;
  margin-block-start: 0.3em;
  margin-block-end: 0.3em;
  margin-inline-start: 0px;
  margin-inline-end: 0px;
  color: #8b969f;
}
.user-detail2 {
  -webkit-box-orient: inherit;
  -webkit-box-direction: inherit;
  -webkit-flex-direction: inherit;
  flex-direction: inherit;
  flex-direction: unset;
}

.user-detail2 {
  -webkit-box-orient: horizontal;
  -webkit-box-direction: reverse;
  -webkit-flex-direction: row-reverse;
  flex-direction: row-reverse;
  -webkit-box-pack: justify;
  -webkit-justify-content: space-between;
  justify-content: space-between;
}

.user-detail {
  position: relative;
  display: -webkit-box;
  display: -webkit-flex;
  display: flex;
  -webkit-box-align: start;
  -webkit-align-items: flex-start;
  align-items: flex-start;
  -webkit-box-orient: vertical;
  -webkit-box-direction: normal;
}

.user-detail2 {
  margin-bottom: 30px;
}
.user-detail2 {
  padding-top: 20px;
  margin: 0 16px 36px;
}

.user-detail2 .user-avatar {
  margin-right: 0;
}

.user-detail2 > div {
  text-align: left;
}
.user-avatar img {
  width: 100%;
  height: 100%;
  border-radius: 50%;
  border: 1px solid hsla(0, 0%, 100%, 0.4);
  box-sizing: border-box;
  object-fit: cover;
}

.user-avatar {
  position: relative;
  display: inline-block;
  width: 64px;
  height: 64px;
}
</style>

<style scoped>
.menulist {
  padding-left: 22px;
  padding-left: -webkit-calc(12px + env(safe-area-inset-left));
  padding-left: calc(12px + env(safe-area-inset-left));
  padding-right: 22px;
  padding-right: -webkit-calc(env(safe-area-inset-right));
  padding-right: calc(env(safe-area-inset-right));
  background: #28333d;
}
.menulist {
  margin-top: 20px;
  margin-bottom: 50px;
  padding-left: -webkit-calc(20px + env(safe-area-inset-left));
  padding-left: calc(20px + env(safe-area-inset-left));
  padding-right: -webkit-calc(20px + env(safe-area-inset-right));
  padding-right: calc(20px + env(safe-area-inset-right));
}
.center-block {
  margin-left: 16px;
  margin-right: 16px;
  border-radius: 8px;
}
ul {
  list-style: none;
}
ul {
  display: block;
  list-style-type: disc;
  margin-block-start: 1em;
  margin-block-end: 1em;
  margin-inline-start: 0px;
  margin-inline-end: 0px;
  padding-inline-start: 40px;
}
.menulist li {
  position: relative;
  line-height: 61px;
  padding-right: 20px;
  border-bottom: 1px solid #2f3c47;
  color: #b4c0c9;
}
.menulist li {
  line-height: 58px;
}
ul {
  list-style: none;
}

.icon-me-arrow {
  width: 16px;
  height: 16px;
  background-image: url(../assets/icon_arrow.png);
}
.menulist li i:last-child {
  float: right;
  margin-top: 22px;
}
li:last-child {
  border-bottom: 0;
}
.menulist li i,
.menulist li span {
  vertical-align: middle;
}
.icon {
  background-position: 50%;
  background-repeat: no-repeat;
  background-size: 100%;
  display: inline-block;
  vertical-align: middle;
}
.menulist li {
  position: relative;
  line-height: 61px;
  padding-right: 20px;
  color: #b4c0c9;
}
</style>

<!-- 分享 -->
<style>
.van-popup {
  background-color: rgb(30, 40, 49);
}

.van-share-sheet__title {
  color: #fff;
}

.van-share-sheet__options--border::before {
  border-top: 1px solid rgb(40, 51, 61);
}

.van-share-sheet__cancel::before {
  display: block;
  height: 2px;
  background-color: rgb(42, 51, 65);
  content: ' ';
}

.van-share-sheet__cancel {
  background-color: rgb(40, 51, 61);
  color: #fff;
}

.van-share-sheet__cancel:hover {
  background-color: rgb(40, 51, 61);
  opacity: 0.5;
}

/* 隐藏水平滚动  */
.van-share-sheet__options {
  scrollbar-width: none;
  -ms-overflow-style: none;
}
.van-share-sheet__options::-webkit-scrollbar {
  display: none; /*ChromeSafari*/
}

/* 用户信息tag */
.user-info span {
  background: hsla(0, 0%, 100%, 0.1);
  border-radius: 12px;
  padding: 3px 8px;
  font-size: 12px;
  color: #b4c0c9;
  margin: 0 8px 4px 0;
  display: inline-block;
}
</style>
