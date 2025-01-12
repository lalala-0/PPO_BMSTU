// src/controllers/ApiDocController.ts

import axios from "axios";
import YAML from "js-yaml";
import { ApiData } from "../models/ApiData"; // Импортируем модель данных

const BASE_URL = "http://localhost:8081/api"; // Базовый URL для запросов

// Создание экземпляра axios с настройками по умолчанию
const axiosInstance = axios.create({
  baseURL: BASE_URL,
  headers: {
    "Content-Type": "application/json", // Устанавливаем правильный заголовок
  },
  timeout: 10000, // Устанавливаем тайм-аут для запросов
  withCredentials: true, // Добавляем это для отправки cookies
});

// Функция для загрузки документации API из YAML файла
export const fetchApiData = async (): Promise<ApiData> => {
  try {
    const response = await axios.get("/swagger.yaml"); // Путь к локальному файлу
    const doc = YAML.load(response.data);
    return doc as ApiData;
  } catch (error) {
    throw new Error("Error loading API doc");
  }
};

// Функция для выполнения API запроса
export const makeApiRequest = async (
  method: string,
  url: string,
  params: any,
) => {
  try {
    let response;

    // Проверяем, какой метод нужно использовать
    if (method.toUpperCase() === "GET") {
      response = await axios.get(url, {
        params: params, // Передаем параметры для GET-запроса
      });
    } else if (method.toUpperCase() === "POST") {
      response = await axios.post(url, params); // Для POST-запроса передаем параметры в теле
    } else if (method.toUpperCase() === "PUT") {
      response = await axios.put(url, params); // Для PUT-запроса
    } else if (method.toUpperCase() === "DELETE") {
      response = await axios.delete(url, {
        data: params, // Для DELETE-запроса передаем данные в теле
      });
    } else {
      throw new Error(`Unsupported HTTP method: ${method}`);
    }

    return response.data;
  } catch (error) {
    console.error("Error executing API request:", error);
    throw new Error("Error executing API request");
  }
};
