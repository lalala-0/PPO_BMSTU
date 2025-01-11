// getCrewsByRatingID.ts
import { useState } from "react";
import axios from "axios";
import { CrewFormData } from "../../models/crewModel";
import { API_URL } from "../../config";
import { handleError } from "../errorHandler"; // Импортируем функцию для обработки ошибок

export const useGetCrewsByRatingID = () => {
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [crews, setCrews] = useState<CrewFormData[] | null>(null);

  const getCrewsByRatingID = async (ratingID: string): Promise<void> => {
    setLoading(true);
    setError(null); // Сбрасываем ошибку перед запросом
    setCrews(null); // Сбрасываем список команд перед новым запросом

    try {
      const response = await axios.get(`${API_URL}/${ratingID}/crews`);
      setCrews(response.data); // Устанавливаем полученные данные
    } catch (err: any) {
      handleError(err, setError); // Используем обработчик ошибок
    } finally {
      setLoading(false); // Завершаем загрузку
    }
  };

  return {
    loading,
    error,
    crews,
    getCrewsByRatingID,
  };
};
