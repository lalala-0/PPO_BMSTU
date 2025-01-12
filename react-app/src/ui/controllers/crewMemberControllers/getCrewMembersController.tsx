import { useState } from "react";
import axios from "axios";
import { API_URL } from "../../config";
import { ParticipantFormData } from "../../models/participantModel";
import { handleError } from "../errorHandler"; // Импортируем функцию для обработки ошибок

export const useGetCrewMembers = (ratingID: string, crewID: string) => {
  const [data, setData] = useState<ParticipantFormData[] | null>(null); // Массив участников
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  const getCrewMembers = async () => {
    setLoading(true);
    setError(null);

    try {
      const response = await axios.get<ParticipantFormData[]>(
        `${API_URL}/ratings/${ratingID}/crews/${crewID}/members`,
      );
      setData(response.data); // Сохраняем полученные данные
    } catch (err: any) {
      handleError(err, setError); // Используем централизованную обработку ошибок
    } finally {
      setLoading(false); // Завершаем загрузку
    }
  };

  return { data, loading, error, getCrewMembers };
};
