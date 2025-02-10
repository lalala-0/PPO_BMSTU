import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useGetCrewMembers } from "../../controllers/crewMemberControllers/getCrewMembersController";
import { ParticipantFormData } from "../../models/participantModel";
import { useDetachCrewMember } from "../../controllers/crewMemberControllers/detachCrewMemberController";
import ParticipantModal from "../participantViews/modalInputParticipant";

interface CrewMembersTableProps {
  crewID: string;
  ratingID: string;
}

const CrewMembersTable: React.FC<CrewMembersTableProps> = ({
  crewID,
  ratingID,
}) => {
  const navigate = useNavigate();
  const { detachCrewMember } = useDetachCrewMember();
  const {
    data: crewMembers,
    loading,
    error,
    getCrewMembers,
  } = useGetCrewMembers(ratingID, crewID);

  useEffect(() => {
    getCrewMembers(); // Загружаем участников при монтировании
  }, [ratingID, crewID]);

  const [filters, setFilters] = useState<Record<string, string>>({});
  const [selectedMember, setSelectedMember] =
    useState<ParticipantFormData | null>(null);
  const [modalType, setModalType] = useState<"update" | null>(null);

  const handleFilterChange = (key: string, value: string) => {
    setFilters((prev) => ({
      ...prev,
      [key]: value,
    }));
  };

  // Фильтрация аналогична RatingsTable
  const filteredMembers = (crewMembers || []).filter(
    (member) =>
      (filters.fio
        ? member.FIO.toLowerCase().includes(filters.fio.toLowerCase())
        : true) &&
      (filters.category
        ? member.Category.toLowerCase().includes(filters.category.toLowerCase())
        : true) &&
      (filters.gender
        ? member.Gender.toLowerCase().includes(filters.gender.toLowerCase())
        : true) &&
      (filters.birthday
        ? member.Birthday.toLowerCase().includes(filters.birthday.toLowerCase())
        : true) &&
      (filters.coach
        ? member.Coach.toLowerCase().includes(filters.coach.toLowerCase())
        : true),
  );

  const handleNavigate = (id: string) => {
    navigate(`/participants/${id}`);
  };

  const onUpdate = (member: ParticipantFormData) => {
    setSelectedMember(member);
    setModalType("update");
  };

  const onDelete = async (id: string) => {
    if (
      window.confirm("Вы уверены, что хотите удалить участника из команды?")
    ) {
      try {
        await detachCrewMember(id);
        getCrewMembers(); // Обновляем список после удаления
      } catch (err) {
        alert("Ошибка при удалении участника");
      }
    }
  };

  const handleModalClose = () => {
    setSelectedMember(null);
    setModalType(null);
    getCrewMembers(); // Обновляем список после закрытия
  };

  if (loading) return <div>Загрузка...</div>;
  if (error) return <div>Ошибка при загрузке участников</div>;

  return (
    <div className="crew-table-container">
      <table className="crew-table">
        <thead>
          <tr>
            {[
              { key: "fio", label: "ФИО" },
              { key: "category", label: "Категория" },
              { key: "gender", label: "Пол" },
              { key: "birthday", label: "Дата рождения" },
              { key: "coach", label: "Тренер" },
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
          {filteredMembers.map((member) => (
            <tr key={member.id}>
              <td>
                <button
                  onClick={() => handleNavigate(member.id)}
                  className="link-button"
                >
                  {member.FIO}
                </button>
              </td>
              <td>{member.Category}</td>
              <td>{member.Gender}</td>
              <td>{member.Birthday}</td>
              <td>{member.Coach}</td>
              <td className="auth-required">
                <div className="buttons-container">
                  <button onClick={() => onDelete(member.id)}>
                    <img
                      src="/delete-icon.svg"
                      alt="Удалить"
                      width="20"
                      height="20"
                    />
                  </button>
                  <button onClick={() => onUpdate(member)}>
                    <img
                      src="/update-icon.svg"
                      alt="Обновить"
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

      {modalType === "update" && selectedMember && (
        <ParticipantModal
          participant={selectedMember}
          type="update"
          onClose={handleModalClose}
        />
      )}
    </div>
  );
};

export default CrewMembersTable;
