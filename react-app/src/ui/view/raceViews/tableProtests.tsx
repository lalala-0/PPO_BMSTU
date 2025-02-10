import React, { useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { useFetchProtests } from "../../controllers/protestControllers/getProtestsController";
import { ProtestFormData, StatusMap } from "../../models/protestModel";
import UpdateProtestModal from "../protestViews/modalUpdateProtest";
import { useDeleteProtestController } from "../../controllers/protestControllers/deleteProtestController";

const ProtestsTable: React.FC = () => {
  const { ratingID, raceID } = useParams<{
    ratingID: string;
    raceID: string;
  }>();
  const navigate = useNavigate();
  const { handleDelete } = useDeleteProtestController(); // Функция удаления

  const [filters, setFilters] = useState<Record<string, string>>({});
  const [selectedProtest, setSelectedProtest] =
    useState<ProtestFormData | null>(null);
  const [isModalOpen, setIsModalOpen] = useState(false);

  const { protests, loading, error } = useFetchProtests(ratingID!, raceID!);

  const handleFilterChange = (key: string, value: string) => {
    setFilters((prev) => ({
      ...prev,
      [key]: value,
    }));
  };

  // Фильтрация, аналогичная RatingsTable
  const filteredProtests = protests.filter(
    (protest) =>
      (filters.ruleNum
        ? protest.RuleNum.toString()
            .toLowerCase()
            .includes(filters.ruleNum.toLowerCase())
        : true) &&
      (filters.reviewDate
        ? protest.ReviewDate.toLowerCase().includes(
            filters.reviewDate.toLowerCase(),
          )
        : true) &&
      (filters.status
        ? StatusMap[Number(protest.Status)]
            .toLowerCase()
            .includes(filters.status.toLowerCase())
        : true) &&
      (filters.comment
        ? protest.Comment.toLowerCase().includes(filters.comment.toLowerCase())
        : true),
  );

  const handleNavigate = (id: string) => {
    navigate(`/ratings/${ratingID}/races/${raceID}/protests/${id}`);
  };

  const handleOpenModal = (protest: ProtestFormData) => {
    setSelectedProtest(protest);
    setIsModalOpen(true);
  };

  const handleCloseModal = () => {
    setIsModalOpen(false);
    setSelectedProtest(null);
  };

  const handleDeleteProtest = async (protestID: string) => {
    if (window.confirm("Вы уверены, что хотите удалить этот протест?")) {
      await handleDelete(ratingID || "", raceID || "", protestID);
      window.location.reload(); // Перезагрузка после удаления
    }
  };

  if (loading) return <div>Загрузка...</div>;
  if (error) return <div>Ошибка при загрузке протестов</div>;

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
            <th className="auth-required">Действия</th>
          </tr>
        </thead>
        <tbody>
          {filteredProtests.map((protest) => (
            <tr key={protest.ID}>
              <td>{protest.RuleNum}</td>
              <td>{protest.ReviewDate}</td>
              <td>{StatusMap[Number(protest.Status)]}</td>
              <td>{protest.Comment}</td>
              <td className="auth-required">
                <button
                  onClick={() => handleNavigate(protest.ID)}
                  className="link-button"
                >
                  Подробнее
                </button>
                <button
                  className="auth-required"
                  onClick={() => handleOpenModal(protest)}
                >
                  Обновить
                </button>
                <button
                  className="auth-required"
                  onClick={() => handleDeleteProtest(protest.ID)}
                >
                  Удалить
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>

      {isModalOpen && selectedProtest && (
        <UpdateProtestModal
          protest={{
            judgeId: selectedProtest.JudgeID,
            ruleNum: Number(selectedProtest.RuleNum),
            reviewDate: selectedProtest.ReviewDate,
            status: Number(selectedProtest.Status),
            comment: selectedProtest.Comment,
          }}
          onClose={handleCloseModal}
        />
      )}
    </div>
  );
};

export default ProtestsTable;
