import { useState } from "react";
import axios from "axios";
import { API_URL } from "../../config";
import { ProtestParticipantDetachInput } from "../../models/protestModel";
import { handleError } from "../errorHandler";

export const useDetachProtestMember = (
  ratingID: string,
  raceID: string,
  protestID: string,
) => {
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  const detachProtestMember = async (
    protestParticipantDetachInput: ProtestParticipantDetachInput,
  ) => {
    setLoading(true);
    setError(null);

    try {
      await axios.delete(
        `${API_URL}/ratings/${ratingID}/races/${raceID}/protests/${protestID}/members/${protestParticipantDetachInput.sailNum}`,
      );
    } catch (err: any) {
      handleError(err, setError); // Обработка ошибок через централизованную функцию
    } finally {
      setLoading(false);
    }
  };

  return { detachProtestMember, loading, error };
};