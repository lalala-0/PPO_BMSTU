import { useState } from "react";

 // Импортируем базовый URL API
import { FinishInput } from "../../models/raceModel";
import { handleError } from "../errorHandler";
import api from "../api"; // Импортируем функцию для обработки ошибок

export const useFinishProcedure = (ratingID: string, raceID: string) => {
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<string | null>(null);

  const finishProcedure = async (input: FinishInput) => {
    setLoading(true);
    setError(null);
    setSuccess(null);

    try {
      const response = await api.post(
        `/ratings/${ratingID}/races/${raceID}/finish`,
        input,
      );
      setSuccess(response.data || "Процедура финиша успешно выполнена");
    } catch (err: any) {
      handleError(err, setError);
    } finally {
      setLoading(false);
    }
  };

  return { loading, error, success, finishProcedure };
};
