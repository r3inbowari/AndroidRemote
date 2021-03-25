<template>
  <div
    id="data-123b1a"
    class="item-card"
    :class="{ 'item-card-background': background }"
  >
    <el-skeleton :loading="loading" animated :count="1">
      <template #template>
        <el-skeleton-item
          variant="h6"
          style="width: 100px; height: 100px; margin-top: 35px"
        />

        <div style="padding: 14px">
          <el-skeleton-item variant="h2" style="height: 24px; width: 80%" />
          <div
            style="
              display: flex;
              align-items: center;
              justify-items: space-between;
              margin-top: 16px;
              height: 16px;
            "
          >
            <el-skeleton-item
              variant="text"
              style="margin: 0 50px 0 50px; height: 26px; width: 60px"
            />
          </div>
        </div>
      </template>

      <template #default>
        <!-- <div v-show="isLoad" class="card-content">
          <img ref="imgdom" class="item-card-img" />
          <div class="item-card-title">原神</div>
          <div class="item-card-action">
            <button class="action-btn">启动</button>
          </div>
        </div> -->
      </template>
    </el-skeleton>

    <!-- skeleton在vue3下目测不是很兼容，无法在插槽未显示时安全获取到ref dom，所以没办法只能这样搞 -->
    <div class="default-card-content">
      <div v-show="!loading" class="card-content">
        <img ref="imgdom" class="item-card-img" />
        <div class="item-card-title">{{ dat.name }}</div>
        <div class="item-card-action">
          <button @click="onRun" class="action-btn">启动</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted, toRefs } from 'vue'
// import { initImage } from './item.ts'

/**
 * 游戏卡片组件
 * @emit runApp 运行按钮按钮按下 return { gameData }
 * @props gameData
 *        track  图片名
 *        ext    扩展名     required false
 *        name   游戏名
 *        aid    普通索引
 * @props background 是否开启卡片背景
 **/
export default defineComponent({
  data() {
    return {
      fits: ['fill', 'contain', 'cover', 'none', 'scale-down'],
      url: 'yuanshen1.png',
      // base: import.meta.env.VITE_STATIC_URL,
      base: '',
    }
  },

  props: {
    dat: {
      // define clearly
      track: { type: String },
      ext: { type: String },
      name: { type: String },
      aid: { type: String },
    },
    background: {
      type: Boolean,
      default: false,
    },
  },

  setup(props, { attrs, slots, emit }) {
    // get refs
    const { dat } = toRefs(props)
    const imgdom = ref(null)
    const loading = ref(true)

    const internalInitPic = () => {
      // 图片加载
      if (dat.value.ext === undefined) {
        imgdom.value.src = dat.value.track
      } else {
        imgdom.value.src = dat.value.track + '.' + dat.value.ext
      }
      let loadTag
      imgdom.value.onload = handleLoad
      imgdom.value.onerror = handleError
      function handleLoad(e) {
        // console.log(imgdom.value)
        // 有性能问题
        loadTag = setInterval(() => {
          if (imgdom.value.complete) {
            clearInterval(loadTag)
            loading.value = false
          }
        }, 50)
      }

      function handleError(e) {
        // 加载错误处理
        console.log(
          '[itemCard] image load error -> ' + track.value + '.' + ext.value
        )
        clearInterval(loadTag)
      }
    }
    onMounted(internalInitPic)
    return {
      imgdom,
      loading,
    }
  },
  // emits: ['runapp'],
  methods: {
    onRun() {
      // 统一交由上层处理
      console.log('123213132')
      this.$emit('run', this.dat)
    },
  },
})
</script>

<style>
/* 使用卡片背景样式 大小必须为 180x260 */
.item-card-background {
  background: url('card-bg.png') no-repeat !important;
}

.item-card .el-skeleton.is-animated .el-skeleton__item {
  /* background: linear-gradient(90deg, #f2f2f2 25%, #e6e6e6 37%, #f2f2f2 63%); */
  background: linear-gradient(90deg, #f2f2f2 25%, #e6e6e6 37%, #f2f2f2 63%);
  background-image: linear-gradient(
    90deg,
    rgb(28, 32, 34) 25%,
    rgb(24, 28, 30) 37%,
    rgb(28, 32, 34) 63%
  );
  background-position-x: initial;
  background-position-y: initial;
  background-size: 400% 100%;
  background-repeat-x: initial;
  background-repeat-y: initial;
  background-attachment: initial;
  background-origin: initial;
  background-clip: initial;
  background-color: initial;
}

.card-content .el-image__error,
.el-image__inner,
.el-image__placeholder {
  border-radius: 6px;
  background-color: rgb(28, 32, 34);
}

.item-card {
  background-color: rgb(20, 20, 20);
  width: 180px;
  height: 260px;
  border-radius: 6px;
  color: aliceblue;

  text-align: center;

  position: relative;
}

/* element + 的 骨架屏存在问题，迫不得已把 div 移出骨架屏 */
.default-card-content {
  position: absolute;
  padding: 30px;
}

.item-card-img {
  /* width: 100px;
  height: 100px; */

  width: 100%;
  height: auto;

  border-radius: 22%;
}

.item-card-title {
  margin-top: 13px;
  font-size: 18px;
  font-weight: 500;
}

.item-card-action button {
  margin-top: 10px;
  background-color: #00a1d5;
}
</style>
