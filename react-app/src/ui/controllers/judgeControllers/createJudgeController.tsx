import { useState } from "react";
import axios from "axios";
import { API_URL } from "../../config";
import { JudgeInput } from "../../models/judgeModel";
import { handleError } from "../errorHandler"; // Импортируем функцию для обработки ошибок

export const useCreateJudge = () => {
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [data, setData] = useState<any | null>(null);

  const createJudge = async (judgeData: JudgeInput) => {
    setLoading(true);
    setError(null);

    try {
      const { data } = await axios.post(`${API_URL}/judges`, judgeData);
      setData(data); // Сохраняем данные судьи
    } catch (err: any) {
      handleError(err, setError); // Используем централизованную обработку ошибок
    } finally {
      setLoading(false); // Завершаем загрузку
    }
  };

  return { createJudge, loading, error, data };
};
