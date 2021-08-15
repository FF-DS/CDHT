import API from "@/services/api";

const rootPath = "/command-dispatcher";

const getPendingCommands = () => {
  return API.get(`${rootPath}/`);
};

const addPendingCommand = ({ command }) => {
  return API.post(`${rootPath}/`, { command });
};

const getCommandResult = (commandId) => {
  return API.get(`${rootPath}/result`, { commandId });
};

const clearCommandsCollection = () => {
  return API.get(`${rootPath}/clear-commands`);
};

const clearCommandResultsCollection = () => {
  return API.get(`${rootPath}/clear-results`);
};

const commandDispatcherAPI = {
  getPendingCommands,
  addPendingCommand,
  getCommandResult,
  clearCommandsCollection,
  clearCommandResultsCollection,
};

export default commandDispatcherAPI;
