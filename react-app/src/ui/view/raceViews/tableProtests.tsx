import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useFetchProtests } from "../../controllers/protestControllers/getProtestsController";
import { ProtestFormData } from "../../models/protestModel";

interface ProtestsTableProps {
  ratingID: string;
  raceID: string;
}

const ProtestsTable: React.FC<ProtestsTableProps> = ({ ratingID, raceID }) => {
  const navigate = useNavigate();

  // Состояние фильтров
  const [filters, setFilters] = useState<Record<string, string>>({});

  const { protests, loading, error } = useFetchProtests(ratingID, raceID);

  // Обработчик изменения фильтров
  const handleFilterChange = (key: string, value: string) => {
    setFilters((prev) => ({
      ...prev,
      [key]: value,
    }));
  };

  // Универсальная фильтрация
  const filteredProtests: ProtestFormData[] =
    protests?.filter((protest: ProtestFormData) => {
      return Object.entries(filters).every(([key, value]) => {
        if (!value) return true; // Если фильтр пустой, пропускаем
        const protestValue = (protest as Record<string, any>)[key]; // Доступ к значению по ключу
        return protestValue
          ?.toString()
          .toLowerCase()
          .includes(value.toLowerCase());
      });
    }) || [];

  const handleNavigate = (id: string) => {
    navigate(`/protest/${id}`);
  };

  if (loading) {
    return <div>Загрузка...</div>;
  }

  if (error) {
    return <div>Ошибка при загрузке протестов</div>;
  }

  return (
    <div className="protests-table-container">
      <table className="protests-table">
        <thead>
          <tr>
            {[
              { key: "ruleNum", label: "Номер правила" },
              { key: "reviewDate", label: "Дата рассмотрения" },
              { key: "status", label: "Статус" },
              { key: "comment", label: "Комментарий" },
            ].map(({ key, label }) => (
              <th key={key}>
                <input
                  type="text"
                  placeholder={`Поиск по ${label.toLowerCase()}`}
                  value={filters[key] || ""}
                  onChange={(e) => handleFilterChange(key, e.target.value)}
                  style={{ width: "100%" }}
                />
                {label}
              </th>
            ))}
            <th>Действия</th>
          </tr>
        </thead>
        <tbody>
          {filteredProtests.map((protest) => (
            <tr key={protest.id}>
              <td>{protest.ruleNum}</td>
              <td>{protest.reviewDate}</td>
              <td>{protest.status}</td>
              <td>{protest.comment}</td>
              <td>
                <button
                  onClick={() => handleNavigate(protest.id)}
                  className="link-button"
                >
                  Подробнее
                </button>
                <button className="update-button">Обновить</button>
                <button className="delete-button">Удалить</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default ProtestsTable;
