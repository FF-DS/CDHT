import API from "@/services/api";

const rootPath = "/report";

const getAllReports = (requestBody) => {
  return API.post(`${rootPath}/all`, requestBody);
};

const getFilteredReports = (filterRequestBody) => {
  return API.post(`${rootPath}/filtered`, filterRequestBody);
};

const reportAPI = {
  getAllReports,
  getFilteredReports,
};

export default reportAPI;
