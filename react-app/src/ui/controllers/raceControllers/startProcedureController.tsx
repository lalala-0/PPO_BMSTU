import { useState } from "react";
import axios from "axios";
import { StartInput } from "../../models/raceModel";
import { API_URL } from "../../config";
import { handleError } from "../errorHandler"; // Импортируем базовый URL API

export const useStartProcedure = (ratingID: string, raceID: string) => {
  const [loading, setLoading] = useState<boolean>(false); // Для индикации загрузки
  const [error, setError] = useState<string | null>(null); // Для отображения ошибок
  const [success, setSuccess] = useState<string | null>(null); // Для сообщения об успешном выполнении

  const startProcedure = async (input: StartInput) => {
    setLoading(true);
    setError(null);
    setSuccess(null);

    try {
      const response = await axios.post(
        `${API_URL}/ratings/${ratingID}/races/${raceID}/start`,
        input,
      );
      setSuccess(response.data || "Процедура старта успешно выполнена");
    } catch (err: any) {
      handleError(err, setError);
    } finally {
      setLoading(false);
    }
  };

  return { loading, error, success, startProcedure };
};
