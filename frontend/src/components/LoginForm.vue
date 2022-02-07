<template>
  <div class="container">
    <div class="row" id="title">
      <h1>Sign In</h1>
    </div>
    <form class="">
      <div class="row">
        <div class="input-group input-group-lg p-3">
          <span class="input-group-text" id="input-email">Email</span>
          <input
            required
            v-model="email"
            type="email"
            class="form-control"
            aria-label="Input your registered email"
            aria-describedby="inputGroup-sizing-sm"
          >
        </div>
      </div>
      <div class="row">
        <div class="input-group input-group-lg p-3">
          <span class="input-group-text" id="input-password">Password</span>
          <input
            required
            v-model="password"
            type="password"
            class="form-control"
            aria-label="Input your registered email"
            aria-describedby="inputGroup-sizing-sm"
          >
        </div>
      </div>
      <div class="row item-align-center txt-center text-center mt-3">
        <div class="col">
          <button v-if="!isFetch" type="submit" @click="signin" class="btn btn-primary">Sign-In</button>
          <button v-else class="btn btn-primary" type="button" disabled>
            <span class="spinner-border spinner-border-sm" role="status" aria-hidden="true"></span>
            Wait...
          </button>
        </div>
      </div>
      <div class="row item-align-center txt-center text-center mt-2" style="color: white">
        <p><i>Don't have any account yet?</i> <a href="#" @click="signUp">Register Now</a> </p>
      </div>
    </form>
  </div>
</template>

<script>
export default {
  data: function () {
    return {
      email: "",
      password: "",
      isFetch: false,
    };
  },
  methods: {
    signUp() {
      this.$swal.fire({
        title: 'This Website Currently Is Use Only For Internal User',
        icon: 'info',
        html: `
          If you still want to register, you could contact me on WhatsApp by clicking this link: <a href="https://wa.me/6289658876167?text=I%20want%20to%20register%20a%20Sharing-is-Caring%20account" target="_blank">Ask For An Account</a> <br/><br/>
          <i>By registering or/and request for an account you agree to our <a href="/terms-conditions">terms and conditions</a></i>
        `,
        showCloseButton: true,
      })
    },
    async signin() {
      this.isFetch = true;
      if (this.email === "" || this.password.length < 8) {
        this.$swal("Please Input Your Email and Password!")
        this.isFetch = false;
      } else {
        try {
          const { data } = await this.$http.post("api/v1/login", {
            email: this.email,
            password: this.password,
          });
          localStorage.setItem("user", JSON.stringify(data.data));
          setTimeout(() => {
            this.isFetch = false;
          }, 500);
          this.$store.dispatch('checkLogin');
          this.$swal("Sign-In Success!");
          this.$router.push("/products");

        } catch (err) {
          if (err.response) {
            if (err.response.status == 401) {
              this.$swal(
                "Bad Request",
                "Perhaps You Entered Wrong Email or/and Password Combination",
                "error"
              );
            } else {
              this.$swal(
                "Internal Server Error",
                "It's not you, we just messed things up :(",
                "error"
              );
            }
          } else {
            this.$swal(
              "Internal Server Error",
              "It's not you, we just messed things up :(",
              "error"
            );
          }
          setTimeout(() => {
            this.isFetch = false;
          }, 500);
        }
      }
    },
  },
}
</script>

<style scoped>
  #title {
    text-align: center;
  }
</style>