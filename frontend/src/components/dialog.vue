<script setup>
  import { OpenDialog } from '~/wailsjs/go/main/App'

  const handleClick2 = () => {
    let dialog = {
      title: '我是一个对话框',
      message: '对话框内容体'
    }
    OpenDialog(dialog)
  }
</script>

<template>
  <section class="bg-white">
    <button @click="handleClick2" class="btn">打开对话框</button>
  </section>
</template>

<style>

</style>