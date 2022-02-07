<template>
  <div class="col">
    <div class="card" style="width: 18rem;">
      <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/08/Netflix_2015_logo.svg/2560px-Netflix_2015_logo.svg.png" class="card-img-top p-3 card-bg" alt="Netflix Logo">
      <div class="card-body card-bg">
        <h5 class="card-title" style="color: #FF0000">Netflix</h5>
        <section class="card-text">
          <ul>
            <li>Host*: <span class="price">Rp 188,000/month</span></li>
            <li>Reg User: <span class="price">Rp 46,500/month</span></li>
          </ul>
        </section>
      </div>
      <ul class="list-group list-group-flush">
        <li class="list-group-item option card-bg"><router-link class="nav-link card-bg" to="/host/netflix">Become a host</router-link></li>
        <li class="list-group-item option card-bg"><router-link class="nav-link card-bg" to="/buy/netflix">Order</router-link></li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    image: String,
    title: String,
    prices: Object,
  },
};
</script>

<style>

</style>