import { useState, useEffect, useCallback } from "react";
import { CrewFormData } from "../../models/crewModel";
import { handleError } from "../errorHandler";
import api from "../api";

export const useGetCrew = () => {
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [crewInfo, setCrewInfo] = useState<CrewFormData | null>(null);

  const getCrew = useCallback(
      async (ratingID: string, crewID: string): Promise<void> => {
        setLoading(true);
        setError(null);
        setCrewInfo(null);

        try {
          const response = await api.get(`/ratings/${ratingID}/crews/${crewID}`);
          setCrewInfo(response.data);
        } catch (err: any) {
          handleError(err, setError);
        } finally {
          setLoading(false);
        }
      },
      [],
  );

  return {
    crewInfo,
    loading,
    error,
    getCrew,
  };
};
