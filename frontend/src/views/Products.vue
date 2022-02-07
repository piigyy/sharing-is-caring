<template>
  <div class="container">
    <div class="row m-2">
      <div class="col">
        <h1 style="text-align:center;">Products And Pricings</h1>
      </div>
    </div>

    <div class="">
      <div class="">
        <div class="row gx-5">
          <div class="col animate__animated animate__backInLeft">
            <div class="card" style="width: 18rem;">
              <img src="https://storage.googleapis.com/pr-newsroom-wp/1/2018/11/Spotify_Logo_CMYK_Green.png" class="card-img-top p-3 card-bg" alt="Spotify Logo">
              <div class="card-body card-bg">
                <h5 class="card-title product-title">Spotify Family</h5>
                <section class="card-text">
                  <ul slot="color: black;">
                    <li>Host*: <span class="price">Rp 79,000/month</span></li>
                    <li>Reg User: <span class="price">Rp 17,000/month</span></li>
                    <li>Reg User: <span class="price">Rp 57,000/3 months</span></li>
                  </ul>
                </section>
              </div>
              <ul class="list-group list-group-flush">
                <li class="list-group-item option card-bg" @click="clickedOrder('/host/spotify')"><span class="action-order">Become a host</span></li>
                <li class="list-group-item option card-bg" @click="clickedOrder('/order/spotify')" ><span class="action-order">Order</span></li>
              </ul>
            </div>
          </div>
          <div class="col animate__animated animate__backInRight">
            <div class="card" style="width: 18rem;">
              <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/08/Netflix_2015_logo.svg/2560px-Netflix_2015_logo.svg.png" class="card-img-top p-3 card-bg" alt="Netflix Logo">
              <div class="card-body card-bg">
                <h5 class="card-title product-title">Netflix</h5>
                <section class="card-text">
                  <ul>
                    <li>Host*: <span class="price">Rp 188,000/month</span></li>
                    <li>Reg User: <span class="price">Rp 46,500/month</span></li>
                  </ul>
                </section>
              </div>
              <ul class="list-group list-group-flush">
                <li class="list-group-item option card-bg" @click="clickedOrder('/host/netflix')"><span class="action-order">Become a host</span></li>
                <li class="list-group-item option card-bg" @click="clickedOrder('/order/netflix')" ><span class="action-order">Order</span></li>
              </ul>
            </div>
          </div>
          <!-- <div class="col">
            <div class="card" style="width: 18rem;">
              <img src="https://logos-world.net/wp-content/uploads/2020/04/Twitter-Logo-2010-2012.png" class="card-img-top p-3 card-bg" alt="Twitter Logo">
              <div class="card-body card-bg">
                <h5 class="card-title product-title">Twitter Private Account</h5>
                <section class="card-text">
                  <ul>
                    <li><i>based on host offers</i></li>
                  </ul>
                </section>
              </div>
              <ul class="list-group list-group-flush">
                <li class="list-group-item option card-bg"><router-link class="nav-link card-bg" to="/host/twitter"><span class="action-order">Become a host</span></router-link></li>
                <li class="list-group-item option card-bg"><router-link class="nav-link card-bg" to="/buy/twitter"><span class="action-order">Order</span></router-link></li>
              </ul>
            </div>
          </div> -->
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  methods: {
    clickedOrder(url) {
      if (!this.$store.state.isLogin) {
        this.$swal("Please Sign-In Before Ordering")
        this.$router.push("/signin");
      } else {
        this.$router.push(url)
      }
    },
  },
}
</script>

<style scoped>
  .card {
    margin: 0 auto; /* Added */
    float: none; /* Added */
    margin-bottom: 10px; /* Added */
  }

  .price {
    font-weight: bold;
  }

  .option {
    cursor: pointer;
  }
  
  .card-bg {
    background-color: #750550;
  }

  .product-title {
    color: snow;
  }

  .action-order {
    color: wheat;
    font-weight: bold;
  }
</style>