import { useState } from "react";
import { ParticipantFormData } from "../../models/participantModel";
import { handleError } from "../errorHandler";
import api from "../api"; // Импортируем функцию для обработки ошибок

export const useGetCrewMember = (
) => {
  const [data, setData] = useState<ParticipantFormData | null>(null);
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  const getCrewMember = async (  ratingID: string, crewID: string, participantID: string) => {
    setLoading(true);
    setError(null);

    try {
      const response = await api.get<ParticipantFormData>(
        `/ratings/${ratingID}/crews/${crewID}/members/${participantID}`,
      );
      setData(response.data); // Сохраняем полученные данные
    } catch (err: any) {
      handleError(err, setError); // Используем централизованную обработку ошибок
    } finally {
      setLoading(false); // Завершаем загрузку
    }
  };

  return { data, loading, error, getCrewMember };
};
