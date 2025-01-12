import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useGetCrewMembers } from "../../controllers/crewMemberControllers/getCrewMembersController";
import { ParticipantFormData } from "../../models/participantModel";
import { useDetachCrewMember } from "../../controllers/crewMemberControllers/detachCrewMemberController";

interface CrewMembersTableProps {
  crewID: string;
  ratingID: string;
}

const CrewMembersTable: React.FC<CrewMembersTableProps> = ({
  crewID,
  ratingID,
}) => {
  const navigate = useNavigate();
  const {
    detachCrewMember,
    success,
    loading: detachLoading,
    error: detachError,
  } = useDetachCrewMember();

  // Состояние фильтров
  const [filters, setFilters] = useState<Record<string, string>>({});
  const [selectedMember, setSelectedMember] =
    useState<ParticipantFormData | null>(null);
  const [modalType, setModalType] = useState<"update" | null>(null);

  const { data, loading, error } = useGetCrewMembers(ratingID, crewID);

  // Обработчик изменения фильтров
  const handleFilterChange = (key: string, value: string) => {
    setFilters((prev) => ({
      ...prev,
      [key]: value,
    }));
  };

  // Универсальная фильтрация
  const filteredMembers: ParticipantFormData[] =
    data?.filter((member: ParticipantFormData) => {
      return Object.entries(filters).every(([key, value]) => {
        if (!value) return true; // Если фильтр пустой, пропускаем
        const memberValue = (member as Record<string, any>)[key]; // Доступ к значению по ключу
        return memberValue
          ?.toString()
          .toLowerCase()
          .includes(value.toLowerCase());
      });
    }) || [];

  const handleNavigate = (id: string) => {
    navigate(`/participants/${id}`);
  };

  const onUpdate = (member: ParticipantFormData) => {
    setSelectedMember(member);
    setModalType("update");
  };

  const onDelete = async (id: string) => {
    if (window.confirm("Вы уверены, что хотите удалить участника?")) {
      try {
        await detachCrewMember(id);
      } catch (err) {
        alert("Ошибка при удалении участника");
      }
    }
  };

  const handleModalClose = () => {
    setSelectedMember(null);
    setModalType(null);
  };

  if (loading) {
    return <div>Загрузка...</div>;
  }

  if (error) {
    return <div>Ошибка при загрузке участников</div>;
  }

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
            <th>Действия</th>
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
                  {member.fio}
                </button>
              </td>
              <td>{member.category}</td>
              <td>{member.gender}</td>
              <td>{member.birthday}</td>
              <td>{member.coach}</td>
              <td>
                <div className="buttons-container">
                  <button
                    className="delete-button"
                    onClick={() => onDelete(member.id)}
                  >
                    <img
                      src="/delete-icon.svg"
                      alt="Удалить"
                      width="20"
                      height="20"
                    />
                  </button>
                  <button
                    className="update-button"
                    onClick={() => onUpdate(member)}
                  >
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
    </div>
  );
};

export default CrewMembersTable;
