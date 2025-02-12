import { useState } from "react";


import { handleError } from "../errorHandler"; // Импортируем функцию для обработки ошибок
import api from "../api"; // Импортируем функцию для обработки ошибок

export const useDeleteJudge = () => {
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<boolean>(false);

  const deleteJudge = async (judgeID: string) => {
    setLoading(true);
    setError(null);
    setSuccess(false);

    try {
      await api.delete(`/judges/${judgeID}`);
      setSuccess(true); // Устанавливаем флаг успеха
    } catch (err: any) {
      handleError(err, setError); // Используем централизованную обработку ошибок
    } finally {
      setLoading(false); // Завершаем загрузку
    }
  };

  return { deleteJudge, loading, error, success };
};
