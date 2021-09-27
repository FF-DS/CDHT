import API from "@/services/api";

const getMonitoringStats = (requestBody) => {
  return API.get("/monitoring/stats", requestBody);
};

const monitoringAPI = {
  getMonitoringStats,
};

export default monitoringAPI;
