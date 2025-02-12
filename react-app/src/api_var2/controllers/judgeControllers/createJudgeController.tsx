import { useState } from "react";
import { JudgeInput } from "../../models/judgeModel";
import { handleError } from "../errorHandler"; // Импортируем функцию для обработки ошибок
import api from "../api"; // Импортируем функцию для обработки ошибок

export const useCreateJudge = () => {
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [data, setData] = useState<any | null>(null);

  const createJudge = async (judgeData: JudgeInput) => {
    setLoading(true);
    setError(null);

    try {
      const { data } = await api.post(`/judges`, judgeData);
      setData(data); // Сохраняем данные судьи
    } catch (err: any) {
      handleError(err, setError); // Используем централизованную обработку ошибок
    } finally {
      setLoading(false); // Завершаем загрузку
    }
  };

  return { createJudge, loading, error, data };
};
