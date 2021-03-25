<template>
  <div class="nav-user-content">
    <div @click="onClickAvatar" class="user-login-btn">
      <el-avatar :size="38" @error="errorHandler">
        <img
          src="https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg"
        />
      </el-avatar>

      <!-- <button @click="onClick">登录</button> -->
    </div>

    <div class="user-avatar"></div>
  </div>
  <LoginDialog ref="nop"></LoginDialog>
</template>

<script lang="ts">
import { Vue, Options } from 'vue-class-component'
import { HELLO } from '../../api/nav'
import LoginDialog from '../Login/Login.vue'

@Options({
  components: {
    LoginDialog,
  },
})
export default class MiniUser extends Vue {
  created() {}

  onClickAvatar() {
    this.$refs['nop'].needLogin()
  }

  errorHandler(e) {}
}
</script>

<style></style>
