<template>
  <div style="margin: 20px">
    <img :src="imgJiatu" alt="" />
  </div>
</template>

<script>
import {
  defineComponent,
  onMounted,
  getCurrentInstance,
  ref,
  reactive,
  watch,
} from 'vue'
import { getBanner } from '../api/banner'
import { getHot } from '../api/hot'
import { getUpdate } from '../api/update'

import { useStore } from 'vuex'
import { key } from '../store'
import { useRouter } from 'vue-router'

import { userLogin, userInfo } from '../api/user'

import { Toast } from 'vant'

import { VueCookieNext } from 'vue-cookie-next'

import imgJiatu from '../assets/jiatu.png'

export default defineComponent({
  components: {},
  setup() {
    const store = useStore(key)
    const router = useRouter()
    return { imgJiatu }
  },
})
</script>
