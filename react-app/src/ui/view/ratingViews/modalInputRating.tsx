import React, { useState, useEffect } from "react";
import { Rating } from "../../models/ratingModel";
import { classOptions } from "../../models/classOptions";
import { useUpdateRatingController } from "../../controllers/ratingControllers/updateRatingController";
import { useCreateRatingController } from "../../controllers/ratingControllers/createRatingController";

interface RatingModalProps {
  rating: Rating;
  type: "update" | "create";
  onClose: () => void;
}

const RatingModal: React.FC<RatingModalProps> = ({ rating, type, onClose }) => {
  const [localRating, setLocalRating] = useState<Rating>(rating);
  const [errorMessage, setErrorMessage] = useState<string | null>(null); // Состояние для отображения ошибки
  const { handleUpdate } = useUpdateRatingController();
  const { handleSubmit } = useCreateRatingController();

  useEffect(() => {
    if (type === "create") {
      setLocalRating({
        ...rating,
        Class: classOptions[0].value.toString(), // Значение по умолчанию
        BlowoutCnt: 0, // Значение по умолчанию
      });
    } else {
      setLocalRating(rating); // Для обновления используем переданный рейтинг
    }
  }, [type, rating]);

  const handleNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setLocalRating((prev) => ({ ...prev, Name: e.target.value }));
  };

  const handleClassChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    setLocalRating((prev) => ({ ...prev, Class: e.target.value }));
  };

  const handleBlowoutCntChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = parseInt(e.target.value) || 0;
    setLocalRating((prev) => ({ ...prev, BlowoutCnt: value }));
  };

  const handleSave = async () => {
    setErrorMessage(null); // Сбрасываем сообщение об ошибке
    const ratingData = {
      name: localRating.Name,
      class: parseInt(localRating.Class),
      blowout_cnt: localRating.BlowoutCnt, // Приводим ключ в соответствие с сервером
    };

    try {
      if (type === "update") {
        await handleUpdate(localRating.id, ratingData);
      } else {
        await handleSubmit(ratingData);
      }
      onClose();
    } catch (error: any) {
      if (error.response && error.response.data) {
        setErrorMessage(
          error.response.data.message || "Ошибка при сохранении данных",
        );
      } else {
        setErrorMessage("Произошла ошибка. Попробуйте позже.");
      }
    }
  };

  const handleClose = () => {
    onClose();
    console.log("Модальное окно закрыто");
  };

  return (
    <div className="modal-overlay">
      <div className="modal-content">
        <h3>
          {type === "update" ? "Обновить рейтинг" : "Создать новый рейтинг"}
        </h3>
        {errorMessage && (
          <div className="error-message">
            <p>{errorMessage}</p>
            <button onClick={() => setErrorMessage(null)}>Закрыть</button>
          </div>
        )}
        <input
          type="text"
          value={localRating.Name}
          onChange={handleNameChange}
          placeholder="Название рейтинга"
        />
        <select value={localRating.Class} onChange={handleClassChange}>
          {classOptions.map((option) => (
            <option key={option.value} value={option.value}>
              {option.label}
            </option>
          ))}
        </select>
        <input
          type="number"
          value={localRating.BlowoutCnt}
          onChange={handleBlowoutCntChange}
          placeholder="Количество срывов"
        />
        <div className="buttons-container">
          <button onClick={handleSave}>Сохранить изменения</button>
          <button onClick={handleClose}>Отмена</button>
        </div>
      </div>
    </div>
  );
};

export default RatingModal;
