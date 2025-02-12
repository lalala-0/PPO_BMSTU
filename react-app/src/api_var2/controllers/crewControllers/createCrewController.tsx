import { useState } from "react";
import { CrewInput } from "../../models/crewModel";
import api from "../api";
import { handleError } from "../errorHandler"; // Импортируем функцию

export const useCreateCrew = (ratingID: string) => {
  const [input, setInput] = useState<CrewInput>({ sailNum: 0 }); // Инициализация данных для создания команды
  const [success, setSuccess] = useState<string | null>(null);
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setInput((prevInput) => ({
      ...prevInput,
      [name]: name === "sailNum" ? parseInt(value) || 0 : value,
    }));
  };

  const handleSubmit = async (updatedData: CrewInput) => {
    setLoading(true);
    setSuccess(null);
    setError(null);

    try {
      const response = await api.post(
          `/ratings/${ratingID}/crews`,
          updatedData,
      ); // Отправляем запрос на создание
      setSuccess("Команда успешно создана");
      setInput({ sailNum: 0 }); // Сброс после успешного создания
    } catch (err: any) {
      handleError(err, setError);
    } finally {
      setLoading(false);
    }
  };

  return {
    input,
    success,
    loading,
    error,
    handleChange,
    handleSubmit,
  };
};
