
// Создаём мокированный API (функции вместо реальных HTTP-запросов)
const api = {
  put: jest.fn().mockResolvedValue({}),  // Это мока для POST-запроса
  post: jest.fn().mockReturnThis(),    // Это мока для метода create
  get: jest.fn().mockResolvedValue({ data: {} }),  // Мока для GET-запроса
  delete: jest.fn().mockResolvedValue({}),  // Мока для DELETE-запроса
  interceptors: {
    request: {
      use: jest.fn(),  // Мока для интерцепторов
    },
  },
};

// Интерцептор для автоматической отправки токена (с помощью мока)
api.interceptors.request.use = jest.fn((config) => {
  const token = sessionStorage.getItem("token");
  if (token) {
    config.headers = config.headers || {};
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export default api;
