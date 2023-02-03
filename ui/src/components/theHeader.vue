<template>
  <nav class="navbar flex-nowrap navbar-light border-bottom header">
    <a class="navbar-brand">
      <img src="@/assets/NFlaquizlogo.png" width="50" height="50" class="d-inline-block align-top" alt="">
    </a>
    <div class="navbar-brand" v-if="$route.name==='TopPage'">
      <button class="btn btn-header" data-toggle="modal" data-target="#loginModal" data-backdrop="false" v-if="!auth">
          <i class="bi bi-person-fill" />
      </button>
      <button class="btn btn-header" @click="logout" v-else-if="auth">
        <i class="bi bi-person-check-fill" />
      </button>
    </div>
  </nav>
</template>

<script>
import { computed } from 'vue';
import { useStore } from 'vuex';
import axios from 'axios';

export default {
  name: 'TheHeader',
  setup() {
    const store = useStore();
    const auth = computed(() => store.state.auth);

    const logout = async () => {
      axios.get('http://localhost:8888/logout')
        .then(() => {
          store.dispatch("setAuth", false);
          store.dispatch('setMapMode', 'correctAnswersRate');
          location.reload();
        })
    };

    return {
      auth,
      logout,
    };
  },
}
</script>

<style>
.header {
  height: 50px;
}
.btn-header i {
  color: #FFCF32;
  font-size: 26px;
}
</style>