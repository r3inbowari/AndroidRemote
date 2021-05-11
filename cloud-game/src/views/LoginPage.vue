<template>
  <div class="login-viewer">
    <div class="page-loading-bar" v-bind:style="{ width: width + '%' }"></div>
    <div class="sign-content">
      <el-card shadow="hover" class="box-card">
        <div class="login-dialog-content">
          <div class="login-logo">
            <img id="logo-bg" :src="logoUrl" alt="logo" />
          </div>
          <div class="login-title">
            <div class="login-top-title"><span>云游</span></div>
            <div class="login-mini-title"><span>云游戏管理中心</span></div>
          </div>
          <transition name="el-fade-in">
            <div v-if="showForm">
              <div class="car-login">
                <div class="login-form">
                  <el-form
                    :model="loginData"
                    ref="loginRef"
                    :rules="loginRules"
                    label-width="80px"
                  >
                    <el-form-item prop="mobile">
                      <el-input
                        prefix-icon="el-icon-mobile-phone"
                        v-model="loginData.mobile"
                        placeholder="手机号"
                        maxlength="11"
                      ></el-input>
                    </el-form-item>

                    <el-form-item prop="password">
                      <el-input
                        prefix-icon="el-icon-lock"
                        placeholder="密码"
                        v-model="loginData.password"
                        maxlength="18"
                        show-password
                      ></el-input>
                    </el-form-item>
                  </el-form>
                </div>
              </div>

              <div class="login-submit">
                <el-button @click="loginClk" type="success">登录</el-button>
              </div>

              <div class="third-login">
                <el-divider>
                  <a class="third-login-a"><span class="dingtalk"></span></a>
                  <a class="third-login-a"><span class="wechat"></span></a>
                </el-divider>
              </div>
            </div>
          </transition>
          <div v-if="showLoading" v-loading="true" style="height: 100px"></div>
          <div class="login-footer">
            <a href="">找回密码</a>
            <span class="login-foot-span"></span>
            <a href="">快速注册</a>
            <span class="login-foot-span"></span>
            <a href="">邮箱登录</a>
            <span class="login-foot-span"></span>
            <a href="">遇到问题</a>
          </div>
        </div>
      </el-card>
    </div>
    <footer class="page-footer">
      <div class="my-links">
        <a
          class="footer-a"
          rel="noopener noreferrer"
          href="http://r3in.top:8881"
          target="_blank"
          >前端视图版本 v{{ v }}</a
        >
        <a
          class="footer-a"
          rel="noopener noreferrer"
          href="https://github.com/r3inbowari/AndroidRemote"
          target="_blank"
          >源代码</a
        >
        <a
          class="footer-a"
          rel="noopener noreferrer"
          href="https://github.com/r3inbowari"
          target="_blank"
          >r3inbowari</a
        >
      </div>
    </footer>
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

import logo_t from '/src/assets/login-bg.png'
import { ElNotification } from 'element-plus'

import { getVersion } from '../api/public'

export default defineComponent({
  data() {
    return {}
  },
  setup() {
    let logoUrl = ref(logo_t)
    let showForm = ref(true)
    let showLoading = ref(false)
    const loginRef = ref(null)

    const loginData = reactive({
      mobile: '',
      password: '',
    })

    let width = ref(0)

    let server_version = ''
    let hash_tag = 'unavailable'
    onMounted(() => {
      getVersion().then((res) => {
        console.log(res.version)
        server_version = res.version
        if (res.hash != 'dev') {
          hash_tag = ' (git-' + res.hash + ')'
        } else {
          hash_tag = ' (开发模式)'
        }
      })
      let b = setInterval(() => {
        width.value += 0.99
        if (width.value > 130) {
          width.value = 0
          clearInterval(b)

          ElNotification({
            title: '全站信息提示',
            dangerouslyUseHTMLString: true,
            message:
              '<div style="font-weight: bold;"><div>服务版本：' +
              server_version +
              hash_tag +
              '</div><div>视图版本：v' +
              import.meta.env['VITE_WEB_VERSION'] +
              '</div></div>',
            duration: 10000,
            position: 'top-left',
            type: 'info',
          })
        }
      }, 16)
    })

    const loginRules = reactive({
      mobile: [
        { required: true, message: '请输入手机号', trigger: 'change' },
        { min: 11, max: 11, message: '手机号长度为11字符', trigger: 'blur' },
      ],
      password: [{ required: true, message: '请输入密码', trigger: 'change' }],
    })

    function loginClk() {
      loginRef.value.validate((valid) => {
        if (valid) {
          // open loading status
          showLoading.value = true
          showForm.value = false
        } else {
          ElNotification({
            title: '发生错误 🤔',
            message: '参数错误(2901)',
            type: 'error',
          })
        }
      })
    }

    let v = import.meta.env['VITE_WEB_VERSION']
    return {
      logoUrl,
      showForm,
      loginData,
      loginRules,
      loginClk,
      showLoading,
      loginRef,

      width,
      v,
    }
  },
})
</script>

<style>
html {
  overflow-x: hidden;
}

body {
  background-color: aliceblue;
  background: aliceblue;
}

/* 全屏背景 */
.login-viewer {
  background-image: url('https://z3.ax1x.com/2021/05/10/gNLSRP.jpg');
  display: -webkit-box;
  display: -ms-flexbox;
  display: flex;
  -webkit-box-orient: vertical;
  -webkit-box-direction: normal;
  -ms-flex-direction: column;
  flex-direction: column;
  -webkit-box-sizing: border-box;
  box-sizing: border-box;
  background-repeat: no-repeat;
  background-color: #b8e5f8;
  background-size: cover;
  width: 100%;
  height: 100vh;
  overflow: auto;
}

.sign-content {
  -webkit-box-flex: 1;
  -ms-flex: 1 1;
  flex: 1 1;
  display: -webkit-box;
  display: -ms-flexbox;
  display: flex;
  -webkit-box-orient: vertical;
  -webkit-box-direction: normal;
  -ms-flex-direction: column;
  flex-direction: column;
  -webkit-box-align: center;
  -ms-flex-align: center;
  align-items: center;
  -webkit-box-pack: center;
  -ms-flex-pack: center;
  justify-content: center;
  border-radius: 2px;
  min-height: 688px;
  height: calc(100% - 42px);
  -webkit-box-sizing: border-box;
  box-sizing: border-box;
}

.cssplogin {
  width: 200px;
  height: 200px;
  background-color: aqua;
}
</style>

<style>
.car-login .el-carousel__item {
  background-color: #fff !important;
}

.third-login span {
  display: inline-block;
  width: 28px;
  height: 29px;
  background-image: url('/src/assets/third-part.png');
  background-repeat: no-repeat;
  background-size: 100%;
  vertical-align: middle;
  margin-right: 5px;
}
.third-login span.wechat {
  background-position: 0 -59.5px;
}
.third-login span.dingtalk {
  background-position: 0 -28.7px;
}
.third-login {
  margin-top: 30px;
  text-align: center;
  width: 320px;
  /* margin-left: 20px; */
  margin-bottom: 30px;
}
.login-logo img {
  width: 320px;
}

.login-logo {
  text-align: center;
}

.login-viewer .login-dialog-content .el-form-item__content {
  margin-left: 0px !important;
}

.login-foot-span {
  display: inline-block;
  width: 1px;
  height: 14px;
  vertical-align: middle;
  margin: 0 8px 3px;
  background-color: #e8e8e8;
}

.login-top-title {
  font-size: 32px;
  line-height: 1.2;
  color: #262626;
}

.login-mini-title {
  color: #595959;
  margin-top: 8px;
  font-size: 18px;
}

.login-title {
  margin-top: 14px;
  text-align: center;
  margin-bottom: 18px;
}

.login-title a,
span::selection {
  background-color: rgb(209, 236, 249);
}
.login-footer a::selection {
  background-color: rgb(209, 236, 249);
}

.login-form {
  width: 320px;
  margin: 0 auto;
  height: 124px;
}

.login-submit .el-button {
  width: 320px;
  /* margin-left: 20px; */
}

.login-submit {
  text-align: center;
}

.login-footer {
  text-align: center;
  width: 320px;
  margin: 0 auto;
  /* margin-left: 20px;
  margin-top: 20px; */
}

.sign-content span {
  font-family: 'Noto Sans SC', sans-serif;
  font-weight: bold;
}

* {
  -webkit-user-select: none;
  -moz-user-select: none;
  -o-user-select: none;
  user-select: none;
}

.box-card .el-card__body {
  padding: 40px;
}

.page-loading-bar {
  width: 500px;
  height: 4px;
  background-color: rgb(27, 120, 207);
  box-shadow: 0 1px 6px 0 rgb(32 33 36 / 28%);
  -webkit-transition: background-color 0.2s;
  transition: background-color 0.2s;
  border-top-right-radius: 2px;
  border-bottom-right-radius: 2px;
}

.el-icon-info:before {
  content: '';
  color: cornflowerblue;
}

.page-footer {
  padding-bottom: 20px;
  font-size: 12px;
  line-height: 21px;
  text-align: center;
  color: #fff;
  text-shadow: 0 1px 2px #999;
}

.footer-a {
  color: #fff !important;
  text-decoration: none;
  color: -webkit-link;
  cursor: pointer;
  -webkit-box-direction: normal;
  -webkit-tap-highlight-color: rgba(18, 18, 18, 0);
}

.footer-a:visited {
  color: none;
}

.footer-a:not(:last-child)::after {
  content: ' · ';
  white-space: pre;
  display: inline-block;
}
</style>
