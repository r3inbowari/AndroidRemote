<template>
  <div id="jiyouCommunity">
    <header class="header main" id="jiyouweb_header">
      <nav class="nav">
        <h1 class="logo" style="color: violet"></h1>
        <ul class="navi">
          <li>
            <a href="/games" id="games"> <span>游戏</span></a>
          </li>
          <li>
            <a href="/games" id="games"> <span>公告/活动</span></a>
          </li>
          <li>
            <a href="/games" id="games"> <span>客服中心</span></a>
          </li>
        </ul>
        <div class="login"></div>
      </nav>
    </header>

    <section class="section-banner">
      <div class="background" :style="bannerBackground">
        <div class="overlay"></div>
      </div>

      <div class="content"></div>

      <div class="slide-progress-bar"></div>
    </section>
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

import bg from '../assets/overlay_main_top.png'

export default defineComponent({
  name: 'App',
  components: {},
  data() {
    return {}
  },
  setup() {
    // let bannerBackground = reactive([
    //   'background-image: url(https://image-glb.qpyou.cn/hubweb/hive_img/web/banner/20200317/41a9e65126c0463ad975ac444027bfa3_1200x490.jpg)',
    // ])
    let bannerBackground = ref(
      'background-image: url(https://image-glb.qpyou.cn/hubweb/hive_img/web/banner/20200317/41a9e65126c0463ad975ac444027bfa3_1200x490.jpg)'
    )

    onMounted(() => {
      function changeBackground(r) {
        r.value =
          'background-image: url("https://image-glb.qpyou.cn/hubweb/hive_img/web/banner/20201203/00d2a88aeb6037a2ada2ffd1ad703e61_1200x490.jpg")'
      }

      setTimeout(() => {
        // bannerBackground.pop()
        // bannerBackground.push(
        //   'background-image: url("https://image-glb.qpyou.cn/hubweb/hive_img/web/banner/20201203/00d2a88aeb6037a2ada2ffd1ad703e61_1200x490.jpg")'
        // )
        // bannerBackground.value =
        changeBackground(bannerBackground)
      }, 5000)
    })

    return {
      bannerBackground,
    }
  },
})
</script>

<style>
body {
  /* 全局污染 白色主题 */
  background-color: white !important;
}
#app {
  background-color: white;
}

.header.main {
  background-color: rgba(0, 0, 0, 0);
}

.header {
  position: relative;
  margin: 0;
  padding: 0;
  background-color: #000;
}

.nav {
  width: 1240px;
  height: 82px;
  position: relative;
  z-index: 100;
  margin: 0 auto;
  padding: 0 20px;
  -webkit-box-sizing: border-box;
  box-sizing: border-box;
  text-align: center;
}

.nav .logo {
  width: 50px;
  height: 50px;
  position: absolute;
  top: 16px;
  left: 20px;
  font-size: 0;
  color: rgba(0, 0, 0, 0);

  /* logo */
  background-color: gray;
  border-radius: 10px;
}

ol,
ul {
  list-style: none;
}

a {
  color: inherit;
}

span {
  font-family: 'Noto Sans SC';
}

.nav .navi li {
  height: 82px;
  display: inline-block;
  position: relative;
  margin: 0 31px;
  padding: 0 2px;
  font-size: 15px;
  font-weight: bold;
  line-height: 1.47;
  color: #fff;
  color: #b0baca;
  vertical-align: middle;
}

.nav .navi li a {
  width: 100%;
  height: 100%;
  display: inline-block;
}

.nav .navi li a:after {
  width: 0;
  height: 100%;
  display: inline-block;
  vertical-align: middle;
  content: '';
}

.nav .navi li:hover a {
  color: #1277ff;
}

.nav .login {
  width: auto;
  height: 100%;
  position: absolute;
  top: 0;
  right: 20px;
  font-size: 14px;
  font-weight: bold;
  line-height: 1.43;
  color: #b0baca;
  text-align: right;
}

/* banner的背景 */
.section-banner .background {
  max-width: 1920px;
  height: 480px;
  position: relative;
  margin: 0 auto;
  background-repeat: no-repeat;
  background-position: 50% 50%;
  background-size: 1920px auto;
  transition: background-image 0.3s ease-in-out;
  opacity: 1;
}

/* background overlay */
.section-banner .background .overlay {
  max-width: 1920px;
  height: 480px;
  position: relative;
  margin: 0 auto;
  background-image: url(../assets/overlay_main_top.png);
  background-repeat: repeat-x;
  background-position: 50% 50%;
  background-size: 1px 480px;
  -webkit-transition: background-image 1s ease-in-out;
  transition: background-image 1s ease-in-out;
}

.section-banner {
  height: 480px;
  margin: -82px auto 92px;
  background-color: #0a0a12;
}
</style>
