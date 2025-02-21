import { useState } from "react";
import { CrewInput } from "../../models/crewModel";
import { handleError } from "../errorHandler";
import api from "../api";

export const useUpdateCrew = () => {
  const [success, setSuccess] = useState<string | null>(null);
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  const handleUpdate = async (ratingID: string, crewID: string, updatedData: CrewInput) => {
    setLoading(true);
    setSuccess(null);
    setError(null);

    try {
      await api.put(
        `/ratings/${ratingID}/crews/${crewID}`,
        updatedData,
      );
      setSuccess("Номер паруса успешно обновлён");
    } catch (err: any) {
      handleError(err, setError);
    } finally {
      setLoading(false);
    }
  };

  return { success, error, loading, handleUpdate };
};
