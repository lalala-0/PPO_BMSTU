import React, { useState, useEffect } from "react";
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
  const { handleDelete } = useDeleteProtestController();

  const [filters, setFilters] = useState<Record<string, string>>({});
  const [selectedProtest, setSelectedProtest] =
      useState<ProtestFormData | null>(null);
  const [isModalOpen, setIsModalOpen] = useState(() => {
    const saved = localStorage.getItem("isModalOpen");
    return saved ? JSON.parse(saved) : false;
  });

  const { protests, loading, error } = useFetchProtests(ratingID!, raceID!);
  const [filteredProtests, setFilteredProtests] = useState<ProtestFormData[]>([]);

  // Эффект для установки начальных протестов, когда они загружены
  useEffect(() => {
    if (!loading && protests.length > 0) {
      setFilteredProtests(protests);
    }
  }, [loading, protests]);

  // Обновляем локальное хранилище, когда модальное окно открыто или закрыто
  useEffect(() => {
    localStorage.setItem("isModalOpen", JSON.stringify(isModalOpen));
  }, [isModalOpen]);

  const handleFilterChange = (key: string, value: string) => {
    setFilters((prev) => {
      const newFilters = { ...prev, [key]: value };
      // Фильтруем протесты после изменения фильтра
      const updatedFilteredProtests = protests.filter(
          (protest) =>
              (newFilters.ruleNum
                  ? protest.RuleNum.toString()
                      .toLowerCase()
                      .includes(newFilters.ruleNum.toLowerCase())
                  : true) &&
              (newFilters.reviewDate
                  ? protest.ReviewDate.toLowerCase().includes(
                      newFilters.reviewDate.toLowerCase()
                  )
                  : true) &&
              (newFilters.status
                  ? StatusMap[Number(protest.Status)]
                      .toLowerCase()
                      .includes(newFilters.status.toLowerCase())
                  : true) &&
              (newFilters.comment
                  ? protest.Comment.toLowerCase().includes(
                      newFilters.comment.toLowerCase()
                  )
                  : true)
      );
      setFilteredProtests(updatedFilteredProtests);
      return newFilters;
    });
  };

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
      // Обновляем список протестов после удаления
      setFilteredProtests((prevProtests) =>
          prevProtests.filter((protest) => protest.ID !== protestID)
      );
    }
  };

  if (loading) return <div>Загрузка...</div>;
  if (error) return <div>Ошибка при загрузке протестов</div>;

  return (
      <div className="tableContent">
        <table className="protests-table">
          <thead>
          <tr>
            {[{ key: "ruleNum", label: "Номер правила" },
              { key: "reviewDate", label: "Дата рассмотрения" },
              { key: "status", label: "Статус" },
              { key: "comment", label: "Комментарий" }]
                .map(({ key, label }) => (
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
              <tr key={protest.ID}>
                <td>{protest.RuleNum}</td>
                <td>{protest.ReviewDate}</td>
                <td>{StatusMap[Number(protest.Status)]}</td>
                <td>{protest.Comment}</td>
                <td>
                  <button
                      onClick={() => handleNavigate(protest.ID)}
                      className="link-button"
                  >
                    Подробнее
                  </button>
                  <div className="buttons-container">
                    <button onClick={() => handleOpenModal(protest)}>
                      <img
                          src="/update-icon.svg"
                          alt="Обновить"
                          width="20"
                          height="20"
                      />
                    </button>
                    <button onClick={() => handleDeleteProtest(protest.ID)}>
                      <img
                          src="/delete-icon.svg"
                          alt="Удалить"
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
