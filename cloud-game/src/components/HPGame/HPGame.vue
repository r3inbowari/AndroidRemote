<template>
  <div
    class="game-list-display"
    :class="[
      { 'game-list-display-1': background1 === 1 },
      { 'game-list-display-2': background1 === 2 },
      { 'game-list-display-3': background1 === 3 },
      { 'game-list-display-4': background1 === 4 },
      { 'game-list-display-5': background1 === 5 },
      { 'game-list-display-6': background1 === 6 },
    ]"
  >
    <!-- <ItemCard class="display-item" :dat="dat" @run="onRun"> </ItemCard> -->

    <div class="display-list">
      <div class="type-content">
        <div
          style="
            background: rgb(0, 161, 213);
            width: 8px;
            height: 40px;
            margin-right: 10px;
          "
        ></div>
        <div class="type-title">???游戏</div>
      </div>
      <hr color="#00A1D5" />
      <el-row type="flex" class="row-bg">
        <el-col v-for="name in gameDatas" :span="spanLen" class="col-class"
          ><ItemCard class="display-item" :dat="name" @run="runApp"> </ItemCard
        ></el-col>
      </el-row>
      <!-- <ItemCard class="display-item" :dat="gameDatas[0]" @run="onb"> </ItemCard> -->
    </div>
  </div>
  <LoginDialog ref="nop2"></LoginDialog>
</template>

<script lang="ts">
import { defineComponent, onMounted } from 'vue'
import { printEnv } from './utils'
import ItemCard from './ItemCard.vue'

// 第二次引用 login 
import LoginDialog from '../Login/Login.vue'

// import HomePage from './components/HelloWorld.vue'

import { VueCookieNext } from 'vue-cookie-next'

export default defineComponent({
  name: 'App',
  components: { ItemCard, LoginDialog },
  data() {
    return {
      yuanshen: '',
      kkk: 1,
      // dat: {
      //   track: 'yuanshen.png',
      //   name: '原神',
      //   aid: '6c1de797-ce90-43db-a51d-c77c440792f6',
      // },
      gameDatas: [
        {
          track: 'yuanshen.png',
          name: '原神',
          aid: '6c1de797-ce90-43db-a51d-c77c440792f6',
        },
        {
          track: 'guangyu.png',
          name: '光遇',
          aid: '4d53d679-de16-49e6-9706-ddca21d6cb82',
        },
      ],
      spanLen: 8,
      screenWidth: document.body.clientWidth,
      background1: 0,
    }
  },
  setup() {

  },

  mounted() {
    const that = this
    window.onresize = () => {
      return (() => {
        window.screenWidth = document.body.clientWidth
        that.screenWidth = window.screenWidth
      })()
    }
    this.fix(this.screenWidth)
  },
  methods: {
    runApp(gameData) {
      let t = VueCookieNext.isCookieAvailable('token')
      // console.log(gameData)
      if (t) {
        console.log('ok');
        // queue task

      } else {
          this.$refs['nop2'].needLogin()
      }
    },
    fix(val) {
      if (val < 3000 && val > 1820) {
        this.background1 = 0
        this.spanLen = 4
      } else if (val < 1820 && val > 1720) {
        this.background1 = 1
        this.spanLen = 4
      } else if (val < 1720 && val > 1600) {
        this.background1 = 2
        this.spanLen = 4
      } else if (val < 1550 && val > 1200) {
        this.background1 = 3
        this.spanLen = 6
      } else if (val < 1150 && val > 970) {
        this.background1 = 4
        this.spanLen = 6
      } else if (val < 970 && val > 765) {
        this.background1 = 5
        this.spanLen = 8
      } else if (val < 765 && val > 600) {
        this.background1 = 6
        this.spanLen = 12
      }
    },
  },
  watch: {
    screenWidth(val) {
      this.screenWidth = val
      // console.log('this.screenWidth', this.screenWidth)
      this.fix(val)
    },
  },
})
</script>

<style>
.type-content {
  /* background-color: rgb(20, 20, 20); */
  margin-bottom: 20px;
  height: 40px;
  color: white;

  display: flex;
  justify-content: left;
  align-items: center;
  font-size: 26px;
  font-weight: bold;
}

.game-list-display {
  margin-top: 40px;
  /* width: 70%;
  left: 15%; */
  padding-left: 350px;
  padding-right: 350px;
  /* position: relative; */
  min-height: 1200px;
  /* margin-left: 15%; */
}

.display-list {
  /* text-align: center; */
}

.display-item {
  margin: 0 auto;
}

.game-list-display-1 {
  padding-left: 300px;
  padding-right: 300px;
}

.game-list-display-2 {
  padding-left: 220px;
  padding-right: 220px;
}

.game-list-display-3 {
  padding-left: 200px;
  padding-right: 200px;
}
.game-list-display-4 {
  padding-left: 100px;
  padding-right: 100px;
}

.game-list-display-5 {
  padding-left: 100px;
  padding-right: 100px;
}

.game-list-display-6 {
  padding-left: 50px;
  padding-right: 50px;
}

.col-class {
  margin-bottom: 20px;
}

.row-bg {
  background-color: rgb(32, 32, 32);
  padding-top: 20px;
  border-radius: 6px;
}
</style>
