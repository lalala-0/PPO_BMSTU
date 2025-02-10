import { useState, useEffect } from "react";
import axios from "axios";
import { API_URL } from "../../config";
import { JudgeFormData } from "../../models/judgeModel";
import { handleError } from "../errorHandler"; // Импортируем функцию для обработки ошибок
import api from "../api"; // Импортируем функцию для обработки ошибок

export const useGetJudge = (judgeID: string) => {
  const [judge, setJudge] = useState<JudgeFormData | null>(null);
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  const fetchJudge = async () => {
    setLoading(true);
    setError(null);

    try {
      const { data } = await api.get(`/judges/${judgeID}`);
      setJudge(data); // Сохраняем данные судьи в состоянии
    } catch (err: any) {
      handleError(err, setError); // Обработка ошибок через централизованную функцию
    } finally {
      setLoading(false); // Завершаем загрузку
    }
  };

  useEffect(() => {
    if (judgeID) {
      fetchJudge(); // Загружаем данные судьи при изменении judgeID
    }
  }, [judgeID]);

  return { judge, loading, error, fetchJudge };
};
