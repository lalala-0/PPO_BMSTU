import api from "../api"; // Импортируем функцию для обработки ошибок
import { useState } from "react";
import axios from "axios";
import { RatingInput } from "../../models/ratingModel";

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
      await api.post("/api/ratings/", updatedData); // Передаем обновленные данные
      setSuccess("Рейтинг успешно создан");
      setInput({ name: "", class: 1, blowout_cnt: 0 });
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

  return { input, success, loading, handleChange, handleSubmit };
};
