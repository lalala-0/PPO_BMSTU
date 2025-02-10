import { useState } from "react";
import {
  ParticipantFormData,
  ParticipantFilters,
} from "../../models/participantModel";
import { handleError } from "../errorHandler"; // Централизованная обработка ошибок
import api from "../api"; // Импорт API с обработкой ошибок

export const useGetAllParticipants = (filters: ParticipantFilters) => {
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [participants, setParticipants] = useState<ParticipantFormData[]>([]);

  const getParticipants = async () => {
    setLoading(true);
    setError(null);
    setParticipants([]);

    try {
      // Оставляем только те параметры, которые заданы (не `null` и не `undefined`)
      const params = Object.fromEntries(
        Object.entries(filters).filter(
          ([_, value]) => value !== null && value !== undefined && value !== "",
        ),
      );

      const { data } = await api.get<ParticipantFormData[]>(`/participants`, {
        params,
      });

      setParticipants(data); // Сохраняем полученных участников
    } catch (err: any) {
      handleError(err, setError); // Обрабатываем ошибку
    } finally {
      setLoading(false); // Завершаем процесс загрузки
    }
  };

  return { getParticipants, loading, error, participants };
};
