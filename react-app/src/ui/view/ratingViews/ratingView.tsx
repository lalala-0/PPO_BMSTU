import React, { useCallback, useMemo } from "react";
import { useParams } from "react-router-dom";
import RatingTable from "./rankingView";
import RatingModal from "./modalInputRating";
import CrewModal from "../crewViews/modalInputCrew";
import RaceModal from "../raceViews/modalInputRace";
import { useGetRating } from "../../controllers/ratingControllers/getRatingController";

const RatingView: React.FC = () => {
  const { ratingInfo, rankingTable, races, loading } = useGetRating();
  const ratingID = useParams().ratingID;
  const currentRating = useMemo(() => ratingInfo, [ratingInfo]);

  const [isRatingModalOpen, setIsRatingModalOpen] = React.useState(false);
  const [isCrewModalOpen, setIsCrewModalOpen] = React.useState(false);
  const [isRaceModalOpen, setIsRaceModalOpen] = React.useState(false);

  const handleRatingModalOpen = useCallback(() => setIsRatingModalOpen(true), []);
  const handleRatingModalClose = useCallback(() => setIsRatingModalOpen(false), []);
  const handleCrewModalOpen = useCallback(() => setIsCrewModalOpen(true), []);
  const handleCrewModalClose = useCallback(() => setIsCrewModalOpen(false), []);
  const handleRaceModalOpen = useCallback(() => setIsRaceModalOpen(true), []);
  const handleRaceModalClose = useCallback(() => setIsRaceModalOpen(false), []);

  if (loading) return <div>Загрузка...</div>;
  if (!ratingInfo) return <div>Рейтинг не найден!</div>;

  return (
      <div className="rating-container">
        <h1>{`Рейтинг: ${ratingInfo.Name}`}</h1>
        <RatingTable rankingTable={rankingTable} races={races} />
        <div className="buttons-container">
          <button className="auth-required" onClick={handleCrewModalOpen}>Создать команду</button>
          <button className="auth-required" onClick={handleRaceModalOpen}>Создать гонку</button>
        </div>
        <button className="auth-required" onClick={handleRatingModalOpen}>Обновить рейтинг</button>

        {/* Модальные окна вынесены в отдельную переменную */}
        {isRatingModalOpen && currentRating && (
            <RatingModal rating={currentRating} type="update" onClose={handleRatingModalClose} />
        )}
        {isCrewModalOpen && ratingID && (
            <CrewModal ratingID={ratingID} crew={{ id: "", ratingId: ratingID, SailNum: 1, Class: "" }} type="create" onClose={handleCrewModalClose} />
        )}
        {isRaceModalOpen && ratingID && (
            <RaceModal race={{ id: "", ratingId: ratingID, date: "", number: 1, class: "" }} type="create" onClose={handleRaceModalClose} />
        )}
      </div>
  );
};

export default RatingView;
