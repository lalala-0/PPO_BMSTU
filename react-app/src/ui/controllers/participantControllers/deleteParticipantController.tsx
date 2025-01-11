import { useState } from "react";
import axios from "axios";
import { API_URL } from "../../config";
import { handleError } from "../errorHandler"; // Импортируем централизованную обработку ошибок

export const useDeleteParticipant = (participantID: string) => {
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<boolean | null>(null);

  const deleteParticipant = async () => {
    setLoading(true);
    setError(null);
    setSuccess(null);

    try {
      await axios.delete(`${API_URL}/participants/${participantID}`);
      setSuccess(true); // Участник успешно удален
    } catch (err: any) {
      handleError(err, setError); // Обработка ошибок через централизованную функцию
    } finally {
      setLoading(false); // Завершаем процесс загрузки
    }
  };

  return { deleteParticipant, loading, error, success };
};
