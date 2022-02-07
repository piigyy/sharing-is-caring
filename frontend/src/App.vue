<template>
  <div id="app">
    <Navbar/>
    <div>
      <router-view/>
    </div>
  </div>
</template>

<script>
import Navbar from '@/components/NavBar.vue';

export default {
  name: 'Main',
  components: {
    Navbar,
  },
  created() {
    this.$store.dispatch('checkLogin');
  },
}
</script>

<style>
  h1 {
    font-family: 'Dongle', sans-serif;
    font-size: 4em;
    color: cadetblue;
  }

  section {
    color: #FFDD93;
  }
</style>
