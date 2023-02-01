<template>
  <div class="btn-toolbar">
    <div class="btn-group" v-if="auth">
      <button class="btn btn-light btn-circle" style="border-radius: 50%" @click="mapSwitch('bookMark')"
        v-if="mapMode == 'correctAnswersRate'">
        <i class="bi bi-layout-sidebar-inset" />
        <div class="btn-text">map</div>
      </button>
      <button class="btn btn-light btn-circle" style="border-radius: 50%" @click="mapSwitch('correctAnswersRate')"
        v-if="mapMode == 'bookMark'">
        <i class="bi bi-layout-sidebar-inset-reverse" />
        <div class="btn-text">map</div>
      </button>
    </div>
    <div class="btn-group" v-if="auth" data-toggle="modal" data-target="#editBookmark">
      <button class="btn btn-light btn-circle" style="border-radius: 50%">
        <i class="bi bi-bookmark-heart" />
        <div class="btn-text btn-text-bookmark">bookmark</div>
      </button>
    </div>
    <div class="btn-group">
      <button class="btn btn-light btn-circle" data-toggle="modal" data-target="#quizFilterModal"
        style="border-radius: 50%">
        <i class="bi bi-flag-fill" />
        <div class="btn-text">quiz</div>
      </button>
    </div>
  </div>
</template>

<script>
import { computed } from 'vue';
import { useStore } from 'vuex';

export default {
  name: 'IconMenu',
  setup() {
    const store = useStore();
    const auth = computed(() => store.state.auth);
    const mapMode = computed(() => store.state.mapMode);

    const mapSwitch = (mapMode) => {
      store.dispatch('setMapMode', mapMode);  
    }

    return {
      auth,
      mapMode,
      mapSwitch,
    };
  },
};
</script>

<style>
.btn-toolbar {
  position: absolute;
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
  margin: 14px -13px 0 -13px;
  color: #fff;
  border-bottom: solid 1px #FFCF32;
}

.btn-text-logout {
  width: calc(100% + 32px);
  margin: 14px -16px 0 -16px;
}

.btn-text-bookmark {
  width: calc(100% + 54px);
  margin: 14px -27px 0 -27px;
}

</style>