import testingAPI from "@/api/testing";

const state = {
  activeOperationId: "",
  activeOperationResults: [],
};

const getters = {
  getActiveOperationId: (state) => state.activeOperationId,
  getActiveOperationResults: (state) => state.activeOperationResults,
};

const actions = {
  sendPingRequest({ commit }, requestBody) {
    testingAPI
      .postPingRequest(requestBody)
      .then((res) => {
        commit("setActiveOperationId", res.data.operation_id);
      })
      .catch((err) => {
        console.log(err);
      });
  },

  sendDNSRequest({ commit }, requestBody) {
    testingAPI
      .postDNSRequest(requestBody)
      .then((res) => {
        commit("setActiveOperationId", res.data.operation_id);
      })
      .catch((err) => {
        console.log(err);
      });
  },

  sendHopCountRequest({ commit }, requestBody) {
    testingAPI
      .postHopCountRequest(requestBody)
      .then((res) => {
        commit("setActiveOperationId", res.data.operation_id);
      })
      .catch((err) => {
        console.log(err);
      });
  },

  getCurrentActiveOperationReport({ commit }, requestBody) {
    testingAPI
      .getCurrentActiveOperationReport(requestBody)
      .then((res) => {
        commit("setActiveOperationResults", res.data);
      })
      .catch((err) => {
        console.log(err);
      });
  },
};

const mutations = {
  setActiveOperationId: (state, activeOperationId) =>
    (state.activeOperationId = activeOperationId),

  setActiveOperationResults: (state, activeOperationResults) =>
    (state.activeOperationResults = activeOperationResults),
};

export default {
  state,
  getters,
  actions,
  mutations,
};
