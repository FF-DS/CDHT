import API from "@/services/api";

const getTest = () => {
  return API.get("/test/");
};

const testingAPI = {
  getTest,
};

export default testingAPI;
