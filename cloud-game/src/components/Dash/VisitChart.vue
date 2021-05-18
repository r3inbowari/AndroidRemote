<template>
  <div
    class="play-chart"
    style="height: 360px; width: 550px; display: flex; flex-direction: column"
  >
    <div class="chart-title">
      <h2>访问频次</h2>
    </div>
    <vue3-chart-js
      :id="playChart.id"
      ref="playChartRef"
      :type="playChart.type"
      :data="playChart.data"
      :options="playChart.options"
    >
    </vue3-chart-js>
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

import { getBanner, delBanner, addBanner } from '../api/banner'
import { getHot, delHot, addHot } from '../api/hot'
import { getUpdate, delUpdate, addUpdate } from '../api/update'

import { ElNotification } from 'element-plus'
import { ops } from '../utils'

import { useStore } from 'vuex'
import { key } from '../../store'

import Vue3ChartJs from '@j-t-mcc/vue3-chartjs'

export default defineComponent({
  data() {
    return {}
  },
  components: {
    Vue3ChartJs,
  },
  setup() {
    const store = useStore(key)

    const playChartRef = ref(null)
    let customClass = ref('dark')
    const playChart = reactive({
      type: 'line',
      data: {
        labels: [
          '6小时前',
          '6小时前',
          '4小时前',
          '3小时前',
          '2小时前',
          '1小时前',
        ],
        // labels: ['VueJs', 'EmberJs', 'ReactJs', 'AngularJs'],
        datasets: [
          {
            label: 'Caffine',
            data: [4, 23, 11, 2, 10, 6],
            fill: false,
            borderColor: '#41B883',
            backgroundColor: 'black',
            color: 'white',
            pointRadius: 0, // 隐藏point
          },
        ],
      },
      options: {
        plugins: {
          legend: {
            display: false, // 隐藏图例标签
          },
        },
        //  close animation
        animation: {
          duration: 0,
        },
        scales: {},
      },
    })

    // 主题监听
    watch(
      () => store.state.theme,
      (oldValue, newValue) => {
        customClass.value = store.state.theme ? 'dark' : 'light'
        // changeTheme(playChartRef, playChart)
      }
    )

    onMounted(() => {
      // setInterval(() => {
      //   pushDate2ChartByQueueWay(playChartRef, playChart, 20)
      // }, 2000)
    })

    function pushDate2ChartByQueueWay(ref, chartInstance, value) {
      chartInstance.data.datasets[0].data.splice(0, 1)
      chartInstance.data.datasets[0].data.push(value)
      ref.value.update(1)
    }

    return {
      playChart,
      playChartRef,
    }
  },
})
</script>

<style scoped>
.dashboard {
  padding: 30px;
}

.play-chart {
  background-color: var(--color-bg-sidebar);
  border-radius: 8px;
  padding: 10px;
}

.chart-title {
  height: 40px;
  /* background-color: aquamarine; */
  margin-bottom: 10px;
}

h2 {
  display: block;
  font-size: 1.5em;
  margin-block-start: 0.1em;
  margin-block-end: 0em;
  margin-inline-start: 0px;
  margin-inline-end: 0px;
  font-weight: bold;
}
</style>
