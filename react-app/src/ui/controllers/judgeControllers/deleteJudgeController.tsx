import { useState } from "react";
import axios from "axios";
import { API_URL } from "../../config";
import { handleError } from "../errorHandler"; // Импортируем функцию для обработки ошибок

export const useDeleteJudge = (judgeID: string) => {
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<boolean>(false);

  const deleteJudge = async () => {
    setLoading(true);
    setError(null);
    setSuccess(false);

    try {
      await axios.delete(`${API_URL}/judges/${judgeID}`);
      setSuccess(true); // Устанавливаем флаг успеха
    } catch (err: any) {
      handleError(err, setError); // Используем централизованную обработку ошибок
    } finally {
      setLoading(false); // Завершаем загрузку
    }
  };

  return { deleteJudge, loading, error, success };
};
