import commandDispatcherAPI from "@/api/command-dispatcher";

const state = {
  pendingCommands: [],
  commandResults: [],
};

const getters = {
  getPendingCommands: (state) => state.pendingCommands,
  getCommandResults: (state) => state.commandResults,
};

const actions = {
  fetchPendingCommands({ commit }) {
    commandDispatcherAPI
      .getPendingCommands()
      .then((res) => {
        commit("setPendingCommands", res.data);
      })
      .catch((err) => {
        console.log(err);
      });
  },

  addPendingCommand({ commit }, command) {
    commandDispatcherAPI
      .addPendingCommand(command)
      .then((res) => {
        res, commit;
        /* 
          perform the appropriate action based on the status of the response
          */
      })
      .catch((err) => {
        console.log(err);
        /* 
        perform the appropriate action based on the type of error 
      */
      });
  },
  getCommandResult({ commit }, commandId) {
    commandDispatcherAPI
      .getCommandResult(commandId)
      .then((res) => {
        /* 
          perform the appropriate action based on the status of the response
          */
        commit("setCommandResults", res.data);
      })
      .catch((err) => {
        console.log(err);
        /* 
        perform the appropriate action based on the type of error 
      */
      });
  },

  clearCommandCollection({ commit }) {
    commandDispatcherAPI
      .clearCommandCollection()
      .then((res) => {
        res, commit;
        /* 
          perform the appropriate action based on the status of the response
          */
      })
      .catch((err) => {
        console.log(err);
        /* 
        perform the appropriate action based on the type of error 
      */
      });
  },
};

const mutations = {
  setPendingCommands: (state, pendingCommands) =>
    (state.pendingCommands = pendingCommands),
  setCommandResults: (state, commandResults) =>
    (state.commandResults = commandResults),
};

export default {
  state,
  getters,
  actions,
  mutations,
};
