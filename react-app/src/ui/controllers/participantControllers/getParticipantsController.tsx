import { useState } from "react";
import axios from "axios";
import { API_URL } from "../../config";
import { ParticipantFormData } from "../../models/participantModel";
import { ParticipantFilters } from "../../models/participantModel";
import { handleError } from "../errorHandler"; // Централизованная обработка ошибок

export const useGetAllParticipants = (filters: ParticipantFilters) => {
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [participants, setParticipants] = useState<ParticipantFormData[]>([]);

  const getParticipants = async () => {
    setLoading(true);
    setError(null);
    setParticipants([]);

    try {
      const { fio, category, gender, birthday, coach } = filters;
      const { data } = await axios.get<ParticipantFormData[]>(
        `${API_URL}/participants`,
        {
          params: {
            fio,
            category,
            gender,
            birthday,
            coach,
          },
        },
      );
      setParticipants(data); // Сохраняем полученных участников
    } catch (err: any) {
      handleError(err, setError); // Обработка ошибок через централизованную функцию
    } finally {
      setLoading(false); // Завершаем процесс загрузки
    }
  };

  return { getParticipants, loading, error, participants };
};
