import { useState } from "react";
import { RatingInput } from "../../models/ratingModel";
import api from "../api"; // Используем свой API

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
      // Проверяем на наличие поля response, если это ошибка с ответом от сервера
      if (err?.response) {
        const serverMessage =
            err.response.data?.message || "Неизвестная ошибка от сервера";
        alert(`Ошибка: ${serverMessage} (код: ${err.response.status})`);
      } else if (err?.request) {
        // Если нет ответа от сервера, но запрос был отправлен
        alert("Ошибка сети. Проверьте подключение к интернету.");
      } else {
        // Ошибка при настройке запроса
        alert(`Ошибка запроса: ${err.message}`);
      }
    } finally {
      setLoading(false); // Завершаем процесс загрузки
    }
  };

  return { success, loading, handleUpdate };
};
