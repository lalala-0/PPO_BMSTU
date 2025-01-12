import { useState } from "react";
import axios from "axios";
import { CrewParticipantAttachInput } from "../../models/crewModel";
import { API_URL } from "../../config";
import { handleError } from "../errorHandler"; // Импортируем функцию

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
      const response = await axios.post(
        `${API_URL}/ratings/${ratingID}/crews/${crewID}/members`,
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
