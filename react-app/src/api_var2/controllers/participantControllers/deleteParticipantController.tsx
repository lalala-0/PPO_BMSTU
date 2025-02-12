import { useState } from "react";


import { handleError } from "../errorHandler"; // Импортируем централизованную обработку ошибок
import api from "../api"; // Импортируем функцию для обработки ошибок

export const useDeleteParticipant = () => {
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<boolean | null>(null);

  const deleteParticipant = async (participantID: string) => {
    setLoading(true);
    setError(null);
    setSuccess(null);

    try {
      await api.delete(`/participants/${participantID}`);
      setSuccess(true); // Участник успешно удален
    } catch (err: any) {
      handleError(err, setError); // Обработка ошибок через централизованную функцию
    } finally {
      setLoading(false); // Завершаем процесс загрузки
    }
  };

  return { deleteParticipant, loading, error, success };
};
