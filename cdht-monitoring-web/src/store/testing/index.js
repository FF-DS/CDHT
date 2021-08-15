import testingAPI from "@/api/testing";

const state = {
  test: "",
};

const getters = {
  getTest: (state) => state.test,
};

const actions = {
  fetchTest({ commit }) {
    testingAPI
      .getTest()
      .then((res) => {
        commit("setTest", res.message);
      })
      .catch((err) => {
        console.log(err);
      });
  },
};

const mutations = {
  setTest: (state, test) => (state.test = test),
};

export default {
  state,
  getters,
  actions,
  mutations,
};
