import { useState } from "react";

import { RaceInput } from "../../models/raceModel"; // Импортируем модель для данных гонки
import { handleError } from "../errorHandler"; // Импортируем API_URL
import api from "../api"; // Импортируем функцию для обработки ошибок

export const useCreateRace = () => {
  const [input, setInput] = useState<RaceInput>({
    date: "",
    number: 0,
    class: 0,
  });
  const [success, setSuccess] = useState<string | null>(null);
  const [err, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState<boolean>(false); // Для индикации загрузки

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>,
  ) => {
    const { name, value } = e.target;
    setInput((prevInput) => ({
      ...prevInput,
      [name]: value,
    }));
  };

  const handleSubmit = async (ratingID: string) => {
    setLoading(true);
    setSuccess(null);
    setError(null);

    try {
      await api.post(`/ratings/${ratingID}/races`, input); // Отправляем данные на сервер
      setSuccess("Гонка успешно создана");
      setInput({ date: "", number: 0, class: 0 }); // Сброс формы после успешного создания
    } catch (err: any) {
      handleError(err, setError);
    } finally {
      setLoading(false);
    }
  };

  return { input, success, loading, handleChange, handleSubmit };
};
