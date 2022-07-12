<template>
<v-container>
    <v-card class="md:mx-4 my-4 px-1 py-20">
      <h1 class="text-2xl md:text-4xl text-center">オープンデータカタログ検索</h1>
    </v-card>

  </v-container>
  <button @click="currentComponent = 'CorrectAnswersRateMap'">正答率</button>
  <button @click="currentComponent = 'BookMarkMap'">ブックマーク</button>
  <CorrectAnswersRateMap v-if="currentComponent === 'CorrectAnswersRateMap'"></CorrectAnswersRateMap>
  <BookMarkMap v-if="currentComponent === 'BookMarkMap'"></BookMarkMap>
</template>

<script>
import CorrectAnswersRateMap from '@/components/correctAnswersRateMap.vue'
import BookMarkMap from '@/components/bookMarkMap.vue'
export default {
  name: 'TopPage',
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