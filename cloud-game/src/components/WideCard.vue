<template>
  <div class="wide-card">
    <a href="">
      <div class="background">
        <div class="img" :style="imgBackground"></div>
      </div>

      <p class="tag">
        <span class="box" v-for="(item, index) in info.tags">
          #{{ info.tags[index] }}
        </span>
        <!-- <span class="box">#{{ info.tags }}</span> -->
      </p>

      <div class="icon" :style="iconBackground"></div>
      <p class="title">{{ info.title }}</p>
      <p class="genre">{{ info.type }}</p>
    </a>
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
  toRefs,
} from 'vue'

export default defineComponent({
  name: 'GameSliderCard',
  components: {},
  props: {
    info: {
      type: Object,
      default: {
        title: '加载中',
        alias: 'Loading',
        type: 'Loading',
        tags: ['默认'],
        cover: '',
        icon: '',
      },
    },
  },
  setup(props) {
    let imgBackground = ref('')
    let iconBackground = ref('')

    /**
       * 虚化背景替换
       * @param r 操作对象
       * @param url 传入背景地址
       * @example changeBackground(
            bannerBackground,
            'https://image-glb.qpyou.cn/hubweb/hive_img/web/banner/20201203/00d2a88aeb6037a2ada2ffd1ad703e61_1200x490.jpg'
          )
       **/
    function changeBackground(r, url) {
      r.value = 'background-image: url("' + url + '")'
    }

    // 响应性获取
    const { info } = toRefs(props)

    onMounted(() => {
      // 写入背景
      if (info.value.cover !== undefined) {
        changeBackground(imgBackground, info.value.cover)
      }
      if (info.value.icon !== undefined) {
        changeBackground(iconBackground, info.value.icon)
      }
    })

    // watch(info, (newProps, oldProps) => {
    //   console.log('asd')
    //   changeBackground(imgBackground, newProps.cover)
    // })

    watch(
      () => props.info.value,
      (first, second) => {
        console.log(
          'Watch props.selected function called with args:',
          first,
          second
        )
        changeBackground(imgBackground, info.value.cover)
      }
    )

    return { imgBackground, iconBackground }
  },
})
</script>

<style></style>

<style scoped>
* {
  margin: 0;
  padding: 0;
  vertical-align: baseline;
  border: 0;
  font-family: 'Noto Sans', sans-serif;
}

.wide-card {
  width: 328px;
  height: 210px;
  margin-top: 20px;

  /* background-color: lawngreen; */
}

.background {
  width: 328px;
  height: 210px;
  position: relative;
  overflow: hidden;
  background-color: #000;
  border-radius: 20px;
  -webkit-transition: box-shadow 0.2s;
  transition: box-shadow 0.2s;
  -webkit-mask-image: -webkit-radial-gradient(white, black);
}

.background .img {
  width: 100%;
  height: 100%;
  position: absolute;
  top: 0;
  left: 0;
  background-position: 50% 50%;
  background-size: auto 100%;
  transition: transform 0.6s cubic-bezier(0.23, 1, 0.32, 1);
  transform: scale(1) translateZ(0);
}

.background:after {
  width: 100%;
  height: 105px;
  position: absolute;
  bottom: 0;
  content: '';
  background-image: linear-gradient(
    to bottom,
    rgba(10, 10, 18, 0) 0%,
    rgba(10, 10, 18, 0.8)
  );
  border-bottom-right-radius: 20px;
  border-bottom-left-radius: 20px;
}

.swiper-slide {
  width: 350px;
}

a:hover .background .img {
  transition: transform 0.6s cubic-bezier(0.23, 1, 0.32, 1);
  transform: scale(1.1) translateZ(0);
}

.tag {
  position: absolute;
  top: 35px;
  left: 15px;
}

.tag .box {
  margin-right: 6px;
  padding: 6px 12px;
  font-size: 11px;
  font-weight: bold;
  line-height: 1.55;
  color: #fff;
  background-color: #1277ff;
  border-radius: 19px;
}

.icon {
  width: 60px;
  height: 60px;
  position: absolute;
  overflow: hidden;
  top: 155px;
  left: 15px;
  background-color: #000;
  background-size: 100% 100%;
  border-radius: 20px;
}

.title {
  position: absolute;
  top: 164px;
  left: 85px;
  font-size: 15px;
  font-weight: bold;
  line-height: 22px;
  color: #fff;
  font-family: 'Noto Sans SC', sans-serif;
}

.genre {
  position: absolute;
  top: 188px;
  left: 85px;
  font-size: 13px;
  line-height: 18px;
  color: #1277ff;
  font-family: 'Noto Sans SC', sans-serif;
}
</style>
