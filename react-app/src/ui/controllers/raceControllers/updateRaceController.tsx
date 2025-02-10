import { useState } from "react";
import axios from "axios";
import { RaceInput } from "../../models/raceModel"; // Импортируем модель данных
import { useParams } from "react-router-dom";
import { handleError } from "../errorHandler"; // Импортируем функцию обработки ошибок
import api from "../api"; // Импортируем функцию для обработки ошибок

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
      await api.put(`/ratings/${ratingID}/races/${raceID}`, updatedData);
      setSuccess("Гонка успешно обновлена");
    } catch (err: any) {
      handleError(err, setError);
    } finally {
      setLoading(false);
    }
  };

  return { success, error, loading, handleUpdate };
};
