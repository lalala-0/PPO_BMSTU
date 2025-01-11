import axios from "axios";

export const useDeleteRatingController = () => {
  const handleDelete = async (id: string) => {
    if (!id) {
      alert("ID рейтинга не указан");
      return;
    }

    try {
      await axios.delete(`/api/ratings/${id}/`);
      alert("Рейтинг успешно удалён");
    } catch (err) {
      alert("Ошибка при удалении рейтинга");
    }
  };

  return { handleDelete };
};
