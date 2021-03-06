import Vue from 'vue'
import VueRouter from 'vue-router'
// import Home from '../views/Home.vue'
import Product from '../views/Products.vue';

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    // component: Home
    component: Product,
  },
  {
    path: '/terms-conditions',
    name: 'Terms And Condition',
    // route level code-splitting
    // this generates a separate chunk (terms-and-condition.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "terms-and-condition" */ '../views/TermsCondition.vue')
  },
  {
    path: '/products',
    name: 'Products',
    // route level code-splitting
    // this generates a separate chunk (products.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "products" */ '../views/Products.vue')
  },
  {
    path: '/signin',
    name: 'Sign In',
    // route level code-splitting
    // this generates a separate chunk (signin.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "signin" */ '../views/Authentication.vue')
  },
  {
    path: '/me/:id',
    name: 'Profile',
    // route level code-splitting
    // this generates a separate chunk (profile.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "profile" */ '../views/Profile.vue')
  },
  {
    path: '/marketplace',
    name: 'Marketplace',
    // route level code-splitting
    // this generates a separate chunk (marketplace.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "marketplace" */ '../views/Marketplace.vue')
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
