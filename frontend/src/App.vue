<template>
  <div id="app">
    <Navbar class="navbar"/>
    <div id="main">
      <router-view/>
    </div>
    <a href="#" class="float" data-bs-toggle="modal" data-bs-target="#exampleModal">
      <i class="fa fa-plus my-float"></i>
    </a>
    <div class="modal fade" id="exampleModal" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="exampleModalLabel">Sell Premium Account</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body">
            ...
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
          </div>
        </div>
      </div>
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
    console.log(`app runing in ${process.env.VUE_APP_ENV} mode`);
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
  .float{
    position:fixed;
    width:60px;
    height:60px;
    bottom:40px;
    right:40px;
    background-color:#0C9;
    color:#FFF;
    border-radius:50px;
    text-align:center;
    box-shadow: 2px 2px 3px #999;
  }

  .my-float{
    margin-top:22px;
  }
</style>
