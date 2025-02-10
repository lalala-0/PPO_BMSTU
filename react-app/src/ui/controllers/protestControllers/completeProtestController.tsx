import { useState, useEffect } from "react";
import axios from "axios";
import { useParams } from "react-router-dom";
import { API_URL } from "../../config";
import { ProtestFormData, ProtestComplete } from "../../models/protestModel";
import { handleError } from "../errorHandler";
import api from "../api"; // Импортируем функцию для обработки ошибок

export const useCompleteProtest = () => {
  const { ratingID, raceID, protestID } = useParams<{
    ratingID: string;
    raceID: string;
    protestID: string;
  }>();

  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [protestData, setProtestData] = useState<ProtestFormData | null>(null);

  useEffect(() => {
    if (!ratingID || !raceID || !protestID) {
      setError("Недостаточно данных для завершения протеста.");
    }
  }, [ratingID, raceID, protestID]);

  const completeProtest = async (protestCompleteData: ProtestComplete) => {
    if (!ratingID || !raceID || !protestID) {
      setError("Недостаточно данных.");
      return;
    }

    setLoading(true);
    setError(null);
    setProtestData(null);

    try {
      const response = await axios.patch<ProtestFormData>(
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
