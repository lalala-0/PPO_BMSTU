import { useState } from "react";
import axios from "axios";
import { API_URL } from "../../config";
import { JudgeInput, JudgeFormData } from "../../models/judgeModel"; // Импортируйте типы
import { handleError } from "../errorHandler"; // Импортируем централизованную обработку ошибок

export const useUpdateJudge = (judgeID: string) => {
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [updatedJudge, setUpdatedJudge] = useState<JudgeFormData | null>(null);

  const updateJudge = async (judgeData: JudgeInput) => {
    setLoading(true);
    setError(null);

    try {
      const { data } = await axios.put(
        `${API_URL}/judges/${judgeID}`,
        judgeData,
      );
      setUpdatedJudge(data); // Сохраняем обновленного судью
      return data;
    } catch (err: any) {
      handleError(err, setError); // Обработка ошибок через централизованную функцию
    } finally {
      setLoading(false); // Завершаем процесс загрузки
    }
  };

  return { updateJudge, loading, error, updatedJudge };
};
