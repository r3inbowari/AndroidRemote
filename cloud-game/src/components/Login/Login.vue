<template>
  <el-dialog
    :close-on-click-modal="false"
    v-model="dialogVisible"
    width="400px"
    :before-close="handleClose"
  >
    <div class="login-dialog-content">
      <div class="login-logo">
        <img id="logo-bg" :src="logoUrl" alt="logo" />
      </div>
      <div class="login-title">
        <div class="login-top-title"><span>云游</span></div>
        <div class="login-mini-title"><span>专业云游戏中心</span></div>
      </div>

      <div class="login-form">
        <el-form ref="loginData" :model="loginData" label-width="80px">
          <el-form-item>
            <el-input
              prefix-icon="el-icon-mobile-phone"
              v-model="loginData.phone"
              placeholder="手机号"
            ></el-input>
          </el-form-item>

          <el-form-item>
            <el-input
              prefix-icon="el-icon-lock"
              placeholder="密码"
              v-model="loginData.passwd"
              show-password
            ></el-input>
          </el-form-item>
        </el-form>
      </div>

      <div class="login-submit">
        <el-button type="success">登录</el-button>
      </div>

      <div class="third-login">
        <el-divider>
          <a class="third-login-a"><span class="dingtalk"></span></a>
          <a class="third-login-a"><span class="wechat"></span></a>
        </el-divider>
      </div>

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
  </el-dialog>
</template>

<script lang="ts">
import { Vue, Options } from 'vue-class-component'
import img1 from '../../assets/login-bg.png'
import imgUrl0 from '/src/assets/login-bg.png'

@Options({
  components: {},
  data() {
    return {
      logoUrl: imgUrl0,
      loginData: {
        phone: '',
        passwd: '',
      },
    }
  },
})
export default class Login extends Vue {
  dialogVisible = true
  created() {}

  onClick() {
    console.log('clk')
  }

  // 登录窗口被关闭时
  handleClose() {}

  mounted() {}
}
</script>

<style>
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
  margin-left: 20px;
  margin-bottom: 30px;
}
.login-logo img {
  width: 320px;
}

.login-logo {
  text-align: center;
}

.login-dialog-content .el-form-item__content {
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
  margin-left: 20px;
}

.login-submit .el-button {
  width: 320px;
  margin-left: 20px;
}

.login-footer {
  text-align: center;
  width: 320px;
  margin-left: 20px;
  margin-top: 20px;
}
</style>
