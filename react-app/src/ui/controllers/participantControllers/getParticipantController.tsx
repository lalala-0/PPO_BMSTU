import { useCallback, useState } from "react";
import axios from "axios";
import { API_URL } from "../../config";
import { ParticipantFormData } from "../../models/participantModel";
import { handleError } from "../errorHandler"; // Импортируем централизованную обработку ошибок

export const useGetParticipant = (participantID: string) => {
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [participant, setParticipant] = useState<ParticipantFormData | null>(
    null,
  );

  const getParticipant = useCallback(async () => {
    setLoading(true);
    setError(null);
    setParticipant(null);

    try {
      const { data } = await axios.get<ParticipantFormData>(
        `${API_URL}/participants/${participantID}`,
      );
      setParticipant(data); // Сохраняем данные об участнике
    } catch (err: any) {
      handleError(err, setError); // Обработка ошибок через централизованную функцию
    } finally {
      setLoading(false); // Завершаем процесс загрузки
    }
  }, [participantID]);

  return { getParticipant, loading, error, participant };
};