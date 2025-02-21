import api from "../api"; // Импортируем функцию для обработки ошибок

export const useDeleteRatingController = () => {
  const handleDelete = async (id: string) => {
    if (!id) {
      alert("ID рейтинга не указан");
      return;
    }

    try {
      await api.delete(`/ratings/${id}/`);
      alert("Рейтинг успешно удалён");
    } catch {
      alert("Ошибка при удалении рейтинга");
    }
  };

  return { handleDelete };
};
