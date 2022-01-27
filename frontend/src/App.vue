<template>
  <div id="app">
    <Navbar/>
    <router-view/>
  </div>
</template>

<script>
import Navbar from '@/components/NavBar.vue';

export default {
  name: 'Main',
  components: {
    Navbar,
  }
}
</script>

<style>

</style>
