<template>
  <div id="app">

    <navbar></navbar>
    <b-container>
      <router-view />
    </b-container>

  </div>
</template>
<script>
import storageService from './service/storageService';
import Navbar from './views/layout/Navbar.vue';

export default {
  components: {
    Navbar,
  },
  mounted () {
    console.log(storageService.USER_TOKEN)
    console.log('toke:', storageService.get(storageService.USER_TOKEN))
  }
}
</script>
<style lang="scss">
</style>
