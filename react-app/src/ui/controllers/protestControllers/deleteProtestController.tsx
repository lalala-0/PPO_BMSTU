import { useState } from "react";
import axios from "axios";
import { API_URL } from "../../config";
import { handleError } from "../errorHandler";

export const useDeleteProtest = (
  ratingID: string,
  raceID: string,
  protestID: string,
) => {
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [isDeleted, setIsDeleted] = useState<boolean>(false);

  const deleteProtest = async () => {
    setLoading(true);
    setError(null);
    setIsDeleted(false);

    try {
      await axios.delete(
        `${API_URL}/ratings/${ratingID}/races/${raceID}/protests/${protestID}`,
      );
      setIsDeleted(true);
    } catch (err: any) {
      handleError(err, setError); // Обработка ошибок через централизованную функцию
    } finally {
      setLoading(false);
    }
  };

  return { deleteProtest, isDeleted, loading, error };
};
