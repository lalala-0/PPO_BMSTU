import { useState, useCallback } from "react";
import { CrewFormData } from "../../models/crewModel";
import { handleError } from "../errorHandler";
import api from "../api";

export const useGetCrewsByRatingID = () => {
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [crews, setCrews] = useState<CrewFormData[] | null>(null);

  const getCrewsByRatingID = useCallback(
    async (ratingID: string): Promise<void> => {
      setLoading(true);
      setError(null);
      setCrews(null);

      try {
        const response = await api.get(`/ratings/${ratingID}/crews`);
        setCrews(response.data);
      } catch (err: any) {
        handleError(err, setError);
      } finally {
        setLoading(false);
      }
    },
    [],
  );

  return {
    loading,
    error,
    crews,
    getCrewsByRatingID,
  };
};
