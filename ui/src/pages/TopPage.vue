<template>
  
  <div id="map">
    <WorldMap v-if="mapMode==='correctAnswersRate'" v-bind:mapMode="mapMode"></WorldMap>
    <WorldMap v-if="mapMode==='bookMark'" v-bind:mapMode="mapMode"></WorldMap>
    <div id="login">
      
      <v-btn elevation="5" @click="mapMode = 'correctAnswersRate'">正答率</v-btn>
      <v-btn elevation="5" @click="mapMode = 'bookMark'">ブックマーク</v-btn>
    </div>
  </div>
</template>

<script>
import WorldMap from '@/components/worldMap.vue'
import { ref } from 'vue'

export default {
  name: 'TopPage',
  components: {
    WorldMap,
    
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
#map{
  position: relative;
};
#login{
  position: absolute;
  top: 0px;

};
</style>