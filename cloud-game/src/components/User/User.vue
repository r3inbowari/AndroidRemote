<template>
  <div class="nav-user-content">
    <div class="user-login-btn">
      <button @click="onClick">登录</button>
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
