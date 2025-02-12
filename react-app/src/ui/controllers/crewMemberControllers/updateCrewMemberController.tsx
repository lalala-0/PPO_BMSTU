import { useState } from "react";


import {
  ParticipantInput,
  ParticipantFormData,
} from "../../models/participantModel";
import { handleError } from "../errorHandler"; // Импортируем функцию для обработки ошибок
import api from "../api"; // Импортируем функцию для обработки ошибок

export const useUpdateCrewMember = (
  ratingID: string,
  crewID: string,
  participantID: string,
) => {
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [data, setData] = useState<ParticipantFormData | null>(null);

  const updateCrewMember = async (data: ParticipantInput) => {
    setLoading(true);
    setError(null);

    try {
      const response = await api.put<ParticipantFormData>(
        `/ratings/${ratingID}/crews/${crewID}/members/${participantID}`,
        data,
      );
      setData(response.data); // Сохраняем данные участника
    } catch (err: any) {
      handleError(err, setError); // Используем централизованную обработку ошибок
    } finally {
      setLoading(false); // Завершаем загрузку
    }
  };

  return { updateCrewMember, loading, error, data };
};
