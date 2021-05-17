<template>
  <div class="other-slider">
    <swiper :slidesPerView="'auto'" :spaceBetween="0" class="mySwiper">
      <swiper-slide v-for="(item, index) in slideItems.dat">
        <a href="/games/313" tabindex="0">
          <img class="imgIcon" :src="item.icon" alt="" />
          <p class="game">{{ item.title }}</p>
          <p class="genre">{{ item.type }}</p>
        </a>
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
  name: 'OtherSlider',
  components: { Swiper, SwiperSlide },
  setup() {
    let slideItems = reactive({
      dat: [
        {
          icon: 'https://z3.ax1x.com/2021/05/17/gR0BsU.jpg',
          type: 'RPG',
          title: '梦幻西游',
        },
        {
          icon: 'https://z3.ax1x.com/2021/05/17/gRB5pq.jpg',
          type: 'MOBA',
          title: '王者荣耀',
        },
      ],
    })

    return { slideItems }
  },
})
</script>

<style scoped>
.swiper-slide {
  width: 16%;
}

.swiper-container {
  width: 100%;
  height: 100%;
}

.imgIcon {
  width: 154px;
  height: 154px;
  position: relative;
  overflow: hidden;
  background-color: #000;
  background-size: 100% 100%;
  border-radius: 30px;
  -webkit-transition: box-shadow 0.2s;
  transition: box-shadow 0.2s;
}

.genre {
  width: 154px;
  height: 20px;
  overflow: hidden;
  font-size: 14px;
  line-height: 20px;
  color: #1277ff;
  text-overflow: ellipsis;
  word-wrap: normal !important;
  white-space: nowrap;
  -o-text-overflow: ellipsis;
}

.game {
  width: 154px;
  height: 22px;
  overflow: hidden;
  margin-top: 15px;
  margin-bottom: 4px;
  font-size: 16px;
  font-weight: bold;
  line-height: 24px;
  color: #000f45;
  text-overflow: ellipsis;
  word-wrap: normal !important;
  white-space: nowrap;
  -o-text-overflow: ellipsis;
}

a {
  width: 154px;
  display: block;
  margin: 0 0 0 20px;
  -webkit-box-sizing: border-box;
  box-sizing: border-box;
  -webkit-transition: background-color 0.5s;
  transition: background-color 0.5s;
}

.other-slider {
  margin-top: 20px !important;
}
</style>
