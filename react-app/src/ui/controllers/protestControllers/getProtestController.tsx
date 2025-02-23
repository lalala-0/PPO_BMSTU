import { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import { ProtestFormData } from "../../models/protestModel";
import { handleError } from "../errorHandler";
import api from "../api"; // Импортируем функцию для обработки ошибок

export const useGetProtest = () => {
  const { ratingID, raceID, protestID } = useParams<{
    ratingID: string;
    raceID: string;
    protestID: string;
  }>();

  const [protestInfo, setProtestInfo] = useState<ProtestFormData | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!ratingID || !raceID || !protestID) {
      setError("Недостаточно данных для загрузки протеста");
      setLoading(false);
      return;
    }

    const fetchProtestData = async () => {
      try {
        const response = await api.get<ProtestFormData>(
          `/ratings/${ratingID}/races/${raceID}/protests/${protestID}`,
        );
        setProtestInfo(response.data);
      } catch (err: any) {
        handleError(err, setError);
      } finally {
        setLoading(false);
      }
    };

    fetchProtestData();
  }, [ratingID, raceID, protestID]);

  return {
    protestInfo,
    loading,
    error,
  };
};
