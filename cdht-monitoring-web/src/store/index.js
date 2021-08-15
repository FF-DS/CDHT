import Vue from "vue";
import Vuex from "vuex";
Vue.use(Vuex);

import report from "./report";
import configuration from "./configuration";
import monitoring from "./monitoring";
import testing from "./testing";
import commandDispatcher from "./command-dispatcher";
import replication from "./replication";

export default new Vuex.Store({
  state: {},
  mutations: {},
  actions: {},
  modules: {
    report,
    configuration,
    monitoring,
    testing,
    commandDispatcher,
    replication,
  },
});
