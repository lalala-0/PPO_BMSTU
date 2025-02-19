import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { Rating } from "../../models/ratingModel";
import RatingTable from "./rankingView";
import RatingModal from "./modalInputRating";
import CrewModal from "../crewViews/modalInputCrew";
import RaceModal from "../raceViews/modalInputRace";
import { useGetRating } from "../../controllers/ratingControllers/getRatingController";


const RatingView: React.FC = () => {
  const { ratingInfo, rankingTable, races, loading } = useGetRating();
  const [currentRating, setCurrentRating] = useState<Rating | null>(null);
  const [isRatingModalOpen, setIsRatingModalOpen] = useState(false);
  const [isCrewModalOpen, setIsCrewModalOpen] = useState(false);
  const [isRaceModalOpen, setIsRaceModalOpen] = useState(false);
  const { ratingID } = useParams(); // Получаем ratingID из URL, оно уже строка

  // Обработчики открытия и закрытия модальных окон
  const handleRatingModalOpen = () => setIsRatingModalOpen(true);
  const handleRatingModalClose = () => setIsRatingModalOpen(false);

  const handleCrewModalOpen = () => setIsCrewModalOpen(true);
  const handleCrewModalClose = () => setIsCrewModalOpen(false);

  const handleRaceModalOpen = () => setIsRaceModalOpen(true);
  const handleRaceModalClose = () => setIsRaceModalOpen(false);

  useEffect(() => {
    if (ratingInfo) {
      setCurrentRating(ratingInfo);
    }
  }, [ratingInfo]);

  if (loading) {
    return <div>Загрузка...</div>;
  }

  if (!ratingInfo) {
    return <div>Рейтинг не найден!</div>;
  }

  return (
    <div style={{ padding: "20px" }}>
      <h1>{`Рейтинг: ${ratingInfo.Name}`}</h1>
      <RatingTable rankingTable={rankingTable} races={races} />
      <div className="buttons-container">
        <button className="auth-required" onClick={handleCrewModalOpen}>
          Создать команду
        </button>
        <button className="auth-required" onClick={handleRaceModalOpen}>
          Создать гонку
        </button>
      </div>
      <button className="auth-required" onClick={handleRatingModalOpen}>
        Обновить рейтинг
      </button>

      {isRatingModalOpen && currentRating && (
        <RatingModal
          rating={currentRating}
          type={"update"} // Передаем тип действия в RatingModal
          onClose={handleRatingModalClose}
        />
      )}

      {isCrewModalOpen && ratingID && (
        <CrewModal
          ratingID={ratingID}
          crew={{ id: "", ratingId: ratingID, SailNum: 1, Class: "" }} // Передаем ratingID как строку
          type={"create"} // Передаем тип действия в CrewModal
          onClose={handleCrewModalClose}
        />
      )}

      {isRaceModalOpen && ratingID && (
        <RaceModal
          race={{ id: "", ratingId: ratingID, date: "", number: 1, class: "" }} // Передаем ratingID как строку
          type={"create"}
          onClose={handleRaceModalClose}
        />
      )}
    </div>
  );
};

export default RatingView;
