<template>
  <WorldMap v-if="mapMode==='correctAnswersRate'" v-bind:mapMode="mapMode"></WorldMap>
  <WorldMap v-if="mapMode==='bookMark'" v-bind:mapMode="mapMode"></WorldMap>
  <div class="btn-toolbar" role="toolbar">
      <div class="btn-group" role="group" v-if="auth">
        <button class="btn btn-light btn-circle" style="border-radius: 50%" @click="mapMode = 'bookMark'" v-if="mapMode=='correctAnswersRate'">
          <i class="bi bi-layout-sidebar-inset"></i>
          <div class="btn-text">map</div>
        </button>
        <button class="btn btn-light btn-circle" style="border-radius: 50%" @click="mapMode = 'correctAnswersRate'" v-if="mapMode=='bookMark'">
          <i class="bi bi-layout-sidebar-inset-reverse"></i>
          <div class="btn-text">map</div>
        </button>
      </div>
  </div>
</template>

<script>
import WorldMap from '@/components/worldMap.vue'
import { ref } from 'vue'
import { computed } from 'vue'
import { useStore } from 'vuex'

export default {
  name: 'TopPage',
  components: {
    WorldMap,
},
  setup() {
    const store = useStore();
    const auth = computed(() => store.state.auth);

    const mapMode = ref("correctAnswersRate")
    return {
      auth,
      mapMode,
    };
  }
};
</script>

<style>
.btn-toolbar {
  position: fixed;
  top: 0;
  left: 0;
  z-index: 1;
  margin-top: 45px;
  margin-left: 46px;
}
.btn-circle {
  width: 45px;
  height: 45px;
  text-align: center;
  padding: 0;
  border-radius: 50%;
  margin-right: 76px;
}
.btn-circle i {
  position: relative;
  color: #FFCF32;
  font-size: 26px;
  display: flex;
  justify-content: center;
  top: 2px;
}
.btn-text {
  font-size: 16px;
  width: calc(100% + 26px);
  margin: 0 -13px;
  margin-top: 14px;
  color: #fff;
  border-bottom: solid 1px #FFCF32;
}
</style>