import React, { useEffect, useState } from "react";
import axios from "axios";
import RatingsTable from "./tableRatings";
import RatingModal from "./modalInputRating";
import Filters from "../filters";
import { useDeleteRatingController } from "../../controllers/ratingControllers/deleteRatingController";
import { Rating } from "../../models/ratingModel";
import "../../styles/styles.css";

const RatingsContainer = () => {
  const [ratings, setRatings] = useState<Rating[]>([]);
  const [filters, setFilters] = useState({
    name: "",
    class: "",
    blowoutCnt: "",
  });
  const [selectedRating, setSelectedRating] = useState<Rating | null>(null);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [error, setError] = useState<string | null>(null); // Для хранения ошибки

  useEffect(() => {
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

  const handleFilterChange = (key: string, value: string) => {
    setFilters((prevFilters) => ({ ...prevFilters, [key]: value }));
  };

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
  };

  return (
    <div>
      {/* Фильтры */}
      <Filters filters={filters} onFilterChange={handleFilterChange} />

      {/* Таблица */}
      <RatingsTable
        ratings={filteredRatings}
        onDelete={handleDeleteRating}
        onUpdate={(rating) => {
          setSelectedRating(rating);
          setIsModalOpen(true);
        }}
      />

      {/* Кнопка для создания нового рейтинга */}
      <button
        onClick={() => {
          setSelectedRating({ id: "", Name: "", Class: "", BlowoutCnt: 0 });
          setIsModalOpen(true);
        }}
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
