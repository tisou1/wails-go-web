<script setup>
  //使用内置的api修改窗口标题
  import { 
    WindowSetTitle,
    WindowUnfullscreen
   } from '../../wailsjs/runtime'

  import { WindowFull } from '../../wailsjs/go/main/App'


  function handleClick( ) {
    WindowSetTitle("测试标题")
  }

  const cancelFullscreen = () => {
    WindowUnfullscreen()
  }

  const windowFullscreen = () => {
    WindowFull()
  }
</script>

<template>
  <section class="bg-white">
    <button @click="handleClick" class="btn">修改</button>
    <button @click="windowFullscreen" class="btn">全屏</button>
    <button @click="cancelFullscreen" mx-2 btn>取消全屏</button>
  </section>
</template>

<style>

</style>