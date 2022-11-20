<template>
  
    <WorldMap v-if="mapMode==='correctAnswersRate'" v-bind:mapMode="mapMode"></WorldMap>
    <WorldMap v-if="mapMode==='bookMark'" v-bind:mapMode="mapMode"></WorldMap>
    <QuizFilter></QuizFilter>
    <div class="btn-toolbar" role="toolbar">
      <div class="btn-group" role="group">
        <button class="btn btn-light btn-circle" style="border-radius: 50%" v-if="!auth">
          <i class="bi bi-person-fill"></i>
          <div class="btn-text">log in</div>
        </button>
        <button class="btn btn-light btn-circle" style="border-radius: 50%" v-if="auth">
          <i class="bi bi-person-check-fill"></i>
          <div class="btn-text">log out</div>
        </button>

        <!-- <LoginDialog v-if="!auth"></LoginDialog>
        <UserInfoDialog v-if="auth"></UserInfoDialog> -->
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
        <!-- <button type="button" class="bi bi-flag-fill" data-toggle="modal" data-target="#quizFilterModal">
        </button> -->
      </div>
  </div>
</template>

<script>
import WorldMap from '@/components/worldMap.vue'
import QuizFilter from "@/components/quizFilter.vue";
import { computed } from 'vue'
import { useStore } from 'vuex'
import { ref } from 'vue'


export default {
  name: 'TopPage',
  components: {
    WorldMap,
    QuizFilter,
},
  setup() {
    const store = useStore();
    const auth = computed(() => store.state.auth);

    const mapMode = ref("correctAnswersRate")
    return {
      mapMode,
      auth,
    };
  }
};
</script>

<style>

.btn-toolbar{
  position: absolute;
  position: fixed;
  top: 0;
  left: 0;
  z-index: 1;
  margin-top: 2.75rem;
  margin-left: 2.75rem;
}
.btn-circle {
  width: 2.75rem;
  height: 2.75rem;
  text-align: center;
  padding: 0;
  border-radius: 50%;
  margin-right: 4.75rem;
}
.btn-circle i {
  position: relative;
  color: #FFCF32;
  font-size: 1.5rem;
  display: flex;
  justify-content: center;
  top: 0.125rem;
}
.btn-text{
  font-size: 1rem;
  width: calc(100% + 1.75rem);
  margin: 0 -0.875rem;
  margin-top: 14px;
  color: #fff;
  border-bottom: solid 1px #FFCF32;
}
#login{
  position: absolute;
  top: 0px;

};

</style>