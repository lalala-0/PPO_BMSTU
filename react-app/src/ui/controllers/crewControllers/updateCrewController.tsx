import { useState } from "react";
import axios from "axios";
import { CrewInput } from "../../models/crewModel";
import { API_URL } from "../../config";
import { useParams } from "react-router-dom";
import { handleError } from "../errorHandler";

export const useUpdateCrew = () => {
  const { ratingID } = useParams<{ ratingID: string }>();
  const [success, setSuccess] = useState<string | null>(null);
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  const handleUpdate = async (crewID: string, updatedData: CrewInput) => {
    setLoading(true);
    setSuccess(null);
    setError(null);

    try {
      const response = await axios.put(
        `${API_URL}/ratings/${ratingID}/crews/${crewID}`,
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
