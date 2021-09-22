import API from "@/services/api";

const rootPath = "/report";

const getAllReports = () => {
  return API.get(`${rootPath}/all/`, { limit: 20 });
};

const reportAPI = {
  getAllReports,
};

export default reportAPI;
