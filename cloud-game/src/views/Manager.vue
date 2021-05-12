<template>
  <div class="devapp" :class="{ dark: theme, light: !theme }">
    <div class="app-sider">
      <!-- 占位 -->
      <div style="height: 15px"></div>
      <!-- 列表 -->
      <div class="sider-content">
        <a
          class="sider-item"
          :class="{ selectside: currIndex === 0 }"
          @click="onJump(0)"
        >
          <svg
            stroke="currentColor"
            fill="currentColor"
            stroke-width="0"
            version="1"
            viewBox="0 0 48 48"
            enable-background="new 0 0 48 48"
            height="1em"
            width="1em"
            xmlns="http://www.w3.org/2000/svg"
          >
            <polygon
              fill="#3F51B5"
              points="42,37 6,37 6,25 16,10 30,17 42,6"
            ></polygon>
            <polygon
              fill="#00BCD4"
              points="42,42 6,42 6,32 16,24 30,26 42,17"
            ></polygon>
          </svg>
          <div class="item-desc">概况</div>
        </a>
        <a
          class="sider-item"
          :class="{ selectside: currIndex === 1 }"
          @click="onJump(1)"
        >
          <svg
            stroke="currentColor"
            fill="currentColor"
            stroke-width="0"
            version="1"
            viewBox="0 0 48 48"
            enable-background="new 0 0 48 48"
            height="1em"
            width="1em"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              fill="#7CB342"
              d="M24,4C13,4,4,13,4,24s9,20,20,20s20-9,20-20S35,4,24,4z"
            ></path>
            <path
              fill="#0277BD"
              d="M45,24c0,11.7-9.5,21-21,21S3,35.7,3,24S12.3,3,24,3S45,12.3,45,24z M23.8,33.7c0-0.4-0.2-0.6-0.6-0.8 c-1.3-0.4-2.5-0.4-3.6-1.5c-0.2-0.4-0.2-0.8-0.4-1.3c-0.4-0.4-1.5-0.6-2.1-0.8c-0.8,0-1.7,0-2.7,0c-0.4,0-1.1,0-1.5,0 c-0.6-0.2-1.1-1.1-1.5-1.7c0-0.2,0-0.6-0.4-0.6c-0.4-0.2-0.8,0.2-1.3,0c-0.2-0.2-0.2-0.4-0.2-0.6c0-0.6,0.4-1.3,0.8-1.7 c0.6-0.4,1.3,0.2,1.9,0.2c0.2,0,0.2,0,0.4,0.2c0.6,0.2,0.8,1,0.8,1.7c0,0.2,0,0.4,0,0.4c0,0.2,0.2,0.2,0.4,0.2 c0.2-1.1,0.2-2.1,0.4-3.2c0-1.3,1.3-2.5,2.3-2.9c0.4-0.2,0.6,0.2,1.1,0c1.3-0.4,4.4-1.7,3.8-3.4c-0.4-1.5-1.7-2.9-3.4-2.7 c-0.4,0.2-0.6,0.4-1,0.6c-0.6,0.4-1.9,1.7-2.5,1.7c-1.1-0.2-1.1-1.7-0.8-2.3c0.2-0.8,2.1-3.6,3.4-3.1c0.2,0.2,0.6,0.6,0.8,0.8 c0.4,0.2,1.1,0.2,1.7,0.2c0.2,0,0.4,0,0.6-0.2c0.2-0.2,0.2-0.2,0.2-0.4c0-0.6-0.6-1.3-1-1.7c-0.4-0.4-1.1-0.8-1.7-1.1 c-2.1-0.6-5.5,0.2-7.1,1.7s-2.9,4-3.8,6.1c-0.4,1.3-0.8,2.9-1,4.4c-0.2,1-0.4,1.9,0.2,2.9c0.6,1.3,1.9,2.5,3.2,3.4 c0.8,0.6,2.5,0.6,3.4,1.7c0.6,0.8,0.4,1.9,0.4,2.9c0,1.3,0.8,2.3,1.3,3.4c0.2,0.6,0.4,1.5,0.6,2.1c0,0.2,0.2,1.5,0.2,1.7 c1.3,0.6,2.3,1.3,3.8,1.7c0.2,0,1-1.3,1-1.5c0.6-0.6,1.1-1.5,1.7-1.9c0.4-0.2,0.8-0.4,1.3-0.8c0.4-0.4,0.6-1.3,0.8-1.9 C23.8,35.1,24,34.3,23.8,33.7z M24.2,14.3c0.2,0,0.4-0.2,0.8-0.4c0.6-0.4,1.3-1.1,1.9-1.5c0.6-0.4,1.3-1.1,1.7-1.5 c0.6-0.4,1.1-1.3,1.3-1.9c0.2-0.4,0.8-1.3,0.6-1.9c-0.2-0.4-1.3-0.6-1.7-0.8c-1.7-0.4-3.1-0.6-4.8-0.6c-0.6,0-1.5,0.2-1.7,0.8 c-0.2,1.1,0.6,0.8,1.5,1.1c0,0,0.2,1.7,0.2,1.9c0.2,1-0.4,1.7-0.4,2.7c0,0.6,0,1.7,0.4,2.1L24.2,14.3z M41.8,29 c0.2-0.4,0.2-1.1,0.4-1.5c0.2-1,0.2-2.1,0.2-3.1c0-2.1-0.2-4.2-0.8-6.1c-0.4-0.6-0.6-1.3-0.8-1.9c-0.4-1.1-1-2.1-1.9-2.9 c-0.8-1.1-1.9-4-3.8-3.1c-0.6,0.2-1,1-1.5,1.5c-0.4,0.6-0.8,1.3-1.3,1.9c-0.2,0.2-0.4,0.6-0.2,0.8c0,0.2,0.2,0.2,0.4,0.2 c0.4,0.2,0.6,0.2,1,0.4c0.2,0,0.4,0.2,0.2,0.4c0,0,0,0.2-0.2,0.2c-1,1.1-2.1,1.9-3.1,2.9c-0.2,0.2-0.4,0.6-0.4,0.8 c0,0.2,0.2,0.2,0.2,0.4c0,0.2-0.2,0.2-0.4,0.4c-0.4,0.2-0.8,0.4-1.1,0.6c-0.2,0.4,0,1.1-0.2,1.5c-0.2,1.1-0.8,1.9-1.3,2.9 c-0.4,0.6-0.6,1.3-1,1.9c0,0.8-0.2,1.5,0.2,2.1c1,1.5,2.9,0.6,4.4,1.3c0.4,0.2,0.8,0.2,1.1,0.6c0.6,0.6,0.6,1.7,0.8,2.3 c0.2,0.8,0.4,1.7,0.8,2.5c0.2,1,0.6,2.1,0.8,2.9c1.9-1.5,3.6-3.1,4.8-5.2C40.6,32.4,41.2,30.7,41.8,29z"
            ></path>
          </svg>
          <div class="item-desc">全局</div>
        </a>
        <a
          class="sider-item"
          :class="{ selectside: currIndex === 2 }"
          @click="onJump(2)"
        >
          <svg
            stroke="currentColor"
            fill="currentColor"
            stroke-width="0"
            version="1"
            viewBox="0 0 48 48"
            enable-background="new 0 0 48 48"
            height="1em"
            width="1em"
            xmlns="http://www.w3.org/2000/svg"
          >
            <rect
              x="16.7"
              y="1.6"
              transform="matrix(-.707 -.707 .707 -.707 24 57.941)"
              fill="#FFA000"
              width="14.6"
              height="44.8"
            ></rect>
            <g fill="#9E6400">
              <rect
                x="17.9"
                y="20.2"
                transform="matrix(-.707 -.707 .707 -.707 21.177 51.125)"
                width="6.6"
                height="2"
              ></rect>
              <rect
                x="22.3"
                y="15.2"
                transform="matrix(-.707 -.707 .707 -.707 29.833 44.71)"
                width="3.7"
                height="2"
              ></rect>
              <rect
                x="25.9"
                y="12.2"
                transform="matrix(-.707 -.707 .707 -.707 40.49 43.125)"
                width="6.6"
                height="2"
              ></rect>
              <rect
                x="31.2"
                y="6.3"
                transform="matrix(.707 -.707 .707 .707 3.643 25.147)"
                width="2"
                height="3.7"
              ></rect>
              <rect
                x="6.3"
                y="31.2"
                transform="matrix(-.707 -.707 .707 -.707 -8.794 60.71)"
                width="3.7"
                height="2"
              ></rect>
              <rect
                x="9.9"
                y="28.2"
                transform="matrix(-.707 -.707 .707 -.707 1.863 59.125)"
                width="6.6"
                height="2"
              ></rect>
              <rect
                x="14.3"
                y="23.2"
                transform="matrix(-.707 -.707 .707 -.707 10.52 52.71)"
                width="3.7"
                height="2"
              ></rect>
            </g>
          </svg>
          <div class="item-desc">规划</div>
        </a>
        <a
          class="sider-item"
          :class="{ selectside: currIndex === 3 }"
          @click="onJump(3)"
        >
          <svg
            stroke="currentColor"
            fill="currentColor"
            stroke-width="0"
            version="1"
            viewBox="0 0 48 48"
            enable-background="new 0 0 48 48"
            height="1em"
            width="1em"
            xmlns="http://www.w3.org/2000/svg"
          >
            <g fill="#1976D2">
              <path
                d="M38,13h-3c-5.5,0-10,4.5-10,10s4.5,10,10,10h3c5.5,0,10-4.5,10-10S43.5,13,38,13z M38,29h-3 c-3.3,0-6-2.7-6-6s2.7-6,6-6h3c3.3,0,6,2.7,6,6S41.3,29,38,29z"
              ></path>
              <path
                d="M13,13h-3C4.5,13,0,17.5,0,23s4.5,10,10,10h3c5.5,0,10-4.5,10-10S18.5,13,13,13z M13,29h-3 c-3.3,0-6-2.7-6-6s2.7-6,6-6h3c3.3,0,6,2.7,6,6S16.3,29,13,29z"
              ></path>
            </g>
            <path
              fill="#42A5F5"
              d="M33,21H15c-1.1,0-2,0.9-2,2s0.9,2,2,2h18c1.1,0,2-0.9,2-2S34.1,21,33,21z"
            ></path>
          </svg>
          <div class="item-desc">连接</div>
        </a>
        <a
          class="sider-item"
          :class="{ selectside: currIndex === 4 }"
          @click="onJump(4)"
        >
          <svg
            stroke="currentColor"
            fill="currentColor"
            stroke-width="0"
            version="1"
            viewBox="0 0 48 48"
            enable-background="new 0 0 48 48"
            height="1em"
            width="1em"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              fill="#607D8B"
              d="M39.6,27.2c0.1-0.7,0.2-1.4,0.2-2.2s-0.1-1.5-0.2-2.2l4.5-3.2c0.4-0.3,0.6-0.9,0.3-1.4L40,10.8 c-0.3-0.5-0.8-0.7-1.3-0.4l-5,2.3c-1.2-0.9-2.4-1.6-3.8-2.2l-0.5-5.5c-0.1-0.5-0.5-0.9-1-0.9h-8.6c-0.5,0-1,0.4-1,0.9l-0.5,5.5 c-1.4,0.6-2.7,1.3-3.8,2.2l-5-2.3c-0.5-0.2-1.1,0-1.3,0.4l-4.3,7.4c-0.3,0.5-0.1,1.1,0.3,1.4l4.5,3.2c-0.1,0.7-0.2,1.4-0.2,2.2 s0.1,1.5,0.2,2.2L4,30.4c-0.4,0.3-0.6,0.9-0.3,1.4L8,39.2c0.3,0.5,0.8,0.7,1.3,0.4l5-2.3c1.2,0.9,2.4,1.6,3.8,2.2l0.5,5.5 c0.1,0.5,0.5,0.9,1,0.9h8.6c0.5,0,1-0.4,1-0.9l0.5-5.5c1.4-0.6,2.7-1.3,3.8-2.2l5,2.3c0.5,0.2,1.1,0,1.3-0.4l4.3-7.4 c0.3-0.5,0.1-1.1-0.3-1.4L39.6,27.2z M24,35c-5.5,0-10-4.5-10-10c0-5.5,4.5-10,10-10c5.5,0,10,4.5,10,10C34,30.5,29.5,35,24,35z"
            ></path>
            <path
              fill="#455A64"
              d="M24,13c-6.6,0-12,5.4-12,12c0,6.6,5.4,12,12,12s12-5.4,12-12C36,18.4,30.6,13,24,13z M24,30 c-2.8,0-5-2.2-5-5c0-2.8,2.2-5,5-5s5,2.2,5,5C29,27.8,26.8,30,24,30z"
            ></path>
          </svg>
          <div class="item-desc">设置</div>
        </a>
        <a
          class="sider-item"
          :class="{ selectside: currIndex === 5 }"
          @click="onJump(5)"
        >
          <svg
            stroke="currentColor"
            fill="currentColor"
            stroke-width="0"
            version="1"
            viewBox="0 0 48 48"
            enable-background="new 0 0 48 48"
            height="1em"
            width="1em"
            xmlns="http://www.w3.org/2000/svg"
          >
            <polygon
              fill="#90CAF9"
              points="40,45 8,45 8,3 30,3 40,13"
            ></polygon>
            <polygon fill="#E1F5FE" points="38.5,14 29,14 29,4.5"></polygon>
            <g fill="#1976D2">
              <rect x="16" y="21" width="17" height="2"></rect>
              <rect x="16" y="25" width="13" height="2"></rect>
              <rect x="16" y="29" width="17" height="2"></rect>
              <rect x="16" y="33" width="13" height="2"></rect>
            </g>
          </svg>
          <div class="item-desc">日志</div>
        </a>
      </div>
      <!-- 按钮 -->
      <div class="sider-buttons">
        <button
          @click="onThemeChange"
          class="_11NITPe4W7 _1kzEILbWHq"
          data-state="tooltip-hidden"
          data-reach-tooltip-trigger=""
        >
          <svg
            v-if="!theme"
            xmlns="http://www.w3.org/2000/svg"
            width="20"
            height="20"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          >
            <circle cx="12" cy="12" r="5"></circle>
            <g
              transform-origin="12px 12px"
              style="transform: none; transform-origin: 12px 12px"
            >
              <line x1="12" y1="1" x2="12" y2="3"></line>
              <line x1="12" y1="21" x2="12" y2="23"></line>
              <line x1="4.22" y1="4.22" x2="5.64" y2="5.64"></line>
              <line x1="18.36" y1="18.36" x2="19.78" y2="19.78"></line>
              <line x1="1" y1="12" x2="3" y2="12"></line>
              <line x1="21" y1="12" x2="23" y2="12"></line>
              <line x1="4.22" y1="19.78" x2="5.64" y2="18.36"></line>
              <line x1="18.36" y1="5.64" x2="19.78" y2="4.22"></line>
            </g>
          </svg>
          <svg
            v-if="theme"
            xmlns="http://www.w3.org/2000/svg"
            width="20"
            height="20"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          >
            <path
              d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"
              transform-origin="12.019092798233032px 11.98090648651123px"
              style="transform: none; transform-origin: 12.0191px 11.9809px"
            ></path>
          </svg>
        </button>
        <a
          class="_11NITPe4W7"
          data-state="tooltip-hidden"
          data-reach-tooltip-trigger=""
          @click="onJump(6)"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="20"
            height="20"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          >
            <circle cx="12" cy="12" r="10"></circle>
            <line x1="12" y1="16" x2="12" y2="12"></line>
            <line x1="12" y1="8" x2="12.01" y2="8"></line>
          </svg>
        </a>
      </div>
      <!-- 退出 -->
      <div class="sider-exit">
        <!-- <div @click="onExit" class="exit-text">退出登录</div> -->
        <el-popover
          :popper-class="themeClass"
          placement="top"
          :width="160"
          v-model:visible="visibleExit"
        >
          <p>确定要退出吗？</p>
          <div style="height: 10px"></div>
          <div style="text-align: right; margin: 0">
            <el-button size="mini" type="text" @click="visibleExit = false"
              >取消</el-button
            >
            <el-button type="primary" size="mini" @click="onExit"
              >确定</el-button
            >
          </div>
          <template #reference>
            <div @click="visibleExit = true" class="exit-text">退出登录</div>
          </template>
        </el-popover>
      </div>
    </div>
    <!-- 嵌套路由 -->
    <div class="app-main">
      <div class="_3WrcX9K2WQ">
        <div class="_24ddJm1Q5s">
          <h1 class="B4QNkMu-0t">{{ navName }}</h1>
        </div>
      </div>

      <router-view></router-view>
    </div>
  </div>

  <!-- <div class="exit-comfirm">asdasd</div> -->
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

import logo_t from '/src/assets/login-bg.png'
import { ElNotification } from 'element-plus'

import { getVersion } from '../api/public'

import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import { key } from '../store'

export default defineComponent({
  data() {
    return {}
  },
  setup() {
    // get store using inject store key
    const store = useStore(key)

    // 主题
    let theme = ref(true)
    let visibleExit = ref(false)
    const router = useRouter()

    function onExit() {
      router.replace({
        name: 'Login',
      })
    }

    let themeClass = ref('light')

    // false light
    // true dark
    function onThemeChange() {
      theme.value = !theme.value
      themeClass.value = theme.value ? 'dark' : 'light'
      store.commit('setTheme', theme.value)
    }

    let currIndex = ref(-1)
    let navName = ref('概况')

    function onJump(index) {
      switch (index) {
        case 0:
          router.push('/center/dashboard')
          currIndex.value = 0
          navName.value = '仪表盘'
          break
        case 1:
          router.push('/center/global')
          currIndex.value = 1
          break
        case 2:
          router.push('/center/measure')
          currIndex.value = 2
          navName.value = '主页规划'
          break
        case 3:
          router.push('/center/connection')
          navName.value = '用户连接'
          currIndex.value = 3
          break
        case 4:
          router.push('/center/setting')
          currIndex.value = 4
          break
        case 5:
          router.push('/center/log')
          currIndex.value = 5
          break
        case 6:
          router.push('/center/about')
          currIndex.value = -1
          navName.value = '关于'
          break
      }
    }

    return {
      theme,
      onExit,
      visibleExit,
      onThemeChange,
      navName,
      onJump,
      currIndex,
      themeClass,
    }
  },
})
</script>

<style>
html {
  overflow-x: hidden;
}

/* #app,
.mHome,
.el-container {
  padding: 0px;
  margin: 0px;
  height: 100%;
} */

#app {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}

.el-aside {
  height: 100vh;
}

body {
  background-color: aliceblue;
  background: aliceblue;
  font-family: Open Sans, -apple-system, BlinkMacSystemFont, Segoe UI, Roboto,
    Helvetica, Arial, sans-serif, Apple Color Emoji, Segoe UI Emoji,
    Segoe UI Symbol, PingFang SC, Microsoft YaHei, 微软雅黑;
  -webkit-tap-highlight-color: rgba(0, 0, 0, 0);
  -webkit-text-size-adjust: 100%;
  -webkit-font-smoothing: antialiased;
}
</style>

<style>
.devapp {
  position: fixed;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  display: flex;
  background: var(--color-background);
  color: var(--color-text);
}

.app-sider {
  min-width: 150px;
  position: relative;
  box-sizing: border-box;
  background: var(--color-bg-sidebar);
}

.app-main {
  flex-grow: 1;
  overflow: auto;
}

.sider-content {
}

.sider-item {
  color: var(--color-text);
  text-decoration: none;
  display: flex;
  align-items: center;
  padding: 8px 20px;
}

*,
:after,
:before {
  box-sizing: border-box;
}

.item-desc {
  font-size: 1em;
  padding-left: 14px;
}

svg {
  width: 24px;
  height: 24px;
}

.sider-buttons {
  position: absolute;
  bottom: 50px;
  left: 50%;
  transform: translateX(-50%);
}

.sider-exit {
  height: 30px;
  width: 120px;
  left: 15px;
  background-color: var(--color-background);
  position: absolute;
  bottom: 10px;
  border-radius: 5px;
}

.sider-exit:hover {
  background-color: var(--color-sb-active-row-bg);
}

.exit-text {
  text-align: center;
  font-weight: bold;
  margin-top: 4px;
}
</style>

<style>
.light {
  --color-background: #fbfbfb;
  --color-background2: hsla(0, 0%, 98.4%, 0.3);
  --color-text: #222;
  --color-text-secondary: #646464;
  --color-text-highlight: #040404;
  --color-bg-sidebar: #e7e7e7;
  --color-sb-active-row-bg: #d0d0d0;
  --color-input-bg: #fff;
  --color-input-border: silver;
  --color-toggle-bg: #fff;
  --color-toggle-selected: #d7d7d7;
  --color-icon: #5b5b5b;
  --color-separator: #ccc;
  --color-btn-bg: #f4f4f4;
  --color-btn-fg: #101010;
  --color-bg-proxy: #e7e7e7;
  --color-row-odd: #f5f5f5;
  --bg-modal: #fbfbfb;
  --bg-near-transparent: rgba(0, 0, 0, 0.1);
  --bg-tooltip: #f0f0f0;
  --bc-tooltip: #ccc;
  --select-border-color: #999;
  --select-bg-hover: url(
    data:image/svg + xml;charset=utf-8,
    %3Csvgwidth='8'height='24'fill='none'xmlns='http://www.w3.org/2000/svg'%3E%3Cpathd='M4 7l3 4H1l3-4zM4 17l-3-4h6l-3 4z'fill='%23222'/%3E%3C/svg%3E
  );

  --color-focus-blue: #1a73e8;
  --btn-bg: #387cec;
}

.light {
  --loading-dot-1-1: rgba(0, 0, 0, 0.1);
  --loading-dot-1-2: rgba(0, 0, 0, 0.5);
  --loading-dot-1-3: rgba(0, 0, 0, 0.3);
  --loading-dot-2-1: rgba(0, 0, 0, 0.3);
  --loading-dot-2-2: rgba(0, 0, 0, 0.1);
  --loading-dot-2-3: rgba(0, 0, 0, 0.5);
  --loading-dot-3-1: rgba(0, 0, 0, 0.5);
  --loading-dot-3-2: rgba(0, 0, 0, 0.3);
  --loading-dot-3-3: rgba(0, 0, 0, 0.1);
}

.dark {
  --loading-dot-1-1: hsla(0, 0%, 100%, 0.5);
  --loading-dot-1-2: hsla(0, 0%, 100%, 0.1);
  --loading-dot-1-3: hsla(0, 0%, 100%, 0.3);
  --loading-dot-2-1: hsla(0, 0%, 100%, 0.3);
  --loading-dot-2-2: hsla(0, 0%, 100%, 0.5);
  --loading-dot-2-3: hsla(0, 0%, 100%, 0.1);
  --loading-dot-3-1: hsla(0, 0%, 100%, 0.1);
  --loading-dot-3-2: hsla(0, 0%, 100%, 0.3);
  --loading-dot-3-3: hsla(0, 0%, 100%, 0.5);
}

.dark {
  --color-background: #202020;
  --color-background2: rgba(32, 32, 32, 0.3);
  --color-text: #ddd;
  --color-text-secondary: #ccc;
  --color-text-highlight: #fff;
  --color-bg-sidebar: #2d2d30;
  --color-sb-active-row-bg: #494b4e;
  --color-input-bg: #2d2d30;
  --color-input-border: #3f3f3f;
  --color-toggle-bg: #353535;
  --color-toggle-selected: #181818;
  --color-icon: #c7c7c7;
  --color-separator: #333;
  --color-btn-bg: #232323;
  --color-btn-fg: #bebebe;
  --color-bg-proxy: #303030;
  --color-row-odd: #282828;
  --bg-modal: #1f1f20;
  --bg-near-transparent: hsla(0, 0%, 100%, 0.1);
  --bg-tooltip: #111;
  --bc-tooltip: #555;
  --select-border-color: #040404;
  --select-bg-hover: url(
    data:image/svg + xml;charset=utf-8,
    %3Csvgwidth='8'height='24'fill='none'xmlns='http://www.w3.org/2000/svg'%3E%3Cpathd='M4 7l3 4H1l3-4zM4 17l-3-4h6l-3 4z'fill='%23fff'/%3E%3C/svg%3E
  );

  --color-focus-blue: #1a73e8;
  --btn-bg: #387cec;
}

/* 定制弹出层 */
.el-popper.is-light {
  background-color: #494b4e;
  border: 0px solid #e4e7ed;
}

.el-popper.is-light .el-popper__arrow::before {
  border: 0px solid #e4e7ed;
  background: rgb(73, 75, 78);
  right: 0;
}
.el-popper p {
  color: white;
}

.builtState[data-v-7ba5bd90] {
  position: fixed;
  float: left;
  bottom: 4px;
  right: 16px;
  font-size: 18px;
  color: #646464;
  font-weight: 600;
}
</style>

<!-- 混淆 css -->
<style>
._1kzEILbWHq {
  -webkit-appearance: none;
  -moz-appearance: none;
  appearance: none;
  -webkit-user-select: none;
  user-select: none;
  background: none;
  cursor: pointer;
}
._11NITPe4W7 {
  --sz: 40px;
  width: var(--sz);
  height: var(--sz);
  display: flex;
  justify-content: center;
  align-items: center;
  outline: none;
  padding: 5px;
  color: var(--color-text);
  border-radius: 100%;
  border: 1px solid transparent;
}
._11NITPe4W7 {
  --sz: 40px;
  width: var(--sz);
  height: var(--sz);
  display: flex;
  justify-content: center;
  align-items: center;
  outline: none;
  padding: 5px;
  color: var(--color-text);
  border-radius: 100%;
  border: 1px solid transparent;
}

._11NITPe4W7:focus {
  border-color: var(--color-focus-blue);
}

._11NITPe4W7:hover {
  opacity: 0.6;
}

.selectside {
  background: var(--color-sb-active-row-bg);
}

/* 导航头部设置 */
._24ddJm1Q5s {
  height: 76px;
  display: flex;
  align-items: center;
}
._3WrcX9K2WQ {
  display: grid;
  grid-template-columns: 1fr minmax(auto, 330px);
  align-items: center;
  padding-right: 40px;
  background-color: var(--color-row-odd);
  box-shadow: 0 0 6px rgb(0 0 0 / 12%);
  /* bilibili阴影 */
}
.B4QNkMu-0t {
  padding: 0 40px;
  font-size: 2em;
}
</style>

<!-- exit -->
<style>
.exit-comfirm {
  width: 200px;
  height: 150px;
  position: relative;
  z-index: 200;
  background-color: coral;
  margin: 0 auto;
  top: 180px;
}
</style>

<style scoped>
a {
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}
</style>
