

import api from "../api"; // Импортируем функцию для обработки ошибок

export const useDeleteProtestController = () => {
  const handleDelete = async (
    ratingID: string,
    raceID: string,
    protestID: string,
  ) => {
    if (!ratingID || !raceID || !protestID) {
      alert("Некорректные параметры для удаления протеста");
      return;
    }

    try {
      await api.delete(
        `/ratings/${ratingID}/races/${raceID}/protests/${protestID}`,
      );
      alert("Протест успешно удалён");
    } catch {
      alert("Ошибка при удалении протеста");
    }
  };

  return { handleDelete };
};
