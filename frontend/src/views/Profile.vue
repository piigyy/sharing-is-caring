<template>
	<div class="container profile">
		<div class="row item-align-center mt-2">
			<div class="card text-center profile-main">
				<div class="card-body">
					<div class="text-center">
						<img :src="'https://ui-avatars.com/api/?name=' + $store.state.name" class="rounded-circle border border-primary profile-image">
					</div>
					<div class="text-center name">
						<h1 style="color:black">{{ $store.state.name }} </h1>
					</div>
					<div class="container">
						<div class="row">
							<div class="col-6">
								<button type="button" class="btn btn-outline-primary">My Order</button>
							</div>
							<div class="col-6">
								<button type="button" class="btn btn-outline-primary">My Products</button>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
export default {
	mounted() {
		// This should be on router guard,
		// but still trying to findout how.
		if (!this.$store.state.isLogin) {
			this.$router.push('/signin')
			this.$swal.fire("Please Sign-In First")
		}
	},
}
</script>

<style scoped>
	.profile {
		color: white;
		padding: 2em;
	}

	@media only screen and (min-width: 768px) {
    .profile-main {
      /* max-width: 50%; */
    }
  }

	.profile-image {
		height: 6em;
	}

</style>