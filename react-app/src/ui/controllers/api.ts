//Не для тестов

import axios from "axios";

export const API_URL = "/api"; // Укажи свой адрес

const api = axios.create({
  baseURL: process.env.REACT_APP_API_URL || API_URL,  // Используйте относительный путь или динамическое значение
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

