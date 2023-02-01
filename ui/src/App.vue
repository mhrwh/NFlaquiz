<template>
  <PageHeader />
  <router-view />
</template>

<script>
import PageHeader from './components/pageHeader.vue';
export default {
  name: 'App',
  components: {
    PageHeader
  },
}
</script>