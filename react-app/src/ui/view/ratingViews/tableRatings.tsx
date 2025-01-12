import React from "react";
import { useNavigate } from "react-router-dom";
import { Rating } from "../../models/ratingModel";

interface RatingsTableProps {
  ratings: Rating[];
  onDelete: (id: string) => void;
  onUpdate: (rating: Rating) => void;
}

const RatingsTable: React.FC<RatingsTableProps> = ({
  ratings,
  onDelete,
  onUpdate,
}) => {
  const navigate = useNavigate();

  const handleNavigate = (id: string) => {
    navigate(`/ratings/${id}`);
  };

  return (
    <div style={{ maxWidth: "100%" }}>
      {/* Скроллируемый контейнер для таблицы */}
      <div
        style={{ maxHeight: "500px", overflowY: "auto", marginBottom: "20px" }}
      >
        <table style={{ tableLayout: "auto", width: "100%" }}>
          <thead>
            <tr>
              <th>Имя</th>
              <th>Класс</th>
              <th>Кол-во выбрасываемых результатов</th>
              <th>Действия</th>
            </tr>
          </thead>
          <tbody>
            {ratings.map((rating) => (
              <tr key={rating.id}>
                <td>
                  <button
                    onClick={() => handleNavigate(rating.id)}
                    className="add-rating-button" // Добавляем класс для стилизации
                  >
                    {rating.Name}
                  </button>
                </td>
                <td>{rating.Class}</td>
                <td>{rating.BlowoutCnt}</td>
                <td>
                  <div className="buttons-container">
                    <button
                      onClick={() => onDelete(rating.id)}
                      style={{
                        backgroundColor: "transparent",
                        border: "none",
                        cursor: "pointer",
                      }}
                    >
                      <img
                        src="/delete-icon.svg"
                        alt="Удалить"
                        width="20"
                        height="20"
                      />
                    </button>
                    <button
                      onClick={() => onUpdate(rating)}
                      style={{
                        backgroundColor: "transparent",
                        border: "none",
                        cursor: "pointer",
                        marginLeft: "10px",
                      }}
                    >
                      <img
                        src="/update-icon.svg"
                        alt="Обновить"
                        width="20"
                        height="20"
                      />
                    </button>
                  </div>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default RatingsTable;
