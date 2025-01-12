import { useState } from "react";
import axios from "axios";
import { API_URL } from "../../config";
import {
  ParticipantInput,
  ParticipantFormData,
} from "../../models/participantModel"; // Импортируем типы для создания участника
import { handleError } from "../errorHandler"; // Импортируем централизованную обработку ошибок

export const useCreateParticipant = () => {
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [createdParticipant, setCreatedParticipant] =
    useState<ParticipantFormData | null>(null);

  const createParticipant = async (participantData: ParticipantInput) => {
    setLoading(true);
    setError(null);

    try {
      const { data } = await axios.post(
        `${API_URL}/participants`,
        participantData,
      );
      setCreatedParticipant(data); // Сохраняем данные о созданном участнике
      return data;
    } catch (err: any) {
      handleError(err, setError); // Обработка ошибок через централизованную функцию
    } finally {
      setLoading(false); // Завершаем процесс загрузки
    }
  };

  return { createParticipant, loading, error, createdParticipant };
};
