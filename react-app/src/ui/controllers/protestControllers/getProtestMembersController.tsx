import { useState, useEffect } from "react";
import axios from "axios";
import { API_URL } from "../../config";
import { ProtestCrewFormData } from "../../models/protestModel";
import { handleError } from "../errorHandler";

export const useGetProtestMembers = (
  ratingID: string,
  raceID: string,
  protestID: string,
) => {
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [protestMembers, setProtestMembers] = useState<ProtestCrewFormData[]>(
    [],
  );

  useEffect(() => {
    const fetchProtestMembers = async () => {
      setLoading(true);
      setError(null);

      try {
        const response = await axios.get<ProtestCrewFormData[]>(
          `${API_URL}/ratings/${ratingID}/races/${raceID}/protests/${protestID}/members`,
        );
        setProtestMembers(response.data);
      } catch (err: any) {
        handleError(err, setError); // Обработка ошибок через централизованную функцию
      } finally {
        setLoading(false);
      }
    };

    fetchProtestMembers();
  }, [ratingID, raceID, protestID]);

  return { protestMembers, loading, error };
};
