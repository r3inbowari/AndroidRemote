<template>
  <el-carousel :interval="5000" arrow="always">
    <el-carousel-item v-for="item in picsObj.count" :key="item">
      <el-image
        style="width: 1920px; height: 500px"
        :src="base + picsObj.pics[item - 1]"
        fit="cover"
      ></el-image>
    </el-carousel-item>
  </el-carousel>
  <div id="internalHeader" v-show="headerHideActive" class="internalHeader">
    <div class="page-header mini-type">
      <div class="page-header-content">
        <div class="nav-link"><MiniLink></MiniLink></div>
        <div class="nav-search-box">
          <MiniSearach></MiniSearach>
        </div>
        <div class="nav-user-center">
          <MiniUser> </MiniUser>
        </div>
      </div>
    </div>
  </div>

  <div class="v-wrap">
    <div class="game-list-content">
      <HPGame> </HPGame>
    </div>
  </div>

  <!-- back to top -->
  <Backtop></Backtop>
</template>

<script lang="ts">
import { Vue, Options } from 'vue-class-component'
import MiniSearach from '../components/MiniSearch/MiniSearch.vue'
import MiniUser from '../components/User/User.vue'
import MiniLink from '../components/MiniLink.vue'
import Backtop from '../components/Backtop.vue'
import LoginDialog from '../components/Login/Login.vue'
import HPGame from '../components/HPGame/HPGame.vue'
import { HELLOQ } from '../api/nav'

import { useStore } from 'vuex'
import { key } from '../store'

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
    MiniLink,
    HPGame,
    Backtop,
  },
  data() {
    return {
      loginSwitch: true,
    }
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
    HELLOQ('15598870762').then((res) => {
      console.log(res)
    })
  }

  picsObj = {
    count: 2,
    pics: ['mihayo.jpg', 'guangyu.jpg'],
    // we can overwrite the base path
    base: '',
  }

  // webstatic setting
  base = import.meta.env.VITE_STATIC_URL
  // created() {
  //   console.log(import.meta.env.DEV)
  //   console.log(import.meta.env)
  // }

  mounted() {
    window.addEventListener('scroll', this.Scroll, true)

    // this.$refs['nop'].needLogin()
    // const store = useStore(key)
    // store.commit('increment')
    // console.log(store.state.count)
  }

  Scroll() {
    if (window.scrollY > 280) {
      this.headerHideActive = true
    } else {
      this.headerHideActive = false
    }
  }
}
</script>

<style>
.nav-link .nav-link-ul {
  height: 36px;
  display: flex;
  align-items: center;
  list-style: none;
}

.nav-link .nav-link-ul .nav-link-item {
  margin-right: 12px;
}

.nav-link a {
  color: #fff;
  font-size: 14px;
  line-height: 32px;
  display: flex;
  white-space: nowrap;
  text-shadow: 0 1px 1px rgb(0 0 0 / 30%);
  transition: color 0.3s;
  text-decoration: none;
}

.nav-link .logo {
  background-color: rgb(43, 43, 43);
  width: 100px;
  height: 34px;
  margin-top: 1px;

  border-radius: 5px;
}
.el-carousel__item h3 {
  color: #475669;
  font-size: 18px;
  opacity: 0.75;
  line-height: 300px;
  margin: 0;
}

.el-carousel__container {
  height: 500px;
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

  background: rgb(21, 21, 21);
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
  /* height: 2000px; */
  /* min-height: 1200px; */
}
</style>
