import React, { useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import ProtestsTable from "./tableProtests"; // Таблица протестов
import RaceModal from "./modalInputRace"; // Модальное окно для обновления информации о гонке
import { useGetRating } from "../../controllers/ratingControllers/getRatingController";
import { useGetRace } from "../../controllers/raceControllers/getRaceController";


const RaceView: React.FC = () => {
  const { ratingID, raceID } = useParams<{
    ratingID: string;
    raceID: string;
  }>();
  const navigate = useNavigate();

  const { ratingInfo, loading: ratingLoading } = useGetRating();
  const { raceInfo, loading: raceLoading } = useGetRace();

  const [isRaceModalOpen, setIsRaceModalOpen] = useState(false);

  const handleRaceModalOpen = () => setIsRaceModalOpen(true);
  const handleRaceModalClose = () => setIsRaceModalOpen(false);

  // Проверка на наличие параметров URL и возвращение ошибки, если они отсутствуют
  if (!ratingID || !raceID) {
    return <div>Неверные параметры URL!</div>;
  }

  // Загрузка данных
  if (ratingLoading || raceLoading) {
    return <div>Загрузка...</div>;
  }

  if (!ratingInfo || !raceInfo) {
    return <div>Информация не найдена!</div>;
  }

  // Обработчики кнопок
  const handleStartProcedure = () => {
    navigate(`/ratings/${ratingID}/races/${raceID}/startProcedure`);
  };

  const handleFinishProcedure = () => {
    navigate(`/ratings/${ratingID}/races/${raceID}/finishProcedure`);
  };

  const handleCreateProtest = () => {
    // Здесь может быть логика для открытия модального окна или перенаправления на страницу создания протеста
    alert("Функционал создания протеста еще не реализован");
  };

  return (
    <div style={{ padding: "20px" }}>
      <h1>{`Рейтинг: ${ratingInfo.Name}`}</h1>
      <h2>{`Гонка №${raceInfo.number} (${raceInfo.class})`}</h2>
      <ProtestsTable /> {/* Таблица протестов */}
      <button className="auth-required" onClick={handleCreateProtest}>
        Создать протест
      </button>
      <div className="buttons-container">
        <button className="auth-required" onClick={handleStartProcedure}>
          Начать стартовую процедуру
        </button>
        <button className="auth-required" onClick={handleFinishProcedure}>
          Начать финишную процедуру
        </button>
      </div>
      <button className="auth-required" onClick={handleRaceModalOpen}>
        Обновить информацию о гонке
      </button>
      {isRaceModalOpen && (
        <RaceModal
          race={raceInfo}
          type={"update"}
          onClose={handleRaceModalClose}
        />
      )}
    </div>
  );
};

export default RaceView;
