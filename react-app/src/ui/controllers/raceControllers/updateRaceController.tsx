import { useState } from "react";
import axios from "axios";
import { RaceInput } from "../../models/raceModel"; // Импортируем модель данных
import { API_URL } from "../../config";
import { useParams } from "react-router-dom";
import { handleError } from "../errorHandler"; // Импортируем функцию обработки ошибок

export const useUpdateRace = () => {
  const { ratingID } = useParams<{ ratingID: string }>();
  const [success, setSuccess] = useState<string | null>(null);
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  const handleUpdate = async (raceID: string, updatedData: RaceInput) => {
    if (!ratingID) {
      setError("Отсутствует идентификатор рейтинга");
      return;
    }

    setLoading(true);
    setSuccess(null);
    setError(null);

    try {
      await axios.put(
        `${API_URL}/ratings/${ratingID}/races/${raceID}`,
        updatedData,
      );
      setSuccess("Гонка успешно обновлена");
    } catch (err: any) {
      handleError(err, setError);
    } finally {
      setLoading(false);
    }
  };

  return { success, error, loading, handleUpdate };
};
