<template>
  <div>
    <WorldMap v-if="mapMode === 'correctAnswersRate'" :mapMode="mapMode">
      <IconMenu />
    </WorldMap>
    <WorldMap v-else-if="mapMode === 'bookMark'" :mapMode="mapMode">
      <IconMenu />
    </WorldMap>
    <QuizFilter />
    <AccountModal />
    <EditBookmark />
  </div>
</template>

<script>
import AccountModal from '@/components/accountModal.vue';
import WorldMap from '@/components/worldMap.vue';
import QuizFilter from "@/components/quizFilter.vue";
import EditBookmark from "@/components/EditBookmark.vue";
import IconMenu from '@/components/iconMenu.vue';
import { computed } from 'vue';
import { useStore } from 'vuex';

export default {
  name: 'TopPage',
  components: {
    WorldMap,
    AccountModal,
    QuizFilter,
    EditBookmark,
    IconMenu
},
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