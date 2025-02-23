import React, { useState, useEffect } from "react";
import { ProtestParticipantAttachInput, RoleMap } from "../../models/protestModel";
import { useAttachProtestMember } from "../../controllers/protestControllers/attachProtestMemberController";
import { useGetCrewsByRatingID } from "../../controllers/crewControllers/getCrewsController";

interface AddProtestMemberModalProps {
    ratingID: string;
    raceID: string;
    protestID: string;
    onClose: () => void;
}

const AddProtestMemberModal: React.FC<AddProtestMemberModalProps> = ({
                                                                         ratingID,
                                                                         raceID,
                                                                         protestID,
                                                                         onClose,
                                                                     }) => {
    const [sailNum, setSailNum] = useState<number | string>(""); // состояние для номера паруса
    const [role, setRole] = useState<number | string>("");

    const { attachProtestMember, successMessage, loading, error } = useAttachProtestMember(
        ratingID,
        raceID,
        protestID
    );

    const { crews, loading: crewsLoading, error: crewsError, getCrewsByRatingID } = useGetCrewsByRatingID();

    // Загружаем участников по рейтингу при открытии модалки
    useEffect(() => {
        if (ratingID) {
            getCrewsByRatingID(ratingID);
        }
    }, [ratingID, getCrewsByRatingID]);

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();

        if (sailNum && role) {
            const protestParticipantInput: ProtestParticipantAttachInput = {
                sailNum: Number(sailNum),
                role: Number(role),
            };

            await attachProtestMember(protestParticipantInput);
             onClose();  // Закрытие модалки после успешного добавления
        }
    };

    return (
        <div className="modal-overlay" onClick={onClose}>
            <div className="modal-content" onClick={(e) => e.stopPropagation()}>
                <h2>Добавить участника в протест</h2>

                {successMessage && <div className="success-message">{successMessage}</div>}
                {error && <div className="error-message">{error}</div>}
                {crewsError && <div className="error-message">{crewsError}</div>}

                <form onSubmit={handleSubmit}>
                    <label>
                        Номер парусника:
                        <select
                            value={sailNum}
                            onChange={(e) => setSailNum(e.target.value)}
                            required
                        >
                            <option value="" disabled>Выберите номер парусника</option>
                            {crewsLoading ? (
                                <option>Загрузка...</option>
                            ) : (
                                crews?.map((crew) => (
                                    <option key={crew.SailNum} value={crew.SailNum}>
                                        {crew.SailNum}
                                    </option>
                                ))
                            )}
                        </select>
                    </label>

                    <label>
                        Роль:
                        <select
                            value={role}
                            onChange={(e) => setRole(e.target.value)}
                            required
                        >
                            <option value="" disabled>Выберите роль</option>
                            {Object.entries(RoleMap).map(([key, value]) => (
                                <option key={key} value={key}>
                                    {value}
                                </option>
                            ))}
                        </select>
                    </label>

                    <div className="modal-actions">
                        <button type="submit" disabled={loading}>
                            {loading ? "Загрузка..." : "Добавить участника"}
                        </button>
                        <button type="button" onClick={onClose}>
                            Закрыть
                        </button>
                    </div>
                </form>
            </div>
        </div>
    );
};

export default AddProtestMemberModal;
