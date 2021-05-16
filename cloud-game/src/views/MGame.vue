<template>
  <header class="header">
    <div class="header-box">
      <div class="nav">
        <div class="nav-mgame-logo"></div>
        <div class="nav-mgame-search">
          <van-search
            v-model="searchValue"
            shape="round"
            background="#1E2831"
            placeholder="你想玩什么呢"
          />
        </div>
        <div class="nav-mgame-msg">
          <div
            style="
              background-color: rgb(40, 51, 61);
              width: 20px;
              height: 34px;
              border-radius: 6px;
            "
          ></div>
        </div>
      </div>
    </div>
  </header>
  <section class="mgame-content">
    <div class="banner">
      <van-swipe
        v-if="bannerData.dat.lenght !== 0"
        :autoplay="3000"
        @change="onChangeSwipe"
      >
        <van-swipe-item
          class="mbanner-swipe"
          v-for="(item, index) in bannerData.dat"
          :key="index"
        >
          <img class="swipe-img" :src="item.cover" />
        </van-swipe-item>
      </van-swipe>
    </div>

    <div class="mglobal-msg">
      <van-notice-bar
        color="#F3AE1F"
        left-icon="volume-o"
        text="【手游公告】《光遇》最新版本「集结季」已上线，让我们一起结伴启程！《原神》全新1.5版本「玉扉绕尘歌」现已推出！各位旅行者集合！【手游上新】《宝可梦大探险》、《航海王热血航线》、《少女的王座》、《梦想协奏曲！少女乐团派对》、《奥奇传说》、《黑暗料理王》、《大征服者：罗马》【端游上新】《假面骑士：剑》、《正当防卫2》（试玩）、《生化危机3重制版》（试玩）、《极品飞车9》"
      />
    </div>

    <div class="nav-main">
      <van-tabs sticky offset-top="48" animated v-model:active="activeTab">
        <van-tab title="推荐">
          <section class="recommend">
            <header class="recommend-name">
              <h2>排行榜</h2>
            </header>

            <div class="recommend-tab">
              <ul class="recommend-list">
                <li
                  v-for="(item, index) in recommendData.dat"
                  class="recommend-item"
                >
                  <div class="recommend-icon">
                    <img :src="item.icon" lazy="loaded" />
                  </div>
                  <div class="recommend-title">
                    <h3>{{ item.title }}</h3>
                    <p>{{ item.desc }}</p>
                  </div>
                  <div class="recommend-btn btn"><p>试玩</p></div>
                </li>
              </ul>
            </div>
          </section>
        </van-tab>
        <van-tab title="手机游戏"></van-tab>
        <van-tab title="更多">更多</van-tab>
      </van-tabs>
      <div>11</div>
      <div>11</div>
      <div>11</div>
      <div>11</div>
      <div>11</div>
      <div>11</div>
      <div>11</div>
      <div>12221</div>
      <div>11</div>
      <div>11</div>
      <div>11</div>
      <div>11</div>
      <div>12221</div>
      <div>11</div>
      <div>11</div>
      <div>11</div>
      <div>11</div>
      <div>12221</div>
    </div>
  </section>
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

export default defineComponent({
  components: {},
  setup() {
    let searchValue = ref('')
    const bannerData = reactive({
      dat: [],
    })

    const recommendData = reactive({
      dat: [],
    })
    getBanner().then((res) => {
      bannerData.dat = res.Data
    })
    getUpdate().then((res) => {
      recommendData.dat = res.Data
    })

    function onChangeSwipe(index) {
      // console.log(index)
    }

    let activeTab = ref(0)
    return { bannerData, recommendData, onChangeSwipe, searchValue, activeTab }
  },
})
</script>

<style scoped>
.headerbox {
  background: #1e2831;
  height: 58px;
  padding-top: constant(safe-area-inset-top);
  padding-top: env(safe-area-inset-top);
  position: fixed;
  width: 100%;
  display: -webkit-box;
  display: -webkit-flex;
  display: flex;
  -webkit-box-align: center;
  -webkit-align-items: center;
  align-items: center;
  z-index: 996;
  top: 0;
  -webkit-transition: background 0.3s cubic-bezier(0.25, 0.8, 0.25, 1);
  transition: background 0.3s cubic-bezier(0.25, 0.8, 0.25, 1);
}
.nav {
  position: relative;
  padding: 0 16px;
  justify-content: center;
  max-width: 1920px;
  padding: 0 8%;
  margin: 0 auto;
  display: -webkit-box;
  display: -webkit-flex;
  display: flex;
  -webkit-box-align: center;
  -webkit-align-items: center;
  align-items: center;
  -webkit-box-flex: 1;
  -webkit-flex: 1;
  flex: 1;
}

.header {
  height: 58px;
  height: -webkit-calc(58px + constant(safe-area-inset-top));
  height: calc(58px + constant(safe-area-inset-top));
  height: -webkit-calc(58px + env(safe-area-inset-top));
  height: calc(58px + env(safe-area-inset-top));
  font-size: 14px;

  position: fixed;
  top: 0px;
  left: 0px;
  z-index: 10;
  background-color: rgb(30, 40, 49);
  width: 100%;
}

.nav-mgame-search {
  -webkit-flex-shrink: 0;
  flex-shrink: 0;
  max-width: 200px;
  margin-left: 16px;
}

.nav-mgame-logo {
  width: 34px;
  height: 34px;
  margin-left: 0;
  margin-right: 0;
  background-color: rgb(40, 51, 61);
  border-radius: 6px;
}

.nav-mgame-msg {
  margin-left: auto;
  margin-right: -8px;
  display: -webkit-box;
  display: -webkit-flex;
  display: flex;
}
</style>

<!-- vant定制 -->
<style>
body {
  color: #fff;
  background-color: #1e2831;
  min-width: 320px;
  line-height: 1.5;
  background-position: top;
  background-repeat: no-repeat;
  -webkit-user-select: none;
  -ms-user-select: none;
  user-select: none;
}

/* 搜索框 */
.van-search__content {
  background-color: rgb(44, 57, 70);
}
.van-field__control::placeholder {
  color: rgba(218, 218, 218, 0.849);
}
.van-icon::before {
  color: rgba(218, 218, 218, 0.849);
}

/* 全局信息框定制 */
.mglobal-msg .van-icon::before {
  color: rgb(243, 174, 31);
}
.mglobal-msg .van-notice-bar {
  background-color: rgb(50, 46, 40);
  border-radius: 6px;
  margin-top: 8px;
}
/* 导航tab定制 */
.nav-main .van-tabs__nav {
  background-color: rgb(30, 40, 49);
}
.nav-main .van-tab {
  font-size: 16px;
}
.nav-main .van-tab span {
  color: #fff;
  opacity: 0.5;
}
.nav-main .van-tab--active span {
  opacity: 1 !important;
  font-weight: 700;
  font-size: 18px;
}
.nav-main .van-tabs__line {
  background-color: rgb(3, 196, 126);
}
</style>

<!-- 内容部分 -->
<style scoped>
.mgame-content {
  padding: 16px;
  padding-top: 8px;
  margin-top: 54px;
}

.swipe-img {
  height: 220px;
  object-fit: cover;
  height: auto;
  max-width: 100%;
  max-height: 100%;
  vertical-align: bottom;
  bottom: 0;
  -o-object-fit: fill;
  object-fit: fill;
  border-radius: 6px;
}

.mbanner-swipe {
  border-radius: 6px;
}

.nav-main {
  margin-top: 8px;
}
</style>

<!-- 推荐卡 -->
<style scoped>
.recommend-content {
  padding-top: 16px;
}

.recommend {
  padding: 16px;
  padding-bottom: 24px;
  background: #28333d;
  border-radius: 6px;
  margin-bottom: 40px;
  margin-top: 10px;
}

.recommend-name {
  display: -webkit-box;
  display: -webkit-flex;
  display: flex;
  -webkit-box-align: center;
  -webkit-align-items: center;
  align-items: center;
  -webkit-box-pack: justify;
  -webkit-justify-content: space-between;
  justify-content: space-between;
}

h2 {
  display: block;
  font-size: 1.5em;
  margin-block-start: 0.83em;
  margin-block-end: 0.83em;
  margin-inline-start: 0px;
  margin-inline-end: 0px;
  font-weight: bold;
  font-size: 18px;
  color: #fff;
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

ul {
  list-style: none;
}

.recommend-tab {
  padding-top: 16px;
}
.recommend-list > li {
  margin-bottom: 32px;
}

.recommend-list > li:last-child {
  margin-bottom: 0;
}

.recommend-item {
  display: -webkit-box;
  display: -webkit-flex;
  display: flex;
  -webkit-box-align: center;
  -webkit-align-items: center;
  align-items: center;
  -webkit-box-pack: justify;
  -webkit-justify-content: space-between;
  justify-content: space-between;
}

.recommend-icon {
  width: 56px;
  height: 56px;
  -webkit-flex-shrink: 0;
  flex-shrink: 0;
  display: -webkit-box;
  display: -webkit-flex;
  display: flex;
  -webkit-box-align: center;
  -webkit-align-items: center;
  align-items: center;
  box-sizing: border-box;
  border-radius: 12px;
  position: relative;
}

.recommend-icon img {
  width: 100%;
  height: 100%;
  border-radius: 12px;
}
.recommend-title {
  margin: 0 16px;
  -webkit-box-flex: 1;
  -webkit-flex: 1;
  flex: 1;
  overflow: hidden;
}
.recommend-title > h3 {
  font-weight: 700;
  color: #fff;
}
.recommend-title > p {
  color: #b4c0c9;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.recommend-btn {
  padding: 0;
  width: 64px;
  -webkit-flex-shrink: 0;
  flex-shrink: 0;
  height: 28px;
  line-height: 26px;
  border: 1px solid #03c47e;
  background: #03c47e;
  color: #fff;
  font-weight: 700;
}

.btn {
  vertical-align: middle;
  text-align: center;
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
}
</style>
