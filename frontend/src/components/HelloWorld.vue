<script setup>
import { reactive } from 'vue'
import { GetBreadList, GetImageUrlByBreed, GetRandomInageUrl } from '../../wailsjs/go/main/App'

const data = reactive({
  randomImageUrl: "",
  breeds: [],
  photos: [],
  selectedBreed: '',
  showRandomPhoto: false,
  showBreedPhotos: false
})

function init() {
  getBreadList()
}
init()


function getRandomInageUrl() {
  data.showBreedPhotos = false
  data.showRandomPhoto = false
  GetRandomInageUrl().then(result => {
    data.randomImageUrl = result
  })
  data.showRandomPhoto = true
}

function getBreadList() {
  GetBreadList().then(result => {
    data.breeds = result
  })
}

function getImageUrlByBreed() {
  init()
  data.showBreedPhotos = false
  data.showRandomPhoto = false
  console.log(data.selectedBreed)
  GetImageUrlByBreed(data.selectedBreed).then(result => {
    data.photos = result
    console.log(result)
  })
  data.showBreedPhotos = true
}
</script>

<template>
  <main>
    <h3>Dog API</h3>
    <div>
      <button class="btn" @click="getRandomInageUrl">获取随机狗子</button>
      下拉选择一个狗子
      <select v-model="data.selectedBreed">
        <option v-for="breed in data.breeds">
          {{breed}}
        </option>
      </select>
      <button class="btn" v-on:click="getImageUrlByBreed">
        获取指定品种狗子
      </button>
      <br/>
      <img v-if="data.showRandomPhoto" v-bind:src="data.randomImageUrl" alt="No dog found" id="breed-photos">
      <div v-if="data.showBreedPhotos" class="img-box">
        <img v-for="photo in data.photos" :src="photo" alt="No dog found"/>
      </div>
    </div>
  </main>
</template>

<style scoped >
  #random-photo {
  width: 600px;
  height: auto;
}

#breed-photos {
  width: 300px;
  height: auto;
}
.btn:hover{
  cursor: pointer;
}

.btn:focus {
  border-width: 3px;
}
.img-box{
  display: grid;
  gap: 0.5rem;
  grid-template:auto /repeat(5, 20%);
}
 
.img-box img{
  width: 100%;
  height: 100%;
  object-fit: contain;
}

</style>