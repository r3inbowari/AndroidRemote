<template>
  <div class="nav-user-content">
    <div @click="onClickAvatar" class="user-login-btn">
      <el-avatar :size="50" @error="errorHandler">
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
import LoginDialog from '../Login/Login.vue'

import { key } from '../../store'
import { useStore } from 'vuex'

@Options({
  components: {
    LoginDialog,
  },
})
export default class MiniUser extends Vue {
  store = useStore(key)

  created() {}

  onClickAvatar() {
    if (this.store.state.token === '') {
      this.$refs['nop'].needLogin()
    }
  }

  errorHandler(e) {}
}
</script>

<style></style>
