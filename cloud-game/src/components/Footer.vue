<template>
  <div class="content">
    <div class="policy">
      <a
        class="link"
        id="privacy_policy_terms"
        href="https://github.com/r3inbowari/AndroidRemote"
        target="_blank"
        ><strong style="font-family: 'Noto Sans', sans-serif"
          >r3inbowari@gmail.com &amp; 使用条款
        </strong>
      </a>
    </div>

    <div class="change_view">
      <a class="go_mobile link_change_platform" @click="onGithub">前往源代码</a>
    </div>

    <div class="img"></div>

    <div style="color: #b0baca">
      <div style="text-align: center">r3inbowari @2021 bye designed</div>
      <div style="text-align: center">powered by aliyun</div>
    </div>
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
  name: 'Footer',
  components: {},
  props: {},
  setup(props) {
    function onGithub() {
      window.open('https://github.com/r3inbowari/AndroidRemote')
    }
    return { onGithub }
  },
})
</script>

<style scoped>
* {
  margin: 0;
  padding: 0;
  vertical-align: baseline;
  border: 0;
  font-family: 'Noto Sans', sans-serif;
}

.content {
  width: 1240px;
  position: relative;
  margin: 0 auto;
  padding: 50px 20px;
  -webkit-box-sizing: border-box;
  box-sizing: border-box;
}

.policy {
  width: 100%;
  display: block;
  position: relative;
  float: left;
  padding: 14px 0;
}

.policy .link:last-child {
  margin-right: 0;
}
.footer .policy .link {
  margin-right: 10px;
  padding: 9px 15px;
  font-size: 13px;
  line-height: 1.38px;
  color: #b0baca;
  background-color: #20202c;
  border-radius: 10px;
  -webkit-transition: background-color 0.2s;
  transition: background-color 0.2s;
}

.policy .link:hover {
  background-color: #363643;
}

.change_view {
  width: 100%;
  display: block;
  margin-top: 80px;
  padding-bottom: 50px;
}

.go_mobile {
  display: inline-block;
  padding: 9px 15px;
  font-size: 13px;
  line-height: 18px;
  color: #b0baca;
  vertical-align: middle;
  border: 1px solid #b0baca;
  border-radius: 10px;
  -webkit-transition: all 0.5s;
  transition: all 0.5s;
}

.change_view:after {
  display: table;
  clear: both;
  content: '';
}

.go_mobile:hover {
  color: #fff;
  border: 1px solid #fff;
}

.img {
  z-index: 500;
  position: absolute;
  top: 70px;
  right: 60px;
  width: 200px;
  height: 100px;
  background-image: url('../assets/footer-img.png');
  background-size: cover;
  background-repeat: no-repeat;
  background-position: center;
}
</style>
