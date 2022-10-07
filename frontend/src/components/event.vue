<script setup>
  import { onMounted }  from 'vue'
  import { EventsOn , EventsOff, EventsEmit} from '~/wailsjs/runtime'
  import { EventsEmit2 } from '~/wailsjs/go/main/App'

  onMounted(() => {
    console.log('组件挂载')
    EventsOn('click', (args) => {
      console.log('触发事件', args)
    })
  })

  const handleClick = () => {
    EventsEmit('click', {name: 'tisou1', age: 22})
  }

  const handleClick2 = () => {
    EventsEmit2('click')
  }
</script>

<template>
  <section class="bg-white">
    <button @click="handleClick" class="btn">js触发事件</button>
    <button @click="handleClick2" class="btn">go触发事件</button>
  </section>
</template>

<style>

</style>