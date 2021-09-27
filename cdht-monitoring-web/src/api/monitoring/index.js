import API from "@/services/api";

const getMonitoringStats = (requestBody) => {
  return API.post("/monitoring/stats", requestBody);
};

const monitoringAPI = {
  getMonitoringStats,
};

export default monitoringAPI;
