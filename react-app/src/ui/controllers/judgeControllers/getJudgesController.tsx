import { useState, useEffect } from "react";
import api from "../api"; // Импортируем API для запросов
import { JudgeFormData } from "../../models/judgeModel";

export const useGetJudges = () => {
  const [judges, setJudges] = useState<JudgeFormData[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const fetchJudges = async () => {
    setLoading(true);
    setError(null);
    try {
      const { data } = await api.get<JudgeFormData[]>("/judges");
      setJudges(data);
    } catch (err: any) {
      setError(err.message || "Ошибка при загрузке судей");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    if (judges.length === 0) {
      // Выполняем запрос только если данные не загружены
      fetchJudges();
    }
  }, [judges]); // Зависимость только от judges, чтобы запрос выполнялся один раз

  return { judges, loading, error, fetchJudges };
};
