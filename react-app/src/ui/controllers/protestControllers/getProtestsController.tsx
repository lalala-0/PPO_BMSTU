import { useState, useEffect } from "react";
import axios from "axios";
import { API_URL } from "../../config";
import { ProtestFormData } from "../../models/protestModel";
import { handleError } from "../errorHandler";

export const useFetchProtests = (ratingID: string, raceID: string) => {
  const [protests, setProtests] = useState<ProtestFormData[]>([]);
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchProtests = async () => {
      setLoading(true);
      setError(null);

      try {
        const response = await axios.get<ProtestFormData[]>(
          `${API_URL}/ratings/${ratingID}/races/${raceID}/protests`,
        );
        setProtests(response.data);
      } catch (err: any) {
        handleError(err, setError); // Обработка ошибок через централизованную функцию
      } finally {
        setLoading(false);
      }
    };

    fetchProtests();
  }, [ratingID, raceID]);

  return { protests, loading, error };
};
