import { useState } from "react";
import axios from "axios";
import { API_URL } from "../../config";
import { ProtestParticipantAttachInput } from "../../models/protestModel";
import { handleError } from "../errorHandler";

export const useAttachProtestMember = (
  ratingID: string,
  raceID: string,
  protestID: string,
) => {
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [successMessage, setSuccessMessage] = useState<string | null>(null);

  const attachProtestMember = async (
    protestParticipantInput: ProtestParticipantAttachInput,
  ) => {
    setLoading(true);
    setError(null);
    setSuccessMessage(null);

    try {
      const response = await axios.post<{ [key: string]: string }>(
        `${API_URL}/ratings/${ratingID}/races/${raceID}/protests/${protestID}/members`,
        protestParticipantInput,
      );
      setSuccessMessage(
        response.data?.message || "Команда-участник успешно добавлена",
      );
    } catch (err: any) {
      handleError(err, setError); // Обработка ошибок через централизованную функцию
    } finally {
      setLoading(false);
    }
  };

  return { attachProtestMember, successMessage, loading, error };
};
