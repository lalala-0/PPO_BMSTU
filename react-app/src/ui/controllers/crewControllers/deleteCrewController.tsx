// deleteCrewByID.ts
import { useState } from "react";
import axios from "axios";
import { API_URL } from "../../config";
import { handleError } from "../errorHandler";
import { useParams } from "react-router-dom"; // Импортируем функцию для обработки ошибок

export const useDeleteCrew = () => {
  const { ratingID } = useParams<{ ratingID: string }>();
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<string | null>(null);

  const deleteCrewByID = async (crewID: string): Promise<void> => {
    setLoading(true);
    setError(null); // Сбрасываем ошибку перед запросом
    setSuccess(null); // Сбрасываем успешный результат

    try {
      await axios.delete(`${API_URL}/ratings/${ratingID}/crews/${crewID}`);
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
