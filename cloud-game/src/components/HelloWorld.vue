<template>
  <p>Recommended IDE setup:</p>
</template>

<script lang="ts">
import { Vue, Options } from 'vue-class-component'

@Options({
  // Define component options
  watch: {
    count: (value) => {
      console.log(value)
    },
  },
})
export default class HelloWorld extends Vue {
  count = 0

  increment() {
    this.count++
  }

  created() {
    this.increment()
  }
}
</script>
