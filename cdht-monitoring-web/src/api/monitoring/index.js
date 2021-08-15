import API from "@/services/api";

const getMonitoring = () => {
  return API.get("/monitoring/");
};

const monitoringAPI = {
  getMonitoring,
};

export default monitoringAPI;
