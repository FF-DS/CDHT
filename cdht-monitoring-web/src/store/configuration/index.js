import configurationAPI from "@/api/configuration";

const state = {
  allConfuigurationProfiles: [],
  currentConfigurationProfile: "",
};

const getters = {
  getAllConfigurationProfiles: (state) => state.allConfuigurationProfiles,
  getCurrentConfigurationProfiles: (state) => state.currentConfigurationProfile,
};

const actions = {
  getAllConfigurationProfiles({ commit }) {
    configurationAPI
      .getAllConfigurationProfiles()
      .then((res) => {
        /* 
        perform the appropriate action with the result based on the status and the data returned
        */
        commit("setAllConfigurationProfiles", res.data);
      })
      .catch((err) => {
        /* 
        perform appropriate action based on the nature and the status of the error
        */
        console.log(err);
      });
  },

  getCurrentConfigurationProfile({ commit }) {
    configurationAPI
      .getCurrentConfigurationProfile()
      .then((res) => {
        /* 
        perform the appropriate action with the result based on the status and the data returned
        */
        commit("setCurrentConfigurationProfile", res.data);
      })
      .catch((err) => {
        /* 
        perform appropriate action based on the nature and the status of the error
        */
        console.log(err);
      });
  },

  changeCurrentConfigurationProfile({ commit }, configurationProfileId) {
    configurationAPI
      .setCurrentConfigurationProfile(configurationProfileId)
      .then((res) => {
        /* 
        perform the appropriate action with the result based on the status and the data returned
        */
        commit("setCurrentConfigurationProfile", res.data);
      })
      .catch((err) => {
        /* 
        perform appropriate action based on the nature and the status of the error
        */
        console.log(err);
      });
  },

  addConfigurationProfile({ commit }, configurationProfile) {
    configurationAPI
      .addConfigurationProfile(configurationProfile)
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

  setJumpSpaceBalancing({ commit }, jumpSpaceValue) {
    configurationAPI
      .setJumpSpaceBalancing(jumpSpaceValue)
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

  deleteConfigurationProfile({ commit }, configurationId) {
    configurationAPI
      .deleteConfigurationProfile(configurationId)
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

  clearConfigurationProfilesCollection({ commit }) {
    configurationAPI
      .clearConfigurationProfileCollection()
      .then((res) => {
        /* 
        perform the appropriate action with the result based on the status and the data returned
        */
        res, commit;
      })
      .catch((err) => {
        console.log(err);
      });
  },
};

const mutations = {
  setAllConfigurationProfiles: (state, allConfuigurationProfiles) =>
    (state.allConfuigurationProfiles = allConfuigurationProfiles),
  setCurrentConfigurationProfile: (state, currentConfigurationProfile) =>
    (state.currentConfigurationProfile = currentConfigurationProfile),
};

export default {
  state,
  getters,
  actions,
  mutations,
};
