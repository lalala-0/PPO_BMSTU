import { useState } from "react";
import axios from "axios";
import { API_URL } from "../../config";
import { handleError } from "../errorHandler"; // Импортируем API_URL

export const useDeleteRace = (ratingID: string, raceID: string) => {
  const [loading, setLoading] = useState<boolean>(false); // Для индикации загрузки
  const [error, setError] = useState<string | null>(null); // Для отображения ошибок
  const [success, setSuccess] = useState<boolean>(false); // Для индикации успешного удаления

  const deleteRace = async () => {
    setLoading(true);
    setError(null);
    setSuccess(false);

    try {
      await axios.delete(`${API_URL}/ratings/${ratingID}/races/${raceID}`);
      setSuccess(true); // Успешное удаление
    } catch (err: any) {
      handleError(err, setError);
    } finally {
      setLoading(false);
    }
  };

  return { loading, error, success, deleteRace };
};
