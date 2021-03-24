<template>
  <div class="item-card">
    <el-skeleton :loading="loading" animated>
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
        <div class="card-content">
          <img ref="imgdom" class="item-card-img" />
          <div class="item-card-title">原神</div>
          <div class="item-card-action">
            <button class="action-btn">启动</button>
          </div>
        </div>
      </template>
    </el-skeleton>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted, toRefs } from 'vue'
import { initImage } from './item.ts'

export default defineComponent({
  data() {
    return {
      // loading: true,
      fits: ['fill', 'contain', 'cover', 'none', 'scale-down'],
      url: 'yuanshen1.png',
      // base: import.meta.env.VITE_STATIC_URL,
      base: '',
    }
  },

  props: {
    track: { type: String },
    ext: { type: String },
  },
  setup(props) {
    // 图片加载
    const imgdom = ref(null)
    const loading = ref(true)

    const { track, ext } = toRefs(props)
    const internalInitPic = () => {
      imgdom.value.src = track.value + '.' + ext.value
    }

    // 完成 callback

    function a() {}
    let pd

    const pre = () => {
      imgdom.value.onload = handleLoad()
      function handleLoad(e) {
        console.log('start', imgdom.value.complete)
        pd = setInterval(() => {
          console.log('process', imgdom.value.complete)
          if (imgdom.value.complete) {
            clearInterval(pd)
            loading.value = false
            console.log('ok')
          }
          // console.log(imgdom.value.complete)
        }, 50)
      }

      imgdom.value.onerror = handleError()
      function handleError(e) {
        loading.value = true
        console.log('error')
        clearInterval(pd)
      }
    }
    onMounted(pre)

    onMounted(internalInitPic)
    return {
      imgdom,
      loading,
    }
  },
})
</script>

<style>
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
  height: 240px;
  border-radius: 6px;
  color: aliceblue;

  text-align: center;
}

.item-card-img {
  width: 100px;
  height: 100px;
  margin-top: 35px;
}

.item-card-title {
  margin-top: 15px;
  font-size: 18px;
  font-weight: 500;
}

.item-card-action button {
  margin-top: 10px;
  background-color: #00a1d5;
}
</style>
