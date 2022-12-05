<template>
  <WorldMap v-if="mapMode==='correctAnswersRate'" v-bind:mapMode="mapMode"></WorldMap>
  <WorldMap v-else-if="mapMode==='bookMark'" v-bind:mapMode="mapMode"></WorldMap>
  <QuizFilter />
  <AccountModal />
  <div class="btn-toolbar" role="toolbar">
      <div class="btn-group" role="group">
        <button class="btn btn-light btn-circle" data-toggle="modal" data-target="#loginModal" data-backdrop="false" style="border-radius: 50%" v-if="!auth">
          <i class="bi bi-person-fill"></i>
          <div class="btn-text">log in</div>
        </button>
        <button class="btn btn-light btn-circle" data-toggle="modal" data-target="#logoutModal" data-backdrop="false" style="border-radius: 50%" v-if="auth">
          <i class="bi bi-person-check-fill"></i>
          <div class="btn-text logout">log out</div>
        </button>
      </div>
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
      <div class="btn-group" role="group">
        <button class="btn btn-light btn-circle" data-toggle="modal" data-target="#quizFilterModal" style="border-radius: 50%">
          <i class="bi bi-flag-fill"></i>
          <div class="btn-text">quiz</div>
        </button>
      </div>
  </div>
</template>

<script>
import AccountModal from '@/components/accountModal.vue';
import WorldMap from '@/components/worldMap.vue';
import QuizFilter from "@/components/quizFilter.vue";
import { ref } from 'vue'
import { computed } from 'vue'
import { useStore } from 'vuex'

export default {
  name: 'TopPage',
  components: {
    WorldMap,
    AccountModal,
    QuizFilter,
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
.logout{
  width: calc(100% + 32px);
  margin: 0 -16px;
  margin-top: 14px;
}
</style>