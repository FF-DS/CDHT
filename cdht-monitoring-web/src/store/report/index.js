import reportAPI from "@/api/report";

const state = {
  normalReportPackets: [],
  testReportpackets: [],
};

const getters = {
  getNormalReportPackets: (state) => state.normalReportPackets,
  getTestReportPackets: (state) => state.testReportPackets,
};

const actions = {
  fetchReports({ commit }) {
    reportAPI
      .getReports()
      .then((res) => {
        console.log("the request to the report end point is successfull");
        console.log(res);
        commit("setReport", res.data);
      })
      .catch((err) => {
        console.log("something is wrong");
        console.log(err);
      });
  },

  fetchNormalReports({ commit }) {
    reportAPI
      .getNormalReports()
      .then((res) => {
        /* 
        perform the appropriate action with the result based on the status and the data returned
        */
        res, commit;
        commit("setNormalReportPackets", res.data);
      })
      .catch((err) => {
        /* 
        perform appropriate action based on the nature and the status of the error
        */
        console.log(err);
      });
  },
  fetchTestReports({ commit }) {
    reportAPI
      .getTestReports()
      .then((res) => {
        /* 
        perform the appropriate action with the result based on the status and the data returned
        */
        commit("setTestReportPackets", res.data);
      })
      .catch((err) => {
        /* 
        perform appropriate action based on the nature and the status of the error
        */
        console.log(err);
      });
  },
  clearNormalReportCollection({ commit }) {
    reportAPI
      .clearNormalReportCollection()
      .then((res) => {
        /* 
        perform the appropriate action with the result based on the status and the data returned
        */
        res, commit;
      })
      .catch((err) => {
        /* 
        perform appropriate action based on the nature and the status of the error
        */
        console.log(err);
      });
  },
  clearTestReportCollection({ commit }) {
    reportAPI
      .clearTestReportCollection()
      .then((res) => {
        /* 
        perform the appropriate action with the result based on the status and the data returned
        */
        res, commit;
      })
      .catch((err) => {
        /* 
        perform appropriate action based on the nature and the status of the error
        */
        console.log(err);
      });
  },
};

const mutations = {
  setNormalReportPackets: (state, normalReportPackets) =>
    (state.normalReportPackets = normalReportPackets),
  seTestReportPackets: (state, testReportPackets) =>
    (state.testReportPackets = testReportPackets),
};

export default {
  state,
  getters,
  actions,
  mutations,
};
