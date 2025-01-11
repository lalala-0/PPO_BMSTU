import { useState } from "react";
import axios from "axios";
import { API_URL } from "../../config";
import {
  ParticipantInput,
  ParticipantFormData,
} from "../../models/participantModel";
import { handleError } from "../errorHandler"; // Импортируем необходимые типы

export const useUpdateParticipant = (participantID: string) => {
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [updatedParticipant, setUpdatedParticipant] =
    useState<ParticipantFormData | null>(null);

  const updateParticipant = async (data: ParticipantInput) => {
    setLoading(true);
    setError(null);
    setUpdatedParticipant(null);

    try {
      const response = await axios.put<ParticipantFormData>(
        `${API_URL}/participants/${participantID}`,
        data,
      );
      setUpdatedParticipant(response.data); // Сохраняем обновленную информацию об участнике
    } catch (err: any) {
      handleError(err, setError); // Обработка ошибок через централизованную функцию
    } finally {
      setLoading(false); // Завершаем процесс загрузки
    }
  };

  return { updateParticipant, loading, error, updatedParticipant };
};
