<template>
 
  <LoginDialog/>
  <!-- <v-btn elevation="5" @click="currentComponent = 'CorrectAnswersRateMap'">正答率</v-btn>
  <v-btn elevation="5" @click="currentComponent = 'BookMarkMap'">ブックマーク</v-btn> -->
  <WorldMap mapMode='correctAnswersRate'></WorldMap>
</template>

<script>
import WorldMap from '@/components/worldMap.vue'

import LoginDialog from '@/components/loginDialog.vue'
//import { ref } from 'vue'

export default {
  name: 'TopPage',
  components: {
    WorldMap,
    LoginDialog,
},
  // setup() {
  //   const currentComponent = ref("CorrectAnswersRateMap")
  //   return {
  //     currentComponent,
  //   };
  // }
};
</script>
