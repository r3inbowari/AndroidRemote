<template>
  <div class="m-setting">
    <van-nav-bar fixed safe-area-inset-top title="设置">
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

    <div style="margin-top: 60px">
      <ul class="menulist center-block">
        <li @click="help" class="">
          <span>帮助与反馈</span>
          <i class="icon icon-me-arrow"></i>
        </li>
        <li @click="onCleanLocalCache">
          <span>本地缓存清除</span>
          <i class="icon icon-me-arrow"></i>
        </li>
        <li @click="onSetting" class="">
          <span>充值</span>
          <i class="icon icon-me-arrow"></i>
        </li>
      </ul>
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
import { Dialog } from 'vant'

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

    function onCleanLocalCache() {
      Dialog.confirm({
        title: '提示',
        message: '确认清除本地缓存吗',
      })
        .then(() => {
          VueCookieNext.removeCookie('token')
          store.commit('setToken', '')
        })
        .catch(() => {
          // on cancel
        })
    }
    return { backMe, onCleanLocalCache }
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

.m-setting .van-nav-bar {
  background-color: rgb(30, 40, 49);
  box-shadow: 0 0 6px rgb(0 0 0 / 20%);
}

.m-setting .van-nav-bar__title {
  color: #fff;
  font-weight: 600;
  font-size: 18px;
}

.m-setting .van-action-sheet__close {
  top: 12px;
  left: 0px;
}
.m-setting .van-action-sheet__close::before {
  color: #fff;
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
