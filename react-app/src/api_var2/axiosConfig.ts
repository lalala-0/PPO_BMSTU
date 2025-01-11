import axios from "axios";

// Создаем экземпляр axios с базовым URL
const apiClient = axios.create({
  baseURL: "http://localhost:8081/api", // Указываем базовый URL
  headers: {
    "Content-Type": "application/json",
  },
});

export default apiClient;
