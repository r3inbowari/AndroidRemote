<template>
  <div class="m-login">
    <van-nav-bar fixed safe-area-inset-top title="登录">
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
    <div class="taken-div" style="height: 46px; width: 0px"></div>
    <div class="login2">
      <div class="login-box">
        <div class="login-header">
          <h1 class="f16">登录后即可免费玩</h1>
        </div>
        <p>手机号</p>
        <div class="f14 input">
          <input
            placeholder="请输入手机号码"
            type="tel"
            maxlength="14"
            minlength="0"
            v-model="user.mobile"
          />
        </div>
        <p>账号密码</p>
        <div class="f14 input">
          <input
            placeholder="请输入密码"
            type="password"
            maxlength="14"
            minlength="0"
            v-model="user.password"
          />
        </div>
        <div class="g-Btn g-Btn-green2">
          <div @click="onLoginXHR" v-if="!loading">登录</div>
          <van-loading
            v-if="loading"
            style="margin-top: 5px"
            size="24px"
            vertical
          ></van-loading>
        </div>
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

    let loading = ref(false)
    function backMe() {
      router.push({
        name: 'MMe',
      })
    }

    const user = reactive({
      mobile: '',
      password: '',
    })
    function onLoginXHR() {
      loading.value = true
      // userLogin()
      if (user.mobile != '' && user.password != '') {
        userLogin(user.mobile, user.password)
          .then((res) => {
            if (res.code === 2005) {
              VueCookieNext.setCookie('token', res.data, { expire: '7d' })
              store.commit('setToken', res.data)
              userInfo().then((res) => {
                if (res.code === 2005) {
                  store.commit('setInfo', res.data)
                  router.push({
                    name: 'MMe',
                  })
                  loading.value = false
                }
              })
            } else {
              Toast({
                forbidClick: true,
                message: '账号密码错误',
                className: 'login-msg',
                icon: 'cross',
              })
              loading.value = false
            }
          })
          .catch(() => {
            Toast({
              forbidClick: true,
              message: '服务错误',
              className: 'login-msg',
              icon: 'cross',
            })
            loading.value = false
          })
      } else {
        Toast({
          forbidClick: true,
          message: '参数错误',
          className: 'login-msg',
          icon: 'cross',
        })
        loading.value = false
      }
    }
    return { backMe, loading, onLoginXHR, user }
  },
})
</script>

<style>
.m-login .van-nav-bar {
  background-color: rgb(30, 40, 49);
  box-shadow: 0 0 6px rgb(0 0 0 / 20%);
}

.m-login .van-nav-bar__title {
  color: #fff;
  font-weight: 600;
  font-size: 18px;
}

.m-login .van-action-sheet__close {
  top: 12px;
  left: 0px;
}
.m-login .van-action-sheet__close::before {
  color: #fff;
}

.login-msg {
  background-color: rgba(125, 131, 136, 0.4);
}
</style>

<style scoped>
.login2 {
  padding: 24px 16px;
  position: relative;
}

.login-box {
  padding-top: 84px;
  padding-left: 24px;
  padding-left: -webkit-calc(16px + env(safe-area-inset-left));
  padding-left: calc(16px + env(safe-area-inset-left));
  padding-right: 24px;
  padding-right: -webkit-calc(16px + env(safe-area-inset-right));
  padding-right: calc(16px + env(safe-area-inset-right));
  top: 0;
  bottom: 0;
  right: 0;
  left: 0;
  position: absolute;
  color: #a5a9ac;
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

.login-avatar,
.login-header {
  position: absolute;
  left: 24px;
  left: -webkit-calc(16px + env(safe-area-inset-left));
  left: calc(16px + env(safe-area-inset-left));
  right: 24px;
  right: -webkit-calc(16px + env(safe-area-inset-right));
  right: calc(16px + env(safe-area-inset-right));
  top: 24px;
  height: 60px;
  display: -webkit-box;
  display: -webkit-flex;
  display: flex;
  -webkit-box-pack: center;
  -webkit-justify-content: center;
  justify-content: center;
  -webkit-box-align: center;
  -webkit-align-items: center;
  align-items: center;
}

.login-box > .input {
  margin-top: 4px;
  margin-bottom: 10px;
  position: relative;
}

.f14 {
  font-size: 16px !important;
}
.login-box > .g-Btn {
  width: 100%;
  margin-bottom: 10px;
}
.g-Btn {
  font-size: 16px;
}

.g-Btn-green2,
a.g-Btn-green2 {
  border: 1px solid #03c47e;
  background: #03c47e;
  color: #fff;
  font-weight: 700;
}

.g-Btn {
  vertical-align: middle;
  text-align: center;
  padding: 0 32px;
  height: 36px;
  line-height: 34px;
  display: inline-block;
  border-radius: 20px;
  box-sizing: border-box;
  white-space: nowrap;
  -webkit-user-select: none;
  -ms-user-select: none;
  user-select: none;
  cursor: pointer;
  margin-top: 10px;
}
.input label {
  position: absolute;
  left: 0;
  right: 0;
  bottom: 0;
  top: 0;
  border-radius: 36px;
  background: #1e2831;
  z-index: -1;
  border: 1px solid transparent;
  pointer-events: none;
}
.input,
.input label {
  box-sizing: border-box;
}

button,
input {
  overflow: visible;
}

button,
input,
optgroup,
select,
textarea {
  font-family: inherit;
  font-size: 100%;
  line-height: 1.15;
  margin: 0;
}

input {
  -webkit-writing-mode: horizontal-tb !important;
  text-rendering: auto;
  color: -internal-light-dark(black, white);
  letter-spacing: normal;
  word-spacing: normal;
  text-transform: none;
  text-indent: 10px;
  text-shadow: none;
  display: inline-block;
  text-align: start;
  appearance: auto;
  background-color: -internal-light-dark(rgb(255, 255, 255), rgb(59, 59, 59));
  -webkit-rtl-ordering: logical;
  cursor: text;
  margin: 0em;
  font: 400 13.3333px Arial;
  padding: 1px 2px;
  border-width: 2px;
  border-style: inset;
  border-color: -internal-light-dark(rgb(118, 118, 118), rgb(133, 133, 133));
  border-image: initial;
  width: 100%;
}
.input input {
  height: 40px;
  line-height: 28px;
  -webkit-box-flex: 1;
  -webkit-flex: 1;
  flex: 1;
  border-radius: 8px;
}

input {
  padding: 0;
  color: #fff;
  border: none;
  outline: none;
  background-color: rgb(40, 51, 61);
}

.input:focus {
  outline: none;
  border: 1px solid red;
  border-radius: 8px;
}
</style>
