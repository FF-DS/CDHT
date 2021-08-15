import API from "@/services/api";

const rootPath = "/report";

const getNormalReports = () => {
  return API.get(`${rootPath}/normal`);
};

const getTestReports = () => {
  return API.get(`${rootPath}/test`);
};

const clearNormalReportCollection = () => {
  return API.get(`${rootPath}/clear-normal`);
};

const clearTestReportCollection = () => {
  return API.get(`${rootPath}/clear-test`);
};

const reportAPI = {
  getNormalReports,
  getTestReports,
  clearNormalReportCollection,
  clearTestReportCollection,
};

export default reportAPI;
