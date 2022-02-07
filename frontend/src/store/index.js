import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    isLogin: false,
    isLoading: true,
    email: "logged user",
  },
  mutations: {
    SET_ISLOGIN(state, payload) {
      state.isLogin = payload
    },
    SET_EMAIL(state, payload) {
      state.email = payload
    },
  },
  actions: {
    checkLogin({ commit }) {
      console.log("checking login...")
      if (localStorage.getItem("user")) {
        console.log("login info found!")
        commit("SET_ISLOGIN", true);
        const { email } = JSON.parse(localStorage.getItem("user"))
        commit("SET_EMAIL", email);
      } else {
        console.log("login info not found!")
        localStorage.removeItem("user");
        commit("SET_ISLOGIN", false);
      }
    },
    signOut({dispatch}) {
      localStorage.removeItem("user");
      dispatch('checkLogin');
    }
  },
  modules: {
  }
})
