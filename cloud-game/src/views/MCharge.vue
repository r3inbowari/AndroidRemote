<template>
  <div class="m-charge">
    <van-nav-bar fixed safe-area-inset-top title="充值中心">
      <template #left>
        <i
          @click="backMe"
          class="
            van-badge__wrapper
            van-icon van-icon-cross
            van-action-sheet__close
          "
        ></i>
      </template>
    </van-nav-bar>

    <div class="m-charge-content">
      <!-- 头部 -->
      <div class="pay-users">
        <div class="user-detail user-detail2">
          <div class="user-avatar" style="width: 48px; height: 48px">
            <img
              src="https://z3.ax1x.com/2021/05/16/gcvd5d.jpg"
              class="avatar-img"
            />
          </div>
          <div class="baseinfo">
            <h3 class="f16 user-phone">15598870762</h3>
            <p>
              <i class="icon icon-cloud"></i>我的剩余时长：{{
                store.state.info.point
              }}
              分钟
            </p>
          </div>
        </div>
      </div>

      <div class="mglobal-msg">
        <van-notice-bar
          color="#F3AE1F"
          left-icon="volume-o"
          text="【充值公告】现在充值打折打断短腿,欢迎尝试。"
        />
      </div>
      <!-- 购买额度 -->
      <div class="pay-title">选择充值额度</div>

      <div class="pay-list">
        <div v-for="(item, index) in payList.dat" :key="item" class="pay-item">
          <span v-if="index === 0" class="g-tag">体验</span>
          <span v-if="index === 5" class="g-tag">我的648</span>
          <div
            @click="onSelectdPayItem(index)"
            :class="{ paySelected: currPayIndex === index }"
          >
            <p style="font-size: 16px">{{ item.point }}小时</p>
            <div class="pay-money">
              <p>
                <b>¥</b>
                <span style="font-size: 20px">{{ item.money }}</span>
              </p>
              <div style="margin-bottom: 10px; color: grey">
                <del>¥{{ item.original }}</del>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div>
        <div @click="onPay" class="g-Btn g-Btn-green2">
          ¥{{ payList.dat[currPayIndex].money }}微信支付
          <van-loading v-if="payLoading" style="display: inline" color="#fff" />
        </div>
      </div>

      <div
        class="token-div"
        style="
          margin-bottom: 20px;
          text-align: center;
          margin-top: 10px;
          color: gray;
        "
      >
        模拟支付系统 · 2017-2021
      </div>
    </div>
  </div>
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
import { useRouter } from 'vue-router'

import { userLogin, userInfo, userPay } from '../api/user'

import { Toast } from 'vant'

import { VueCookieNext } from 'vue-cookie-next'

export default defineComponent({
  components: {},
  setup() {
    const store = useStore(key)
    const router = useRouter()

    function backMe() {
      router.push({
        name: 'MMe',
      })
    }

    const payList = reactive({
      dat: [
        {
          point: 3,
          money: 2,
          original: 3,
          tag: true,
        },
        {
          point: 12,
          money: 8,
          original: 12,
        },
        {
          point: 72,
          money: 48,
          original: 72,
        },
        {
          point: 128,
          money: 98,
          original: 128,
        },
        {
          point: 328,
          money: 246,
          original: 328,
        },
        {
          point: 648,
          money: 520,
          original: 648,
        },
      ],
    })

    let currPayIndex = ref(0)
    let payLoading = ref(false)
    function onSelectdPayItem(index) {
      currPayIndex.value = index
    }

    function onPay() {
      if (payLoading.value === false) {
        window.location.href = 'weixin://dl/business/?ticket=abbc14432b6fcc'
        payLoading.value = true
        setTimeout(() => {
          userPay(payList.dat[currPayIndex.value].point)
            .then((res) => {
              payLoading.value = false
              userInfo().then((res) => {
                if (res.code === 2005) {
                  store.commit('setInfo', res.data)
                }
              })
              Toast.success('充值成功')
            })
            .catch(() => {
              payLoading.value = false
            })
        }, 10000)
      }
    }

    return {
      backMe,
      payList,
      currPayIndex,
      onSelectdPayItem,
      onPay,
      store,
      payLoading,
    }
  },
})
</script>

<style>
.m-charge .van-nav-bar {
  background-color: rgb(30, 40, 49);
  box-shadow: 0 0 6px rgb(0 0 0 / 20%);
}

.m-charge .van-nav-bar__title {
  color: #fff;
  font-weight: 600;
  font-size: 18px;
}

.m-charge .van-action-sheet__close {
  top: 12px;
  left: 0px;
}
.m-charge .van-action-sheet__close::before {
  color: #fff;
}
</style>

<style scoped>
.m-charge-content {
  margin-top: 60px;
  padding: 5px 16px 0;
}
.pay-users {
  display: -webkit-box;
  display: -webkit-flex;
  display: flex;
  -webkit-box-align: center;
  -webkit-align-items: center;
  align-items: center;
}
.pay-users {
  margin-bottom: 20px;
}
.user-detail2 {
  -webkit-box-orient: inherit;
  -webkit-box-direction: inherit;
  -webkit-flex-direction: inherit;
  flex-direction: inherit;
  flex-direction: unset;
}
.user-detail2 > div {
  text-align: left;
}
.user-detail {
  position: relative;
  display: -webkit-box;
  display: -webkit-flex;
  display: flex;
  -webkit-box-align: start;
  -webkit-align-items: flex-start;
  align-items: flex-start;
  -webkit-box-orient: vertical;
  -webkit-box-direction: normal;
}
.user-avatar {
  position: relative;
  display: inline-block;
  width: 48px;
  height: 48px;
  margin-right: 12px;
}

.user-avatar img {
  width: 100%;
  height: 100%;
  border-radius: 50%;
  border: 1px solid hsla(0, 0%, 100%, 0.4);
  box-sizing: border-box;
  object-fit: cover;
}
.user-detail2 .baseinfo {
  -webkit-box-flex: 1;
  -webkit-flex: 1;
  flex: 1;
}
.user-detail2 > div {
  text-align: left;
}
.f16 {
  font-size: 18px !important;
}
.user-detail > div h3 {
  color: #fff;
  display: -webkit-box;
  display: -webkit-flex;
  display: flex;
  -webkit-box-align: center;
  -webkit-align-items: center;
  align-items: center;
  -webkit-box-pack: center;
  -webkit-justify-content: center;
  justify-content: center;
}
.user-detail2 h3 {
  -webkit-box-pack: start !important;
  -webkit-justify-content: flex-start !important;
  justify-content: flex-start !important;
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
.user-detail > div p {
  color: #8b969f;
  word-break: break-all;
}
p {
  display: block;
  margin-inline-start: 0px;
  margin-inline-end: 0px;
}
</style>

<!-- 额度和充值的部分 -->
<style scoped>
.pay-title {
  margin-top: 20px;
  margin-bottom: 16px;
  font-size: 14px;
  color: #b6b6b6;
}
.pay-title,
.pay-users {
  display: -webkit-box;
  display: -webkit-flex;
  display: flex;
  -webkit-box-align: center;
  -webkit-align-items: center;
  align-items: center;
}
</style>

<style scoped>
.pay-list {
  display: -webkit-box;
  display: -webkit-flex;
  display: flex;
  -webkit-flex-wrap: wrap;
  flex-wrap: wrap;
  margin-bottom: 16px;
  margin-top: -16px;
}

.pay-item {
  /* background-color: aqua; */
  width: 33.33%;
  padding-left: 8px;
  padding-right: 8px;
  padding-top: 16px;
  position: relative;
  box-sizing: border-box;
}

.pay-money {
  color: #03c37e;
  margin-top: auto;
  margin-bottom: auto;
  text-align: center;
}

.pay-item > div {
  cursor: pointer;
  position: relative;
  background: linear-gradient(
    135deg,
    rgba(136, 139, 153, 0.07),
    rgba(77, 80, 97, 0.08)
  );
  border: 1px solid rgba(106, 115, 107, 0.24);
  border-radius: 2px;
  display: -webkit-box;
  display: -webkit-flex;
  display: flex;
  -webkit-box-orient: vertical;
  -webkit-box-direction: normal;
  -webkit-flex-direction: column;
  flex-direction: column;
  -webkit-box-align: center;
  -webkit-align-items: center;
  align-items: center;
  color: #fff;
}

.paySelected {
  border: 1px solid rgb(3, 196, 125, 0.9) !important;
}

p {
  display: block;
  margin-block-start: 0.5em;
  margin-block-end: 0em;
  margin-inline-start: 0px;
  margin-inline-end: 0px;
}

.g-tag {
  right: -3px;
  top: 2px;
  font-weight: 700;
  position: absolute;
  z-index: 2000;
}
.g-tag {
  padding: 2px 4px;
  color: #664a1a;
  background-image: -webkit-gradient(
    linear,
    left top,
    left bottom,
    from(#f3e1ba),
    to(#e4bf84)
  );
  background-image: linear-gradient(180deg, #f3e1ba, #e4bf84);
  border-radius: 0 8px 0 8px;
}

.g-Btn-green2:hover {
  -webkit-transition: inherit;
  transition: inherit;
  -webkit-transition: unset;
  transition: unset;
  background: #03c47e;
}

.g-Btn {
  margin-top: 16px;
  font-size: 16px;
  border-radius: 48px;
  height: 48px;
  line-height: 46px;
  width: 100%;
  letter-spacing: 2px;
  padding: 0;
}

.g-Btn-green2,
a.g-Btn-green2 {
  border: 1px solid #03c47e;
  background: #03c47e;
  color: #fff;
  font-weight: 700;
  text-align: center;
}
</style>
