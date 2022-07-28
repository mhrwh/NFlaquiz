<template>
  
  <v-btn elevation="5" @click="mapMode = 'correctAnswersRate'">正答率</v-btn>
  <v-btn elevation="5" @click="mapMode = 'bookMark'">ブックマーク</v-btn>
  <WorldMap v-if="mapMode==='correctAnswersRate'" v-bind:mapMode="mapMode"></WorldMap>
  <WorldMap v-if="mapMode==='bookMark'" v-bind:mapMode="mapMode"></WorldMap>
  <LoginDialog/>
    
</template>

<script>
import WorldMap from '@/components/worldMap.vue'

import LoginDialog from '@/components/loginDialog.vue'
import { ref } from 'vue'

export default {
  name: 'TopPage',
  components: {
    WorldMap,
    LoginDialog,
},
  setup() {
    const mapMode = ref("correctAnswersRate")
    return {
      mapMode,
    };
  }
};
</script>

<style>

</style>