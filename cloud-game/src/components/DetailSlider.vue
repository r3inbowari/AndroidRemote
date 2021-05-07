<template>
  <div class="long-slider">
    <swiper :slidesPerView="'auto'" :spaceBetween="0" class="mySwiper">
      <swiper-slide v-for="(item, index) in slideItems">
        <img class="img" :src="item" alt="" />
      </swiper-slide>
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

export default defineComponent({
  name: 'GameSlider',
  components: { Swiper, SwiperSlide },
  setup() {
    let slideItems = reactive([
      'https://image-glb.qpyou.cn/hubweb/hive_img/A/A/1317/20210428/8fec9be66af60b0426fbecc952ee2c74.jpg',
      'https://image-glb.qpyou.cn/hubweb/hive_img/A/A/1317/20210428/5c916d8f50702ee09d61448adea9d5df.jpg',
      'https://image-glb.qpyou.cn/hubweb/hive_img/A/A/1317/20210428/f4ecbd05e2eb7489552ddd2e87f39164.jpg',
      'https://image-glb.qpyou.cn/hubweb/hive_img/A/A/1317/20210428/74353b21935b0afd3804ef30a96d338f.jpg',
      'https://image-glb.qpyou.cn/hubweb/hive_img/A/A/1317/20210428/60788e41ff8eff4a824c92af2c68d61c.jpg',
      'https://image-glb.qpyou.cn/hubweb/hive_img/A/A/1317/20210428/74353b21935b0afd3804ef30a96d338f.jpg',
      'https://image-glb.qpyou.cn/hubweb/hive_img/A/A/1317/20210428/60788e41ff8eff4a824c92af2c68d61c.jpg',
      'https://image-glb.qpyou.cn/hubweb/hive_img/A/A/1317/20210428/74353b21935b0afd3804ef30a96d338f.jpg',
      'https://image-glb.qpyou.cn/hubweb/hive_img/A/A/1317/20210428/60788e41ff8eff4a824c92af2c68d61c.jpg',
    ])

    return { slideItems }
  },
})
</script>

<style>
.swiper-slide {
  width: 48%;
}

.swiper-container {
  width: 100%;
  height: 100%;
}

.img {
  width: 551px;
  height: 310px;
  display: block;
  position: relative;
  overflow: hidden;
  margin: 0 0 30px 20px;
  background-color: #000;
  background-size: 100% 100%;
  border-radius: 20px;
}
</style>
