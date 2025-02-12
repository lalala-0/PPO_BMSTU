// deleteCrewByID.ts
import { useState } from "react";
import { handleError } from "../errorHandler";
import { useParams } from "react-router-dom";
import api from "../api"; // Импортируем функцию для обработки ошибок

export const useDeleteCrew = () => {
  const { ratingID } = useParams<{ ratingID: string }>();
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<string | null>(null);

  const deleteCrewByID = async (CrewID: string): Promise<void> => {
    setLoading(true);
    setError(null); // Сбрасываем ошибку перед запросом
    setSuccess(null); // Сбрасываем успешный результат

    try {
      await api.delete(`/ratings/${ratingID}/crews/${CrewID}`);
      setSuccess("Команда успешно удалена");
    } catch (err: any) {
      handleError(err, setError); // Используем обработчик ошибок
    } finally {
      setLoading(false); // Завершаем загрузку
    }
  };

  return {
    loading,
    error,
    success,
    deleteCrewByID,
  };
};
