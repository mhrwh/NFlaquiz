<template>
  <TheHeader />
  <router-view />
</template>

<script>
import TheHeader from './components/theHeader.vue';
export default {
  name: 'App',
  components: {
    TheHeader
  },
}
</script>