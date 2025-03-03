import { useParams } from "react-router-dom";
import { useGetRating } from "../../controllers/ratingControllers/getRatingController";
import { useGetRace } from "../../controllers/raceControllers/getRaceController";
import { useGetProtest } from "../../controllers/protestControllers/getProtestController";
import { useState, useEffect } from "react";
import { StatusMap } from "../../models/protestModel";
import UpdateProtestModal from "./modalUpdateProtest";
import { CompleteProtestModal } from "./modalComplete";
import React from "react";
import ProtestCrewMembersTable from "./tableProtestMembers";

const ProtestView: React.FC = () => {
  const { ratingID, raceID, protestID } = useParams<{
    ratingID: string;
    raceID: string;
    protestID: string;
  }>();

  // Получаем данные с сервера через контроллеры
  const { ratingInfo, loading: ratingLoading } = useGetRating();
  const { raceInfo, loading: raceLoading } = useGetRace();
  const { protestInfo, loading: protestLoading } = useGetProtest();

  const [isProtestModalOpen, setIsProtestModalOpen] = useState<boolean>(false);
  const [isCompleteModalOpen, setIsCompleteModalOpen] = useState<boolean>(false);

  // Восстановление состояний из localStorage при монтировании компонента
  useEffect(() => {
    const protestModalState = localStorage.getItem("isProtestModalOpen");
    const completeModalState = localStorage.getItem("isCompleteModalOpen");

    if (protestModalState) {
      setIsProtestModalOpen(JSON.parse(protestModalState));
    }
    if (completeModalState) {
      setIsCompleteModalOpen(JSON.parse(completeModalState));
    }
  }, []);

  // Сохранение состояний в localStorage при их изменении
  useEffect(() => {
    localStorage.setItem("isProtestModalOpen", JSON.stringify(isProtestModalOpen));
  }, [isProtestModalOpen]);

  useEffect(() => {
    localStorage.setItem("isCompleteModalOpen", JSON.stringify(isCompleteModalOpen));
  }, [isCompleteModalOpen]);

  const handleProtestModalOpen = () => setIsProtestModalOpen(true);
  const handleProtestModalClose = () => setIsProtestModalOpen(false);

  const handleCompleteModalOpen = () => setIsCompleteModalOpen(true);
  const handleCompleteModalClose = () => setIsCompleteModalOpen(false);

  // Если параметры не переданы
  if (!ratingID || !raceID || !protestID) {
    return <div>Неверные параметры URL!</div>;
  }

  // Загрузка данных
  if (ratingLoading || raceLoading || protestLoading) {
    return <div>Загрузка...</div>;
  }

  // Если данные не найдены
  if (!ratingInfo || !raceInfo || !protestInfo) {
    return <div>Информация не найдена!</div>;
  }

  const showCompleteButton = Number(protestInfo.Status) === 1; // Предположим, что 3 - это статус завершённого протеста

  return (
      <div style={{padding: "20px"}}>
        <h1 className={"headerH1"}>{`Рейтинг: ${ratingInfo.Name}`}</h1>
        <h2 className={"headerH2"}>{`Гонка №${raceInfo.number}`}</h2>
        <h3 className={"headerH3"}>{`Протест`}</h3>

        <div className="protest-info">
          <p>
            <strong>Номер правила:</strong> {protestInfo.RuleNum}
          </p>
          <p>
            <strong>Дата рассмотрения:</strong> {protestInfo.ReviewDate}
          </p>
          <p>
            <strong>Статус:</strong>{" "}
            {StatusMap[Number(protestInfo.Status)] || "Неизвестный статус"}
          </p>
          <p>
            <strong>Комментарий:</strong> {protestInfo.Comment}
          </p>
        </div>

        <div className="buttons-container">
          {showCompleteButton && (
              <button className="auth-required" onClick={handleCompleteModalOpen}>
                Завершить рассмотрение
              </button>
          )}
          <button className="auth-required" onClick={handleProtestModalOpen}>
            Обновить информацию
          </button>
        </div>

        {/* Модальное окно для завершения рассмотрения */}
        {isCompleteModalOpen && (
            <CompleteProtestModal
                isOpen={isCompleteModalOpen}
                onClose={handleCompleteModalClose}
            />
        )}

        {/* Таблица участников */}
        <ProtestCrewMembersTable
            ratingID={ratingID}
            raceID={raceID}
            protestID={protestID}
        />


        {/* Модальное окно для обновления информации */}
        {isProtestModalOpen && (
            <UpdateProtestModal
                protest={{
                  judgeId: protestInfo.JudgeID,
                  ruleNum: Number(protestInfo.RuleNum),
                  reviewDate: protestInfo.ReviewDate,
                  status: Number(protestInfo.Status),
                  comment: protestInfo.Comment,
                }}
                onClose={handleProtestModalClose}
            />
        )}
      </div>
  );
};

export default ProtestView;
