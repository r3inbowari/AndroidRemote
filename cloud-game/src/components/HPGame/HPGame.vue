<template>
  <div class="game-list-display">
    <ItemCard :dat="dat" @run="onRun"> </ItemCard>
    <!-- <div class="display-list">
      <el-row :gutter="10">
        <el-col :xs="4" :sm="6" :md="4" :lg="3" :xl="1"
          ><div class="grid-content bg-purple"></div
        ></el-col>
        <el-col :xs="4" :sm="6" :md="8" :lg="9" :xl="11"
          ><div class="grid-content bg-purple-light"></div
        ></el-col>
        <el-col :xs="4" :sm="6" :md="8" :lg="9" :xl="11"
          ><div class="grid-content bg-purple"></div
        ></el-col>
        <el-col :xs="4" :sm="6" :md="4" :lg="3" :xl="1"
          ><div class="grid-content bg-purple-light"></div
        ></el-col>
      </el-row>
    </div> -->
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted } from 'vue'
import { printEnv } from './utils'
import ItemCard from './ItemCard.vue'

// import HomePage from './components/HelloWorld.vue'

export default defineComponent({
  name: 'App',
  components: { ItemCard },
  data() {
    return {
      yuanshen: '',
      kkk: 1,
      dat: {
        track: 'yuanshen.png',
        name: '原神',
        aid: '6c1de797-ce90-43db-a51d-c77c440792f6',
      },
    }
  },
  setup() {
    // const hello = () => {
    //   console.log('hahaha')
    // }
    // onMounted(hello)
  },

  methods: {
    onRun(gameData) {
      console.log(gameData.aid)
    },
  },

  mounted() {
    // setTimeout(() => {
    //   this.dat.name = ''
    // }, 4000)
    // setInterval(() => {
    //   this.kkk++
    // }, 2000)
  },
})
</script>

<style>
.el-row {
  margin-bottom: 20px;
  &:last-child {
    margin-bottom: 0;
  }
}
.el-col {
  border-radius: 4px;
}
.bg-purple-dark {
  background: #99a9bf;
}
.bg-purple {
  background: #d3dce6;
}
.bg-purple-light {
  background: #e5e9f2;
}
.grid-content {
  border-radius: 4px;
  min-height: 36px;
}
.row-bg {
  padding: 10px 0;
  background-color: #f9fafc;
}

.game-list-display {
  margin-top: 40px;
  width: 70%;
  left: 15%;
  /* position: relative; */
  min-height: 1200px;
  margin-left: 15%;
}
</style>
