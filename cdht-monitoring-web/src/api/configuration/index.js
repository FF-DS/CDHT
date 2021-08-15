import API from "@/services/api";

const rootPath = "/configuration";

const getAllConfigurationProfiles = () => {
  return API.get(`${rootPath}/`);
};

const getCurrentConfigurationProfile = () => {
  return API.get(`${rootPath}/current`);
};

const setCurrentConfigurationProfile = (configurationProfileId) => {
  return API.post(`${rootPath}/current`, { configurationProfileId });
};

const addConfigurationProfile = (configurationProfile) => {
  return API.post(`${rootPath}/add-config-profile`, {
    configurationProfile,
  });
};

const setJumpSpaceBalancing = (jumpSpaceValue) => {
  return API.post(`${rootPath}/set-jump-space`, { jumpSpaceValue });
};

const deleteConfigurationProfile = (configurationId) => {
  return API.delete(`${rootPath}/delete-config-profile`, {
    configurationId,
  });
};

const clearConfigurationProfileCollection = () => {
  return API.get(`${rootPath}/clear`);
};

const configurationAPI = {
  getAllConfigurationProfiles,
  getCurrentConfigurationProfile,
  setCurrentConfigurationProfile,
  addConfigurationProfile,
  setJumpSpaceBalancing,
  deleteConfigurationProfile,
  clearConfigurationProfileCollection,
};

export default configurationAPI;
