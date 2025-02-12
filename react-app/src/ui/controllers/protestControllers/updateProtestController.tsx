import { useState } from "react";


import { ProtestInput, ProtestFormData } from "../../models/protestModel";
import { useParams } from "react-router-dom";
import { handleError } from "../errorHandler";
import api from "../api"; // Импортируем функцию для обработки ошибок

export const useUpdateProtest = () => {
  const { ratingID, raceID, protestID } = useParams<{
    ratingID: string;
    raceID: string;
    protestID: string;
  }>();

  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [updatedProtest, setUpdatedProtest] = useState<ProtestFormData | null>(
    null,
  );

  const updateProtest = async (protestInput: ProtestInput) => {
    if (!ratingID || !raceID || !protestID) {
      setError("Недостаточно данных для обновления протеста");
      return;
    }

    setLoading(true);
    setError(null);
    setUpdatedProtest(null);

    try {
      const response = await api.put<ProtestFormData>(
        `/ratings/${ratingID}/races/${raceID}/protests/${protestID}`,
        protestInput,
      );
      setUpdatedProtest(response.data);
    } catch (err: any) {
      handleError(err, setError);
    } finally {
      setLoading(false);
    }
  };

  return { updateProtest, updatedProtest, loading, error };
};
