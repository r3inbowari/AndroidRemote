<template>
  <div class="mini-search">
    <form class="nav-search-form">
      <input
        type="text"
        autocomplete="off"
        accesskey="s"
        x-webkit-speech=""
        x-webkit-grammar="builtin:translate"
        placeholder="刻骨铭心初恋情侣套餐132"
        class="nav-search-keyword"
      />
      <div class="nav-search-btn">
        <button
          type="button"
          class="nav-search-submit iconfont icon-electronics"
        ></button>
      </div>
    </form>
  </div>
</template>

<script lang="ts">
import { Vue, Options } from 'vue-class-component'
import { HELLO } from '../../api/nav'

// Component definition
@Options({
  // Define component options
  // watch: {
  //   count: (value) => {
  //     console.log(value)
  //   },
  // },
})
export default class MiniSearach extends Vue {
  // The behavior in class is the same as the current

  created() {
    console.log('asd')
    HELLO(1).then((res) => {
      console.log(res)
    })
  }
}
</script>

<style scoped>
.mini-search {
  position: relative;
}

.nav-search-form {
  border: 1px solid #e7e7e7;
  display: block;
  padding: 0 38px 0 16px;
  border-radius: 2px;
  background: #151515 !important;
  border-color: #151515 !important;
}

.nav-search-keyword {
  overflow: hidden;
  width: 100%;
  height: 34px;
  border: none;
  background-color: transparent;
  box-shadow: none;
  color: #999;
  font-size: 14px;
  line-height: 34px;
  transition: all 0.2s;
}

.nav-search-btn {
  position: absolute;
  top: 0px;
  right: 0;
  margin: 1px;
  padding: 0;
  width: 56px;
  height: 34px;
  border: none;
  border-radius: 2px;
  background: rgb(43, 43, 43);
  line-height: 26px;
  cursor: pointer;
}

.nav-search-submit {
  position: absolute;
  top: 8px;
  right: 16px;
  margin: 0;
  padding: 0;
  border: none;
  background: none;
  color: #505050;
  font-size: 16px;
  line-height: 20px;
  cursor: pointer;
  transition: all 0.2s;
}

.nav-search-submit::before {
  content: '\e8aa';
  font-size: 22px;
}

.nav-search-submit:hover {
  color: rgb(3, 152, 200);
}
</style>
