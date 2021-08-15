import monitoringAPI from "@/api/monitoring";

const state = {
  monitor: "",
};

const getters = {
  getMonitor: (state) => state.monitor,
};

const actions = {
  fetchMonitor({ commit }) {
    monitoringAPI
      .getMonitoring()
      .then((res) => {
        commit("setMonitor", res.message);
      })
      .catch((err) => {
        console.log(err);
      });
  },
};

const mutations = {
  setMonitor: (state, monitor) => (state.monitor = monitor),
};

export default {
  state,
  getters,
  actions,
  mutations,
};
