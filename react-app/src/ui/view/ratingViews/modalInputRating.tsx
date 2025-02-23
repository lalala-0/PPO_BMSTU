import React, { useState, useEffect } from "react";
import { Rating } from "../../models/ratingModel";
import { classOptions } from "../../models/classOptions";
import { useUpdateRatingController } from "../../controllers/ratingControllers/updateRatingController";
import { useCreateRatingController } from "../../controllers/ratingControllers/createRatingController";
import { useParams } from "react-router-dom";

interface RatingModalProps {
  rating: Rating;
  type: "update" | "create";
  onClose: () => void;
}

const RatingModal: React.FC<RatingModalProps> = ({ rating, type, onClose }) => {
  const { ratingID } = useParams<{ ratingID: string }>();

  const [localRating, setLocalRating] = useState<Rating>({
    id: rating.id || "",
    Name: rating.Name || "",
    Class: rating.Class || classOptions[0].value.toString(),
    BlowoutCnt: rating.BlowoutCnt || 0,
  });
  const [errorMessage, setErrorMessage] = useState<string | null>(null);

  const { handleUpdate } = useUpdateRatingController();
  const { handleSubmit } = useCreateRatingController();

  // Восстановление данных формы из localStorage
  useEffect(() => {
    const storedRating = localStorage.getItem(`rating-${ratingID}`);
    if (storedRating) {
      setLocalRating(JSON.parse(storedRating));
    } else {
      setLocalRating({
        id: rating.id || "",
        Name: rating.Name || "",
        Class: rating.Class || classOptions[0].value.toString(),
        BlowoutCnt: rating.BlowoutCnt || 0,
      });
    }
  }, [ratingID, rating]);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setLocalRating((prev) => {
      const updatedRating = {
        ...prev,
        [name]: name === "BlowoutCnt" ? parseInt(value) : value,
      };
      // Сохраняем состояние в localStorage
      localStorage.setItem(`rating-${ratingID}`, JSON.stringify(updatedRating));
      return updatedRating;
    });
  };

  const handleSave = async () => {
    setErrorMessage(null); // Сбрасываем сообщение об ошибке

    const ratingData = {
      name: localRating.Name,
      class: parseInt(localRating.Class),
      blowout_cnt: localRating.BlowoutCnt,
    };

    try {
      if (type === "update") {
        await handleUpdate(localRating.id, ratingData);
      } else {
        await handleSubmit(ratingData);
      }
      onClose();
      // После сохранения очищаем состояние из localStorage
      localStorage.removeItem(`rating-${ratingID}`);
    } catch (error: any) {
      if (error.response && error.response.data) {
        setErrorMessage(
            error.response.data.message || "Ошибка при сохранении данных"
        );
      } else {
        setErrorMessage("Произошла ошибка. Попробуйте позже.");
      }
    }
  };

  return (
      <div className="modal-overlay">
        <div className="modal-content">
          <h3>{type === "update" ? "Обновить рейтинг" : "Создать новый рейтинг"}</h3>
          {errorMessage && (
              <div className="error-message">
                <p>{errorMessage}</p>
                <button onClick={() => setErrorMessage(null)}>Закрыть</button>
              </div>
          )}
          <input
              type="text"
              name="Name"
              value={localRating.Name}
              onChange={handleChange}
              placeholder="Название рейтинга"
          />
          <select
              name="Class"
              value={localRating.Class}
              onChange={handleChange}
          >
            {classOptions.map((option) => (
                <option key={option.value} value={option.value}>
                  {option.label}
                </option>
            ))}
          </select>
          <input
              type="number"
              name="BlowoutCnt"
              value={localRating.BlowoutCnt}
              onChange={handleChange}
              placeholder="Количество выбрасываемых результатов"
          />
          <div className="buttons-container">
            <button onClick={handleSave}>Сохранить изменения</button>
            <button onClick={onClose}>Отмена</button>
          </div>
        </div>
      </div>
  );
};

export default RatingModal;
