import React, { useState } from "react";
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

  // Состояния для фильтров
  const [nameFilter, setNameFilter] = useState("");
  const [classFilter, setClassFilter] = useState("");
  const [blowoutFilter, setBlowoutFilter] = useState("");

  const handleNavigate = (id: string) => {
    navigate(`/ratings/${id}`);
  };

  // Фильтрация данных
  const filteredRatings = ratings.filter(
      (rating) =>
          rating.Name.toLowerCase().includes(nameFilter.toLowerCase()) &&
          rating.Class.toLowerCase().includes(classFilter.toLowerCase()) &&
          (blowoutFilter === "" ||
              rating.BlowoutCnt.toString().includes(blowoutFilter)),
  );

  return (
      <div className="table-container">
        {/* Скроллируемый контейнер для таблицы */}
        <div className="tableContent">
          <table className="table">
            <thead>
            <tr>
              <th className={"stickyHeader"}>
                Имя
                <input
                    type="text"
                    placeholder="Поиск по имени"
                    value={nameFilter}
                    onChange={(e) => setNameFilter(e.target.value)}
                />
              </th>
              <th className={"stickyHeader"}>
                Класс
                <input
                    type="text"
                    placeholder="Поиск по классу"
                    value={classFilter}
                    onChange={(e) => setClassFilter(e.target.value)}
                />
              </th>
              <th className={"stickyHeader"}>
                Кол-во выбрасываемых результатов
                <input
                    type="number"
                    placeholder="Поиск по кол-ву выбрасываемых результатов"
                    value={blowoutFilter}
                    onChange={(e) => setBlowoutFilter(e.target.value)}
                />
              </th>
              <th className={"stickyHeader"}>
                <th className="auth-required">Действия</th>
              </th>
            </tr>
            </thead>
            <tbody>
            {filteredRatings.map((rating) => (
                <tr key={rating.id}>
                  <td>
                    <button
                        onClick={() => handleNavigate(rating.id)}
                        className="add-rating-button"
                    >
                      {rating.Name}
                    </button>
                  </td>
                  <td>{rating.Class}</td>
                  <td>{rating.BlowoutCnt}</td>
                  <td className="auth-required">
                    <div className="buttons-container">
                      <button
                          onClick={() => onDelete(rating.id)}
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
