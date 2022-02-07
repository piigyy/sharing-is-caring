<template>
  <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
		<div class="container-fluid">
			<router-link class="navbar-brand title" to="/">Sharing is Caring</router-link>
			<button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNavDropdown" aria-controls="navbarNavDropdown" aria-expanded="false" aria-label="Toggle navigation">
				<span class="navbar-toggler-icon"></span>
			</button>
			<div class="collapse navbar-collapse" id="navbarNavDropdown">
				<ul class="navbar-nav">
					<li class="nav-item">
						<router-link class="nav-link" to="/products">Product and Pricing</router-link>
					</li>
					<li class="nav-item">
						<router-link class="nav-link" to="/terms-conditions">Terms and Condition</router-link>
					</li>
					<li class="nav-item" v-if="!$store.state.isLogin">
						<router-link class="nav-link btn btn-secondary" to="/signin">Sign-In</router-link>
					</li>
					<li class="nav-item" v-if="$store.state.isLogin">
						<span class="nav-link">{{ $store.state.email }}</span>
					</li>
					<li class="nav-item" v-if="$store.state.isLogin">
						<span @click="doLogout" class="nav-link btn btn-secondary">Sign-Out</span>
					</li>
				</ul>
			</div>
		</div>
	</nav>
</template>

<script>
export default {
	methods: {
		doLogout() {
			this.$store.dispatch('signOut');
			this.$swal("Sign Out success!");
			this.$router.push("/signin");
		},
	},
}
</script>

<style scoped>
	.title {
		font-family: 'Neonderthaw', cursive;
		font-weight: bold;
		font-size: 1.5em;
	}
</style>