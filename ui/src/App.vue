<template>
  <button @click="currentComponent = 'CorrectAnswersRateMap'">正答率</button>
  <button @click="currentComponent = 'BookMarkMap'">ブックマーク</button>
  <CorrectAnswersRateMap v-if="currentComponent === 'CorrectAnswersRateMap'"></CorrectAnswersRateMap>
  <BookMarkMap v-if="currentComponent === 'BookMarkMap'"></BookMarkMap>
</template>

<script>
import CorrectAnswersRateMap from './components/correctAnswersRateMap.vue'
import BookMarkMap from './components/bookMarkMap.vue'
export default {
  name: 'App',
  components: {
    CorrectAnswersRateMap,
    BookMarkMap
  },
  data() {
    return {
      currentComponent: "CorrectAnswersRateMap",
    };
  }
}
</script>
