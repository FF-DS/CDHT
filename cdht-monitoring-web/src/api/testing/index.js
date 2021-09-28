import API from "@/services/api";

const postPingRequest = (requestBody) => {
  return API.post("/test/ping", requestBody);
};

const postDNSRequest = (requestBody) => {
  return API.post("/test/dns-lookup", requestBody);
};

const postHopCountRequest = (requestBody) => {
  return API.post("/test/hop-count", requestBody);
};

const getCurrentActiveOperationReport = (requestBody) => {
  return API.post("/test/filter-results", requestBody);
};

const testingAPI = {
  postPingRequest,
  postDNSRequest,
  postHopCountRequest,
  getCurrentActiveOperationReport,
};

export default testingAPI;
