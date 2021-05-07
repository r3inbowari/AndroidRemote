<template>
  <div class="long-slider">
    <swiper :slidesPerView="'auto'" :spaceBetween="0" class="mySwiper">
      <swiper-slide v-for="(item, index) in slideItems"
        ><WideCard :info="slideItems[index]"></WideCard
      ></swiper-slide>
    </swiper>
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

import SwiperCore, { Navigation, Pagination } from 'swiper'
import { Swiper, SwiperSlide } from 'swiper/vue'
import 'swiper/swiper-bundle.css'
SwiperCore.use([Navigation, Pagination])

import WideCard from './WideCard.vue'

export default defineComponent({
  name: 'GameSlider',
  components: { Swiper, SwiperSlide, WideCard },
  setup() {
    let slideItems = reactive([
      {
        title: '内核活动0',
        alias: 'kernel event',
        cover:
          'https://image-glb.qpyou.cn/hubweb/hive_img/web/banner/20200317/41a9e65126c0463ad975ac444027bfa3_1200x490.jpg',
        icon:
          'https://image-glb.qpyou.cn/hubweb/gmnotice/appcenter/1606972197901.jpg',
        type: 'RPG',
        tags: ['app', '最受喜爱'],
      },
      {
        title: '内核活动1',
        alias: 'kernel event',
        cover:
          'https://image-glb.qpyou.cn/hubweb/hive_img/web/banner/20201203/00d2a88aeb6037a2ada2ffd1ad703e61_1200x490.jpg',
        icon:
          'https://image-glb.qpyou.cn/hubweb/gmnotice/appcenter/1560919617216.jpg',
        type: 'RPG',
        tags: ['app', '最受喜爱'],
      },
      {
        title: '内核活动2',
        alias: 'kernel event',
        cover:
          'https://image-glb.qpyou.cn/hubweb/hive_img/web/banner/20210204/0aa3bbb693b4d1b9cb6fc61ce9d9ac8e_1200x490.jpg',
        icon:
          'https://image-glb.qpyou.cn/hubweb/gmnotice/appcenter/1612412918869.jpg',
        type: 'RPG',
        tags: ['app', '最受喜爱'],
      },
      {
        title: '内核活动2',
        alias: 'kernel event',
        cover:
          'https://image-glb.qpyou.cn/hubweb/hive_img/web/banner/20210204/0aa3bbb693b4d1b9cb6fc61ce9d9ac8e_1200x490.jpg',
        icon:
          'https://image-glb.qpyou.cn/hubweb/gmnotice/appcenter/1612412918869.jpg',
        type: 'RPG',
        tags: ['app', '最受喜爱'],
      },
    ])

    return { slideItems }
  },
})
</script>

<style>
.swiper-slide {
  width: 30%;
}

.swiper-container {
  width: 100%;
  height: 100%;
}
</style>
