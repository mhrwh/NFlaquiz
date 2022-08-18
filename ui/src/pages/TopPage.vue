<template>
  <div id="map">
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
#map{
  position: relative;
};

</style>