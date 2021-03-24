<template>
  <div class="nav-user-content">
    <div class="user-login-btn">
      <el-avatar :size="38" src="https://empty" @error="errorHandler">
        <img
          src="https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg"
        />
      </el-avatar>

      <!-- <button @click="onClick">登录</button> -->
    </div>

    <div class="user-avatar"></div>
  </div>
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

  onClick() {
    console.log('clk')
  }
}
</script>

<style></style>
