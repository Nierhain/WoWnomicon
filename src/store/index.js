import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    charInfo: {
      name: String,
      realm: String,
      level: Number,
      class: String,
      race: String,
      itemLevel: Number
    }
  },
  mutations: {},
  actions: {},
  modules: {}
});
