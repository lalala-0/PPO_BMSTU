// src/models/ApiMethods.ts

import axios from "axios";
import YAML from "js-yaml";

// Функция для загрузки документации API из YAML файла
export const fetchApiData = async () => {
  try {
    const response = await axios.get("/swagger.yaml"); // Путь к локальному файлу
    const doc = YAML.load(response.data);
    return doc;
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
    const config = {
      method: method,
      url: url,
      data: params,
    };
    const response = await axios(config);
    return response.data;
  } catch (error) {
    throw new Error("Error executing API request");
  }
};
