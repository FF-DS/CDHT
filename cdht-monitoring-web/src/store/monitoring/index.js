import monitoringAPI from "@/api/monitoring";

const state = {
  monitoringStats: "",
};

const getters = {
  getMonitoringStats: (state) => state.monitoringStats,
};

const actions = {
  fetchMonitoringStats({ commit }, requestBody) {
    monitoringAPI
      .getMonitoringStats(requestBody)
      .then((res) => {
        commit("setMonitoringStats", res.data);
      })
      .catch((err) => {
        console.log(err);
      });
  },
};

const mutations = {
  setMonitoringStats: (state, monitoringStats) =>
    (state.monitoringStats = monitoringStats),
};

export default {
  state,
  getters,
  actions,
  mutations,
};
