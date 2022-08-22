<template>
  <div id="map">
    <div class="btn-group" role="group" style="position: absolute; z-index: 1;">
      <button type="button" class="btn btn-secondary" @click="mapMode = 'correctAnswersRate'">正答率</button>
      <button type="button" class="btn btn-secondary" @click="mapMode = 'bookMark'">ブックマーク</button>
    </div>
    <WorldMap v-if="mapMode==='correctAnswersRate'" v-bind:mapMode="mapMode"></WorldMap>
    <WorldMap v-if="mapMode==='bookMark'" v-bind:mapMode="mapMode"></WorldMap>
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
#map {
  position: relative;
};
</style>