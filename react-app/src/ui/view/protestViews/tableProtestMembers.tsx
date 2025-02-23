import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useGetProtestMembers } from "../../controllers/protestControllers/getProtestMembersController";
import { useDetachProtestMember } from "../../controllers/protestControllers/detachProtestMemberController";
import AddProtestMemberModal from "./modalAttachMember";

interface ProtestCrewMembersTableProps {
    ratingID: string;
    raceID: string;
    protestID: string;
}

const ProtestCrewMembersTable: React.FC<ProtestCrewMembersTableProps> = ({
                                                                             ratingID,
                                                                             raceID,
                                                                             protestID,
                                                                         }) => {
    const navigate = useNavigate();
    const { detachProtestMember } = useDetachProtestMember(ratingID, raceID, protestID);
    const { protestMembers, loading, error } = useGetProtestMembers(ratingID, raceID, protestID);

    const [filters, setFilters] = useState<Record<string, string>>({});
    const [modalType, setModalType] = useState<"add" | "update" | null>(null);

    const handleFilterChange = (key: string, value: string) => {
        setFilters((prev) => ({
            ...prev,
            [key]: value,
        }));
    };

    // Фильтрация участников протеста
    const filteredMembers = (protestMembers || []).filter(
        (member) =>
            (filters.sailNum ? member.sailNum.toString().includes(filters.sailNum) : true) &&
            (filters.role ? member.role.toString().includes(filters.role) : true)
    );

    const handleNavigate = (ratingID: string, crewID: string) => {
        navigate(`/ratings/${ratingID}/crews/${crewID}`);
    };

    const onDelete = async (ratingID: string, sailNum: number) => {
        if (window.confirm("Вы уверены, что хотите удалить участника из протеста?")) {
            try {
                await detachProtestMember({ ratingID, sailNum });
            } catch {
                alert("Ошибка при удалении участника");
            }
        }
    };

    const handleCloseModal = () => {
        setModalType(null);
        localStorage.setItem("modalState", "");
    };

    const handleModalOpen = () => {
        setModalType("add"); // Открытие модалки для добавления
        localStorage.setItem("modalState", "open");
    };

    // Проверка состояния модалки в localStorage не обязательно, но можно оставить если нужно
    useEffect(() => {
        const modalState = localStorage.getItem("modalState");
        if (modalState === "open") {
            setModalType("add"); // Если модалка была открыта при последнем визите
        }
    }, []);

    if (loading) return <div>Загрузка...</div>;
    if (error) return <div>Ошибка при загрузке участников</div>;

    return (
        <div>
            <div className="tableContent">
                <table className="table">
                    <thead>
                    <tr>
                        {[{ key: "sailNum", label: "Номер паруса" }, { key: "role", label: "Роль" }].map(({ key, label }) => (
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
                    {filteredMembers.map((member) => (
                        <tr key={member.sailNum}>
                            <td>
                                {/* Кнопка с номером паруса */}
                                <button onClick={() => handleNavigate(member.ratingID, member.id)}>
                                    {member.sailNum}
                                </button>
                            </td>
                            <td>{member.role}</td>
                            <td className="auth-required">
                                <div className="buttons-container">
                                    <button onClick={() => onDelete(member.ratingID, member.sailNum)}>
                                        <img src="/delete-icon.svg" alt="Удалить" width="20" height="20" />
                                    </button>
                                </div>
                            </td>
                        </tr>
                    ))}
                    </tbody>
                </table>
            </div>

            {/* Кнопка для добавления участника */}
            <button onClick={handleModalOpen}>
                Добавить участника
            </button>

            {/* Показываем модалку, если она должна быть открыта */}
            {modalType === "add" && (
                <AddProtestMemberModal
                    ratingID={ratingID}
                    raceID={raceID}
                    protestID={protestID}
                    onClose={handleCloseModal}
                />
            )}
        </div>
    );
};

export default ProtestCrewMembersTable;
