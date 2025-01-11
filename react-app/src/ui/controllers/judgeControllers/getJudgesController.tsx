import { useState, useEffect } from "react";
import axios from "axios";
import { API_URL } from "../../config";
import { JudgeFilters, JudgeFormData } from "../../models/judgeModel";
import { handleError } from "../errorHandler"; // Импортируем централизованную обработку ошибок

export const useGetJudges = (filters: JudgeFilters = {}) => {
  const [judges, setJudges] = useState<JudgeFormData[]>([]);
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  const fetchJudges = async () => {
    setLoading(true);
    setError(null);

    try {
      const queryParams = new URLSearchParams(filters as any).toString();
      const { data } = await axios.get(`${API_URL}/judges?${queryParams}`);
      setJudges(data); // Сохраняем полученные данные судей
    } catch (err: any) {
      handleError(err, setError); // Обработка ошибок через централизованную функцию
    } finally {
      setLoading(false); // Завершаем процесс загрузки
    }
  };

  useEffect(() => {
    fetchJudges(); // Загружаем данные судей при изменении фильтров
  }, [filters]);

  return { judges, loading, error, fetchJudges };
};
