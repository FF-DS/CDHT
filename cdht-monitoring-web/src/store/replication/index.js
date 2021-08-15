import replicationAPI from "@/api/replication";

const state = {
  replicasForCurrentNode: [],
};

const getters = {
  getReplicasForCurrentNode: (state) => state.replicasForCurrentNode,
};

const actions = {
  fetchReplicasForCurrentNode({ commit }, nodeId) {
    replicationAPI
      .getReplicasForNode(nodeId)
      .then((res) => {
        /* 
        perform the appropriate action with the result based on the status and the data returned
        */
        commit("setReplicasForCurrentNode", res.data);
      })
      .catch((err) => {
        /* 
        perform appropriate action based on the nature and the status of the error
        */
        console.log(err);
      });
  },
  addReplicasForCurrentNode({ commit }, nodeId, replicaObject) {
    replicationAPI
      .addReplicaForNode(nodeId, replicaObject)
      .then((res) => {
        res, commit;
        /* 
        perform the appropriate action with the result based on the status and the data returned
        */
      })
      .catch((err) => {
        /* 
        perform appropriate action based on the nature and the status of the error
        */
        console.log(err);
      });
  },
  deleteReplicasFromCurrentNode({ commit }, nodeId, replicaId) {
    replicationAPI
      .deletereplicaFromNode(nodeId, replicaId)
      .then((res) => {
        res, commit;
        /* 
        perform the appropriate action with the result based on the status and the data returned
        */
      })
      .catch((err) => {
        /* 
        perform appropriate action based on the nature and the status of the error
        */
        console.log(err);
      });
  },
  clearReplicaCollection({ commit }) {
    replicationAPI
      .clearReplicaCollection()
      .then((res) => {
        res, commit;
        /* 
        perform the appropriate action with the result based on the status and the data returned
        */
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
  setReplicasForCurrentNode: (state, replicasForCurrentNode) =>
    (state.replicasForCurrentNode = replicasForCurrentNode),
};

export default {
  state,
  getters,
  actions,
  mutations,
};
