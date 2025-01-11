import { useState } from "react";
import axios from "axios";
import { API_URL } from "../../config";
import { ProtestInput } from "../../models/protestModel";
import { ProtestFormData } from "../../models/protestModel";
import { handleError } from "../errorHandler";

export const useUpdateProtest = (
  ratingID: string,
  raceID: string,
  protestID: string,
) => {
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [updatedProtest, setUpdatedProtest] = useState<ProtestFormData | null>(
    null,
  );

  const updateProtest = async (protestInput: ProtestInput) => {
    setLoading(true);
    setError(null);
    setUpdatedProtest(null);

    try {
      const response = await axios.put<ProtestFormData>(
        `${API_URL}/ratings/${ratingID}/races/${raceID}/protests/${protestID}`,
        protestInput,
      );
      setUpdatedProtest(response.data);
    } catch (err: any) {
      handleError(err, setError); // Обработка ошибок через централизованную функцию
    } finally {
      setLoading(false);
    }
  };

  return { updateProtest, updatedProtest, loading, error };
};
