import API from "@/services/api";

const rootPath = "/report";

const getAllReports = (requestBody) => {
  return API.post(`${rootPath}/all`, requestBody);
};

const reportAPI = {
  getAllReports,
};

export default reportAPI;
