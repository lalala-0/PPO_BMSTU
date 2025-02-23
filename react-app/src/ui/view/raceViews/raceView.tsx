import React, { useState, useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";
import ProtestsTable from "./tableProtests"; // Таблица протестов
import RaceModal from "./modalInputRace"; // Модальное окно для обновления информации о гонке
import { useGetRating } from "../../controllers/ratingControllers/getRatingController";
import { useGetRace } from "../../controllers/raceControllers/getRaceController";
import CreateProtestModal from "../protestViews/modalCreateProtest";

const RaceView: React.FC = () => {
    const { ratingID, raceID } = useParams<{ ratingID: string; raceID: string }>();
    const navigate = useNavigate();

    const { ratingInfo, loading: ratingLoading } = useGetRating();
    const { raceInfo, loading: raceLoading } = useGetRace();

    // Загружаем сохраненные состояния для всех модальных окон, если они есть
    const [isRaceModalOpen, setIsRaceModalOpen] = useState(() => {
        const saved = localStorage.getItem("isRaceModalOpen");
        return saved ? JSON.parse(saved) : false;
    });

    const [isProtestModalOpen, setIsProtestModalOpen] = useState(() => {
        const saved = localStorage.getItem("isProtestModalOpen");
        return saved ? JSON.parse(saved) : false;
    });

    // При изменении состояния модальных окон сохраняем в localStorage
    useEffect(() => {
        localStorage.setItem("isRaceModalOpen", JSON.stringify(isRaceModalOpen));
    }, [isRaceModalOpen]);

    useEffect(() => {
        localStorage.setItem("isProtestModalOpen", JSON.stringify(isProtestModalOpen));
    }, [isProtestModalOpen]);

    const handleRaceModalOpen = () => setIsRaceModalOpen(true);
    const handleRaceModalClose = () => setIsRaceModalOpen(false);

    const handleCreateProtest = () => setIsProtestModalOpen(true);
    const handleProtestModalClose = () => setIsProtestModalOpen(false);

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

    return (
        <div style={{padding: "20px"}}>
            <h1 className={"headerH1"}>{`Рейтинг: ${ratingInfo.Name} (${raceInfo.class})`}</h1>
            <h2 className={"headerH2"}>{`Гонка №${raceInfo.number} (${raceInfo.date})`}</h2>
            <h3 className={"headerH3"}>{`Список протестов`}</h3>
            <ProtestsTable/> {/* Таблица протестов */}
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
                <RaceModal race={raceInfo} type={"update"} onClose={handleRaceModalClose}/>
            )}
            {isProtestModalOpen && <CreateProtestModal onClose={handleProtestModalClose}/>}
        </div>
    );
};

export default RaceView;
