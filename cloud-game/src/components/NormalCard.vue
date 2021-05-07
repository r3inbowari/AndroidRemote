<template>
  <li>
    <a :href="href">
      <p class="tag">
        <span v-if="info.new" class="new">新作</span>
        <span v-if="info.recommend" class="recommend">推荐</span>
        <span v-if="info.popular" class="popular">人气</span>
        <span v-if="info.update" class="update">更新</span>
      </p>

      <div class="icon">
        <div class="img" :style="iconBackground"></div>
      </div>
      <p class="title">
        <i class="caution"></i>{{ info.title }}
        <span class="genre">{{ info.type }}</span>
      </p>
      <p class="desc">{{ info.desc }}</p>
    </a>
  </li>
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
  name: 'NormalCard',
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
        desc: '加载中',
        new: false,
        recommend: false,
        popular: false,
        update: false,
        aid: '6c1de797-ce90-43db-a51d-c77c440792f6',
      },
    },
  },
  setup(props) {
    // 响应性获取
    const { info } = toRefs(props)

    let iconBackground = ref('')

    let href = '/game/' + info.value.aid
    // console.log(href)

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

    onMounted(() => {
      changeBackground(iconBackground, info.value.icon)
    })
    return {
      href,
      iconBackground,
    }
  },
})
</script>

<style scoped>
a {
  width: 100%;
  height: 100%;
  display: block;
}
.tag {
  position: absolute;
  top: 16px;
  right: 20px;
}

.icon {
  width: 90px;
  height: 90px;
  position: absolute;
  overflow: hidden;
  top: 20px;
  left: 20px;
  background: none;
  border-radius: 30px;
  -webkit-mask-image: -webkit-radial-gradient(white, black);
}

.new {
  color: #00c6a5;
  background-color: #c2f4f8;
}

.recommend {
  color: #1277ff;
  background-color: #d4dfff;
}

.popular {
  color: #964aff;
  background-color: #ebdcff;
}

.update {
  color: #ff4758;
  background-color: #ffe4e6;
}

.title {
  position: absolute;
  top: 40px;
  left: 124px;
  font-size: 16px;
  font-weight: bold;
  line-height: 1.5;
  color: #000f45;
}

.img {
  width: 100%;
  height: 100%;
  position: absolute;
  top: 0;
  left: 0;
  background-position: 50% 50%;
  background-size: 100% 100%;
}

.genre {
  margin-left: 6px;
  font-size: 13px;
  font-weight: normal;
  line-height: 1.38;
  color: #1277ff;
  font-family: 'Noto Sans SC', sans-serif;
}

.desc {
  width: 445px;
  height: 30px;
  position: absolute;
  overflow: hidden;
  top: 65px;
  left: 124px;
  font-size: 13px;
  line-height: 30px;
  color: #7e8592;
  font-family: 'Noto Sans SC', sans-serif;
}

li {
  width: 589px;
  height: 130px;
  position: relative;
  float: left;
  margin-right: 20px;
  margin-bottom: 20px;
  background-color: #fff;
  border-radius: 20px;
  -webkit-transition: box-shadow 0.2s;
  transition: box-shadow 0.2s;
}

* {
  margin: 0;
  padding: 0;
  vertical-align: baseline;
  border: 0;
  font-family: 'Noto Sans', sans-serif;
}

.tag span {
  display: inline-block;
  margin: 0 3px;
  padding: 6px 12px;
  font-size: 11px;
  font-weight: bold;
  line-height: 1.55;
  border-radius: 14px;
}

li {
  display: list-item;
  text-align: -webkit-match-parent;
}

/* 偶次项靠右 */
.list li:nth-child(even) {
  margin-right: 0;
}
</style>
