import { useState } from "react";
import api from "../api"; // Импортируем свой API
import { RatingInput } from "../../models/ratingModel"; // Импортируем модель данных

export const useCreateRatingController = () => {
  const [input, setInput] = useState<RatingInput>({
    name: "",
    class: 1,
    blowout_cnt: 0,
  });
  const [success, setSuccess] = useState<string | null>(null);
  const [loading, setLoading] = useState<boolean>(false); // Для индикации загрузки

  const handleChange = (
      e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>,
  ) => {
    const { name, value } = e.target;
    setInput((prevInput) => ({
      ...prevInput,
      [name]: name === "class" ? parseInt(value) : value,
    }));
  };

  const handleSubmit = async (updatedData: RatingInput) => {
    setLoading(true);
    setSuccess(null);

    try {
      console.log("Submitting data:", updatedData); // Логируем отправляемые данные для дебага
      const response = await api.post("/ratings/", updatedData); // Передаем обновленные данные
      console.log("Server response:", response); // Логируем ответ от сервера

      setSuccess("Рейтинг успешно создан");
      setInput({ name: "", class: 1, blowout_cnt: 0 }); // Сброс состояния после успешного создания
    } catch (err: any) {
      // Стандартная обработка ошибок
      if (err?.response) {
        // Ошибки от сервера (HTTP 4xx или 5xx)
        const serverMessage =
            err.response.data?.message || "Неизвестная ошибка от сервера";
        alert(`Ошибка: ${serverMessage} (код: ${err.response.status})`);
      } else if (err?.request) {
        // Проблемы с сетью
        alert("Ошибка сети. Проверьте подключение к интернету.");
      } else {
        // Ошибка при настройке запроса
        alert(`Ошибка запроса: ${err.message}`);
      }
    } finally {
      setLoading(false); // Завершаем процесс загрузки
    }
  };

  return { input, success, loading, handleChange, handleSubmit };
};
