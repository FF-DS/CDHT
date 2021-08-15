import API from "@/services/api";

const rootPath = "/replication";

const getReplicasForNode = (nodeId) => {
  return API.get(`${rootPath}/`, { nodeId });
};

const addReplicaForNode = (nodeId, replicaObject) => {
  return API.post(`${rootPath}/`, { nodeId, replicaObject });
};

const deleteReplicaFromNode = (nodeId, replicaId) => {
  return API.delete(`${rootPath}/delte-replica`, { nodeId, replicaId });
};

const clearReplicaCollection = () => {
  return API.get(`${rootPath}/clear`);
};

const replicationAPI = {
  getReplicasForNode,
  addReplicaForNode,
  deleteReplicaFromNode,
  clearReplicaCollection,
};

export default replicationAPI;
