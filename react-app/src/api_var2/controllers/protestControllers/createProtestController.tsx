import { useState } from "react";


import { ProtestCreate } from "../../models/protestModel";
import { ProtestFormData } from "../../models/protestModel";
import { handleError } from "../errorHandler";
import { useParams } from "react-router-dom";
import api from "../api"; // Импортируем функцию для обработки ошибок

export const useCreateProtest = () => {
  const { ratingID, raceID } = useParams<{
    ratingID: string;
    raceID: string;
  }>();
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [newProtest, setNewProtest] = useState<ProtestFormData | null>(null);

  const createProtest = async (protestCreate: ProtestCreate) => {
    setLoading(true);
    setError(null);
    setNewProtest(null);

    try {
      const response = await api.post(
        `/ratings/${ratingID}/races/${raceID}/protests`,
        protestCreate,
      );
      setNewProtest(response.data);
    } catch (err: any) {
      handleError(err, setError); // Обработка ошибок через централизованную функцию
    } finally {
      setLoading(false);
    }
  };

  return { createProtest, newProtest, loading, error };
};
