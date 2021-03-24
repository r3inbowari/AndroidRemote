<template>
  <el-carousel :interval="5000" arrow="always">
    <el-carousel-item v-for="item in 4" :key="item">
      <h3>{{ item }}</h3>
    </el-carousel-item>
  </el-carousel>
  <div id="internalHeader" v-if="headerHideActive" class="internalHeader">
    <div class="page-header mini-type">
      <div class="page-header-content">
        <div class="nav-link">游戏</div>
        <div class="nav-search-box">
          <MiniSearach></MiniSearach>
        </div>
        <div class="nav-user-center">
          <MiniUser> </MiniUser>
        </div>
      </div>
    </div>
  </div>
  <LoginDialog></LoginDialog>
  <div class="v-wrap"></div>
</template>

<script lang="ts">
import { Vue, Options } from 'vue-class-component'
import MiniSearach from '../components/MiniSearch/MiniSearch.vue'
import MiniUser from '../components/User/User.vue'
import LoginDialog from '../components/Login/Login.vue'
import { HELLO } from '../api/nav'

// Component definition
@Options({
  // Define component options
  watch: {
    count: (value) => {
      console.log(value)
    },
  },
  components: {
    MiniSearach,
    MiniUser,
    LoginDialog,
  },
})
export default class HomePage extends Vue {
  // The behavior in class is the same as the current
  count = 0
  headerHideActive = false
  increment() {
    this.count++
  }

  created() {
    HELLO(1).then((res) => {
      console.log(res)
    })
  }

  // created() {
  //   console.log(import.meta.env.DEV)
  //   console.log(import.meta.env)
  // }

  mounted() {
    window.addEventListener('scroll', this.Scroll, true)
  }

  Scroll() {
    // console.log(window.scrollY)
    if (window.scrollY > 280) {
      console.log('sd')
      this.headerHideActive = true
    } else {
      this.headerHideActive = false
    }
  }
}
</script>

<style>
.el-carousel__item h3 {
  color: #475669;
  font-size: 18px;
  opacity: 0.75;
  line-height: 300px;
  margin: 0;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n + 1) {
  background-color: #d3dce6;
}

.el-carousel__button {
  width: 10px;
}

.internalHeader {
  min-width: 998;
  min-height: 56;

  z-index: 999;
  display: flex;
  /* display: none; */
  margin: 0;
  padding: 0;

  position: fixed;

  top: 0px;

  width: 100%;
}

.internalHeader .mini-type {
  position: relative;
  box-shadow: 0 2px 4px 0 rgb(0 0 0 / 8%);
  box-sizing: border-box;
}

.page-header {
  width: 100%;
  height: 56px;
  z-index: 1;
}

.page-header-content {
  box-sizing: border-box;
  padding: 10px 24px;
  line-height: 30px;
  position: relative;
  margin: 0 auto;
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
}

.nav-search-box {
  width: 500px;
}

.v-wrap {
  height: 2000px;
}
</style>
