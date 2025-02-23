import { useState } from "react";
import { ProtestFormData, ProtestComplete } from "../../models/protestModel";
import { handleError } from "../errorHandler";
import api from "../api";

export const useCompleteProtest = () => {
    const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [protestData, setProtestData] = useState<ProtestFormData | null>(null);

  const completeProtest = async (ratingID : string, raceID: string, protestID: string, protestCompleteData: ProtestComplete) => {
    if (!ratingID || !raceID || !protestID) {
      setError("Недостаточно данных.");
      return;
    }

    setLoading(true);
    setError(null);
    setProtestData(null);

    try {
      const response = await api.post(
        `/ratings/${ratingID}/races/${raceID}/protests/${protestID}/complete`,
        protestCompleteData,
      );
      setProtestData(response.data);
    } catch (err: any) {
      handleError(err, setError);
    } finally {
      setLoading(false);
    }
  };

  return { completeProtest, protestData, loading, error };
};
