import { useState } from "react";
import axios from "axios";
import { API_URL } from "../../config";
import { ProtestFormData } from "../../models/protestModel";
import { handleError } from "../errorHandler";

export const useGetProtest = (
  ratingID: string,
  raceID: string,
  protestID: string,
) => {
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [protestData, setProtestData] = useState<ProtestFormData | null>(null);

  const getProtest = async () => {
    setLoading(true);
    setError(null);
    setProtestData(null);

    try {
      const response = await axios.get<ProtestFormData>(
        `${API_URL}/ratings/${ratingID}/races/${raceID}/protests/${protestID}`,
      );
      setProtestData(response.data);
    } catch (err: any) {
      handleError(err, setError); // Обработка ошибок через централизованную функцию
    } finally {
      setLoading(false);
    }
  };

  return { getProtest, protestData, loading, error };
};
