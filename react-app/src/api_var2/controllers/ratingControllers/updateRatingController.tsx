import { useState } from "react";

import { RatingInput } from "../../models/ratingModel";
import api from "../api";
import axios from "axios"; // Импортируем функцию для обработки ошибок

export const useUpdateRatingController = () => {
  const [success, setSuccess] = useState<string | null>(null);
  const [loading, setLoading] = useState<boolean>(false); // Для индикации загрузки

  const handleUpdate = async (id: string, updatedData: RatingInput) => {
    setLoading(true);
    setSuccess(null);

    try {
      await api.put(`/ratings/${id}`, updatedData);
      setSuccess("Рейтинг успешно обновлён");
    } catch (err: any) {
      if (axios.isAxiosError(err)) {
        if (err.response) {
          // Ошибки от сервера (HTTP 4xx или 5xx)
          const serverMessage =
            err.response.data?.message || "Неизвестная ошибка от сервера";
          alert(`Ошибка: ${serverMessage} (код: ${err.response.status})`);
        } else if (err.request) {
          // Проблемы с сетью
          alert("Ошибка сети. Проверьте подключение к интернету.");
        } else {
          // Ошибка при настройке запроса
          alert(`Ошибка запроса: ${err.message}`);
        }
      } else {
        // Непредвиденная ошибка
        alert("Произошла неизвестная ошибка.");
      }
    } finally {
      setLoading(false);
    }
  };

  return { success, loading, handleUpdate };
};
