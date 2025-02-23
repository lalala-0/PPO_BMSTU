import React, { useCallback, useMemo, useEffect } from "react";
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

    // Восстановление состояния модалок из localStorage
    useEffect(() => {
        const storedModalState = localStorage.getItem("ratingViewModalState");
        if (storedModalState) {
            const parsedState = JSON.parse(storedModalState);
            setIsRatingModalOpen(parsedState.isRatingModalOpen || false);
            setIsCrewModalOpen(parsedState.isCrewModalOpen || false);
            setIsRaceModalOpen(parsedState.isRaceModalOpen || false);
        }
    }, []);

    const handleRatingModalOpen = useCallback(() => {
        setIsRatingModalOpen(true);
        // Сохраняем состояние в localStorage
        localStorage.setItem("ratingViewModalState", JSON.stringify({
            isRatingModalOpen: true,
            isCrewModalOpen,
            isRaceModalOpen,
        }));
    }, [isCrewModalOpen, isRaceModalOpen]);

    const handleRatingModalClose = useCallback(() => {
        setIsRatingModalOpen(false);
        // Сохраняем состояние в localStorage
        localStorage.setItem("ratingViewModalState", JSON.stringify({
            isRatingModalOpen: false,
            isCrewModalOpen,
            isRaceModalOpen,
        }));
    }, [isCrewModalOpen, isRaceModalOpen]);

    const handleCrewModalOpen = useCallback(() => {
        setIsCrewModalOpen(true);
        // Сохраняем состояние в localStorage
        localStorage.setItem("ratingViewModalState", JSON.stringify({
            isRatingModalOpen,
            isCrewModalOpen: true,
            isRaceModalOpen,
        }));
    }, [isRatingModalOpen, isRaceModalOpen]);

    const handleCrewModalClose = useCallback(() => {
        setIsCrewModalOpen(false);
        // Сохраняем состояние в localStorage
        localStorage.setItem("ratingViewModalState", JSON.stringify({
            isRatingModalOpen,
            isCrewModalOpen: false,
            isRaceModalOpen,
        }));
    }, [isRatingModalOpen, isRaceModalOpen]);

    const handleRaceModalOpen = useCallback(() => {
        setIsRaceModalOpen(true);
        // Сохраняем состояние в localStorage
        localStorage.setItem("ratingViewModalState", JSON.stringify({
            isRatingModalOpen,
            isCrewModalOpen,
            isRaceModalOpen: true,
        }));
    }, [isRatingModalOpen, isCrewModalOpen]);

    const handleRaceModalClose = useCallback(() => {
        setIsRaceModalOpen(false);
        // Сохраняем состояние в localStorage
        localStorage.setItem("ratingViewModalState", JSON.stringify({
            isRatingModalOpen,
            isCrewModalOpen,
            isRaceModalOpen: false,
        }));
    }, [isRatingModalOpen, isCrewModalOpen]);

    if (loading) return <div>Загрузка...</div>;
    if (!ratingInfo) return <div>Рейтинг не найден!</div>;

    return (
        <div className="rating-container">
            <h1 className={"headerH1"}>{`Рейтинг: ${ratingInfo.Name}`}</h1>
            <RatingTable rankingTable={rankingTable} races={races} />
            <div className="buttons-container">
                <button className="auth-required" onClick={handleCrewModalOpen}>Создать команду</button>
                <button className="auth-required" onClick={handleRaceModalOpen}>Создать гонку</button>
            </div>
            <button className="auth-required" onClick={handleRatingModalOpen}>Обновить рейтинг</button>

            {/* Модальные окна */}
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
