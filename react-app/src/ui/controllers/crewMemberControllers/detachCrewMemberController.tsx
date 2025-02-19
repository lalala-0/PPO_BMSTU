import { useState } from "react";
import { handleError } from "../errorHandler";
import api from "../api"; // Импортируем функцию для обработки ошибок

export const useDetachCrewMember = () => {
  const [success, setSuccess] = useState<string | null>(null);
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  const detachCrewMember = async (ratingID: string, crewID: string, participantID: string) => {
    if (!ratingID || !crewID || !participantID) {
      setError("Недостаточно данных для выполнения запроса");
      return;
    }

    setLoading(true);
    setSuccess(null);
    setError(null);

    try {
      await api.delete(
        `/ratings/${ratingID}/crews/${crewID}/members/${participantID}`,
      );
      setSuccess("Участник успешно удален");
    } catch (err: any) {
      handleError(err, setError);
    } finally {
      setLoading(false);
    }
  };

  return { detachCrewMember, success, loading, error };
};
