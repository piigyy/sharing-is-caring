<template>
  <div>
		<div class="col">
			<div class="card card-prod">
				<img :src="image" class="card-img-top p-3">
				<div class="card-body">
					<h5 class="card-title">{{ title }}</h5>
					<table class="table">
						<thead>
							<tr>
								<th scope="col">#</th>
								<th scope="col">Duration</th>
								<th scope="col">Price</th>
								<th scope="col">Action</th>
							</tr>
						</thead>
						<tbody>
							<tr v-for="(price, i) in prices" :key="i">
								<th scope="row">{{ i+1 }}</th>
								<td>{{ price.duration }}</td>
								<td>{{ price.price }}</td>
								<td><button type="button" class="ml-2 btn btn-outline-primary" @click="placeOrder(productID, price)">Buy</button></td>
							</tr>
						</tbody>
					</table>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
export default {
	props: {
		productID: String,
		image: String,
		title: String,
		prices: Array,
		seller: Object,
	},

	methods: {
		placeOrder(productID, price) {
			alert(`you buy ${productID} for ${price.duration} (${price.price})`)
		}
	}
}
</script>

<style scoped>

</style>