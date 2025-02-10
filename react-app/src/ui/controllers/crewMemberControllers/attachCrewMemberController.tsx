import { useState } from "react";
import { CrewParticipantAttachInput } from "../../models/crewModel";
import { handleError } from "../errorHandler"; // Импортируем функцию
import api from "../api"; // Импортируем функцию для обработки ошибок

export const useAttachCrewMember = (ratingID: string, crewID: string) => {
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<string | null>(null);
  const [loading, setLoading] = useState<boolean>(false);

  const attachCrewMember = async (
    crewParticipantAttachInput: CrewParticipantAttachInput,
  ) => {
    setLoading(true);
    setSuccess(null);
    setError(null);

    try {
      const response = await api.post(
        `/ratings/${ratingID}/crews/${crewID}/members`,
        crewParticipantAttachInput,
      );
      setSuccess("Член экипажа успешно добавлен");
      return response.data;
    } catch (err: any) {
      handleError(err, setError);
    } finally {
      setLoading(false);
    }
  };

  return { attachCrewMember, success, loading, error };
};
