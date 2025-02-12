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



// // Создаём мокированный API (функции вместо реальных HTTP-запросов)
// const api = {
//   put: jest.fn().mockResolvedValue({}),  // Это мока для POST-запроса
//   post: jest.fn().mockReturnThis(),    // Это мока для метода create
//   get: jest.fn().mockResolvedValue({ data: {} }),  // Мока для GET-запроса
//   delete: jest.fn().mockResolvedValue({}),  // Мока для DELETE-запроса
//   interceptors: {
//     request: {
//       use: jest.fn(),  // Мока для интерцепторов
//     },
//   },
// };
//
// // Интерцептор для автоматической отправки токена (с помощью мока)
// api.interceptors.request.use = jest.fn((config) => {
//   const token = sessionStorage.getItem("token");
//   if (token) {
//     config.headers = config.headers || {};
//     config.headers.Authorization = `Bearer ${token}`;
//   }
//   return config;
// });