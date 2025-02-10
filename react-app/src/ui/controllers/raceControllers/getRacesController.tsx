// getRacesByRatingID.ts
import { useState } from "react";
import axios from "axios";
import { RaceFormData } from "../../models/raceModel"; // Импортируем модель для гонок
import api from "../api"; // Импортируем функцию для обработки ошибок
import { handleError } from "../errorHandler"; // Импортируем функцию для обработки ошибок

export const useGetRacesByRatingID = () => {
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [races, setRaces] = useState<RaceFormData[] | null>(null);

  const getRacesByRatingID = async (ratingID: string): Promise<void> => {
    setLoading(true);
    setError(null); // Сбрасываем ошибку перед запросом
    setRaces(null); // Сбрасываем список гонок перед новым запросом

    try {
      const response = await api.get(`/ratings/${ratingID}/races`);
      setRaces(response.data); // Устанавливаем полученные данные
    } catch (err: any) {
      handleError(err, setError); // Используем обработчик ошибок
    } finally {
      setLoading(false); // Завершаем загрузку
    }
  };

  return {
    loading,
    error,
    races,
    getRacesByRatingID,
  };
};
