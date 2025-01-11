import { useState } from "react";
import axios from "axios";
import { useParams } from "react-router-dom"; // Импортируем useParams
import { API_URL } from "../../config";
import { handleError } from "../errorHandler"; // Импортируем функцию для обработки ошибок

export const useDetachCrewMember = () => {
  const { ratingID, crewID } = useParams<{
    ratingID: string;
    crewID: string;
  }>(); // Получаем параметры из URL
  const [success, setSuccess] = useState<string | null>(null);
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  const detachCrewMember = async (participantID: string) => {
    if (!ratingID || !crewID || !participantID) {
      setError("Недостаточно данных для выполнения запроса");
      return;
    }

    setLoading(true);
    setSuccess(null);
    setError(null);

    try {
      await axios.delete(
        `${API_URL}/ratings/${ratingID}/crews/${crewID}/members/${participantID}`,
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
