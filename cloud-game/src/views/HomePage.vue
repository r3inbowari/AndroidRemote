<template>
  <div id="internalHeader" class="internalHeader">
    <div class="page-header mini-type">
      <div class="page-header-content">
        <div class="nav-link">logo</div>
        <div class="nav-search-box">
          <MiniSearach></MiniSearach>
        </div>
        <div class="nav-user-center">center</div>
      </div>
    </div>
  </div>

  <div class="v-wrap">123123</div>
</template>

<script lang="ts">
import { Vue, Options } from 'vue-class-component'
import MiniSearach from '../components/MiniSearch/MiniSearch.vue'

// Component definition
@Options({
  // Define component options
  watch: {
    count: (value) => {
      console.log(value)
    },
  },
  components: {
    MiniSearach,
  },
})
export default class HomePage extends Vue {
  // The behavior in class is the same as the current
  count = 0

  increment() {
    this.count++
  }

  created() {
    console.log('asd')
  }
}
</script>

<style>
.internalHeader {
  min-width: 998;
  min-height: 56;

  z-index: 999;
  position: relative;
  margin: 0;
  padding: 0;
}

.internalHeader .mini-type {
  position: relative;
  box-shadow: 0 2px 4px 0 rgb(0 0 0 / 8%);
  box-sizing: border-box;
}

.page-header {
  width: 100%;
  height: 56px;
  z-index: 1;
}

.page-header-content {
  box-sizing: border-box;
  padding: 10px 24px;
  line-height: 30px;
  position: relative;
  margin: 0 auto;
  /* display: flex; */
  align-items: flex-start;
  justify-content: space-between;
}

.v-wrap {
  height: 1000px;
}
</style>
