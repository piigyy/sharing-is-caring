import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    userID: '',
    isLogin: false,
    isLoading: true,
    email: "logged@user.com",
    name: "logged user",
  },
  mutations: {
    SET_ISLOGIN(state, payload) {
      state.isLogin = payload;
    },
    SET_EMAIL(state, payload) {
      state.email = payload;
    },
    SET_NAME(state, payload) {
      state.name = payload;
    },
    SET_USERID(state, payload) {
      state.userID = payload;
    },
  },
  actions: {
    checkLogin({ commit }) {
      const loginModalCloseButton = document.getElementById("close-modal-login");
      if (loginModalCloseButton) {
        loginModalCloseButton.click();
      }

      console.log("checking login...")
      if (localStorage.getItem("user")) {
        console.log("login info found!")
        commit("SET_ISLOGIN", true);
        const { email, name, id } = JSON.parse(localStorage.getItem("user"))
        commit("SET_EMAIL", email);
        commit("SET_NAME", name);
        commit("SET_USERID", id);
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
