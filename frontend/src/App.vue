<template>
  <div id="app">
    <Navbar class="navbar"/>
    <div id="main">
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
    color: white;
  }

  section {
    color: white;
  }

  /* .navbar {
    position: fixed;
    top: 0;
    width: 100%;
  }

  .navbar + #main {
    padding-top: 60px;
  }

  #main {
    padding: 16px;
  } */
</style>
