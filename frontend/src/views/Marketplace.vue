<template>
  <div class="container">
		<div class="row row-cols-1 row-cols-md-3 g-4">
			<marketplace
				class="p-4"
				v-for="product in products" :key="product.id"
				:productID="product.id.toString()"
				:image="product.image"
				:title="product.name"
				:prices="product.prices"
			/>
		</div>
	</div>
</template>

<script>
import Marketplace from '@/components/Marketplace.vue';

export default {
	components: {
		Marketplace,
	},
	data: function() {
		return {
			products: [
				{
					id: 1,
					name: "Spotify Family",
					image: "https://storage.googleapis.com/pr-newsroom-wp/1/2018/11/Spotify_Logo_CMYK_Green.png",
					prices: [
						{
							duration: "1 months",
							price: "10,000 IDR"
						},
						{
							duration: "3 months",
							price: "25,000 IDR"
						},
					],
				},
				{
					id: 2,
					name: "Spotify Family",
					image: "https://storage.googleapis.com/pr-newsroom-wp/1/2018/11/Spotify_Logo_CMYK_Green.png",
					prices: [
						{
							duration: "1 months",
							price: "15,000 IDR"
						},
						{
							duration: "3 months",
							price: "20,000 IDR"
						},
					],
				},
				{
					id: 3,
					name: "Spotify Family",
					image: "https://storage.googleapis.com/pr-newsroom-wp/1/2018/11/Spotify_Logo_CMYK_Green.png",
					prices: [
						{
							duration: "1 months",
							price: "15,000 IDR"
						},
						{
							duration: "3 months",
							price: "20,000 IDR"
						},
					],
				},
				{
					id: 4,
					name: "Spotify Family",
					image: "https://storage.googleapis.com/pr-newsroom-wp/1/2018/11/Spotify_Logo_CMYK_Green.png",
					prices: [
						{
							duration: "1 months",
							price: "15,000 IDR"
						},
						{
							duration: "3 months",
							price: "20,000 IDR"
						},
					],
				},
				{
					id: 5,
					name: "Spotify Family",
					image: "https://storage.googleapis.com/pr-newsroom-wp/1/2018/11/Spotify_Logo_CMYK_Green.png",
					prices: [
						{
							duration: "1 months",
							price: "15,000 IDR"
						},
						{
							duration: "3 months",
							price: "20,000 IDR"
						},
					],
				},
				{
					id: 6,
					name: "Netlfix",
					image: "https://upload.wikimedia.org/wikipedia/commons/thumb/0/08/Netflix_2015_logo.svg/2560px-Netflix_2015_logo.svg.png",
					prices: [
						{
							duration: "1 months",
							price: "15,000 IDR"
						},
						{
							duration: "3 months",
							price: "20,000 IDR"
						},
						{
							duration: "1 year",
							price: "50,000 IDR"
						},
					],
				},
			],
		};
	},
}
</script>

<style>

</style>