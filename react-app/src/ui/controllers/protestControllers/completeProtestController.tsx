import { useState } from "react";
import axios from "axios";
import { API_URL } from "../../config";
import { ProtestFormData } from "../../models/protestModel";
import { ProtestComplete } from "../../models/protestModel";
import { handleError } from "../errorHandler";

export const useCompleteProtest = (
  ratingID: string,
  raceID: string,
  protestID: string,
) => {
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [protestData, setProtestData] = useState<ProtestFormData | null>(null);

  const completeProtest = async (protestCompleteData: ProtestComplete) => {
    setLoading(true);
    setError(null);
    setProtestData(null);

    try {
      const response = await axios.patch<ProtestFormData>(
        `${API_URL}/ratings/${ratingID}/races/${raceID}/protests/${protestID}/complete`,
        protestCompleteData,
      );
      setProtestData(response.data);
    } catch (err: any) {
      handleError(err, setError); // Обработка ошибок через централизованную функцию
    } finally {
      setLoading(false);
    }
  };

  return { completeProtest, protestData, loading, error };
};
