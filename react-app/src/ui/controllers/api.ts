import axios from "axios";

export const API_URL = "/api"; // Укажи свой адрес

const api = axios.create({
  baseURL: API_URL,
});

// Добавляем интерцептор для автоматической отправки токена
api.interceptors.request.use((config) => {
  const token = sessionStorage.getItem("token");
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export default api;
