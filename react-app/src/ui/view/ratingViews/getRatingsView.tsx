import React, { useEffect, useState } from "react";
import axios from "axios";
import RatingsTable from "./tableRatings";
import RatingModal from "./modalInputRating";
import { useDeleteRatingController } from "../../controllers/ratingControllers/deleteRatingController";
import { Rating } from "../../models/ratingModel";

const RatingsContainer = () => {
  const [ratings, setRatings] = useState<Rating[]>([]);
  const [filters] = useState({
    name: "",
    class: "",
    blowoutCnt: "",
  });
  const [selectedRating, setSelectedRating] = useState<Rating | null>(null);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [error, setError] = useState<string | null>(null); // Для хранения ошибки

  useEffect(() => {
    // Восстановление состояния из localStorage
    const storedModalState = localStorage.getItem("modalState");
    if (storedModalState) {
      const parsedState = JSON.parse(storedModalState);
      setIsModalOpen(parsedState.isModalOpen);
      setSelectedRating(parsedState.selectedRating);
    }

    const fetchRatings = async () => {
      try {
        const response = await axios.get("api/ratings");
        setRatings(response.data);
      } catch (err) {
        console.error("Ошибка при загрузке данных:", err);
        setError("Ошибка при загрузке данных. Попробуйте обновить страницу.");
      }
    };

    fetchRatings();
  }, []);

  const filteredRatings = ratings.filter((rating) => {
    const { name, class: classFilter, blowoutCnt } = filters;
    return (
        rating.Name.toLowerCase().includes(name.toLowerCase()) &&
        (classFilter ? rating.Class.toString().includes(classFilter) : true) &&
        (blowoutCnt ? rating.BlowoutCnt.toString().includes(blowoutCnt) : true)
    );
  });

  const deleteRatingController = useDeleteRatingController();

  const handleDeleteRating = async (id: string) => {
    try {
      await deleteRatingController.handleDelete(id);
      setRatings((prev) => prev.filter((rating) => rating.id !== id)); // Удаляем из списка
    } catch (err) {
      console.error("Ошибка при удалении записи:", err);
      setError("Не удалось удалить запись. Попробуйте снова.");
    }
  };

  const closeModal = () => {
    setIsModalOpen(false);
    setSelectedRating(null);
    // Сохраняем состояние в localStorage
    localStorage.removeItem("modalState");
  };

  const openModal = (rating: Rating | null) => {
    setSelectedRating(rating);
    setIsModalOpen(true);
    // Сохраняем состояние в localStorage
    localStorage.setItem("modalState", JSON.stringify({ isModalOpen: true, selectedRating: rating }));
  };

  return (
      <div>
        {/* Таблица */}
        <RatingsTable
            ratings={filteredRatings}
            onDelete={handleDeleteRating}
            onUpdate={(rating) => openModal(rating)}
        />

        {/* Кнопка для создания нового рейтинга */}
        <button
            className="auth-required"
            onClick={() => openModal({ id: "", Name: "", Class: "", BlowoutCnt: 0 })}
        >
          Создать новый рейтинг
        </button>

        {/* Модальное окно */}
        {isModalOpen && selectedRating && (
            <RatingModal
                rating={selectedRating}
                type={selectedRating.id ? "update" : "create"}
                onClose={closeModal}
            />
        )}

        {/* Сообщение об ошибке */}
        {error && (
            <div className="error-message">
              <p>{error}</p>
              <button onClick={() => setError(null)}>Закрыть</button>
            </div>
        )}
      </div>
  );
};

export default RatingsContainer;
