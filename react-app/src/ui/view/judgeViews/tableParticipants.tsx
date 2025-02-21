import React, { useState, useEffect } from "react";
import { ParticipantFormData } from "../../models/participantModel";
import { useGetAllParticipants } from "../../controllers/participantControllers/getParticipantsController";
import ParticipantModal from "../participantViews/modalInputParticipant";
import { useDeleteParticipant } from "../../controllers/participantControllers/deleteParticipantController"; // Импортируем хук

const ParticipantTable: React.FC = () => {
  const [showParticipantModal, setShowParticipantModal] = useState(false);
  const [selectedParticipant, setSelectedParticipant] =
    useState<ParticipantFormData | null>(null);

  const {
    getParticipants,
    loading: loadingParticipants,
    error: participantsError,
    participants,
  } = useGetAllParticipants({});

  const {
    deleteParticipant,
  } = useDeleteParticipant(); // Хук для удаления

  // Фильтры
  const [fioFilter, setFioFilter] = useState("");
  const [categoryFilter, setCategoryFilter] = useState("");
  const [genderFilter, setGenderFilter] = useState("");
  const [birthdayFilter, setBirthdayFilter] = useState("");
  const [coachFilter, setCoachFilter] = useState("");

  const handleAddParticipantClick = () => {
    setSelectedParticipant(null);
    setShowParticipantModal(true);
  };

  const handleParticipantModalClose = () => {
    setShowParticipantModal(false);
  };

  useEffect(() => {
    if (participantsError) {
      alert(`Ошибка загрузки участников: ${participantsError}`);
    }
  }, [participantsError]);

  useEffect(() => {
    if (participants.length === 0 && !loadingParticipants) {
      getParticipants();
    }
  }, [participants, getParticipants, loadingParticipants]);

  const handleDeleteParticipant = (id: string) => {
    if (
      window.confirm("Вы уверены, что хотите удалить участника из команды?")
    ) {
      try {
        deleteParticipant(id);
        getParticipants(); // Обновляем список после удаления
      } catch {
        alert("Ошибка при удалении участника");
      }
    }
  };

  // Фильтрация участников
  const filteredParticipants = participants.filter((participant) => {
    return (
      (fioFilter === "" ||
        participant.FIO.toLowerCase().includes(fioFilter.toLowerCase())) &&
      (categoryFilter === "" ||
        participant.Category.toLowerCase().includes(
          categoryFilter.toLowerCase(),
        )) &&
      (genderFilter === "" ||
        participant.Gender.toLowerCase().includes(
          genderFilter.toLowerCase(),
        )) &&
      (birthdayFilter === "" ||
        participant.Birthday.toLowerCase().includes(
          birthdayFilter.toLowerCase(),
        )) &&
      (coachFilter === "" ||
        participant.Coach.toLowerCase().includes(coachFilter.toLowerCase()))
    );
  });

  if (loadingParticipants) {
    return <div>Загрузка...</div>;
  }

  return (
    <div>
      {/* Таблица участников */}
      <h2>Участники</h2>
      <div className={"tableContent"}>
        <table >
          <thead>
            <tr>
              <th className={"stickyHeader"}>ФИО</th>
              <th className={"stickyHeader"}>Категория</th>
              <th className={"stickyHeader"}>Пол</th>
              <th className={"stickyHeader"}>Дата рождения</th>
              <th className={"stickyHeader"}>Тренер</th>
              <th className={"stickyHeader"}>Действия</th>
            </tr>
            <tr>
              <th>
                <input
                  type="text"
                  placeholder="ФИО"
                  value={fioFilter}
                  onChange={(e) => setFioFilter(e.target.value)}
                />
              </th>
              <th>
                <input
                  type="text"
                  placeholder="Категория"
                  value={categoryFilter}
                  onChange={(e) => setCategoryFilter(e.target.value)}
                />
              </th>
              <th>
                <input
                  type="text"
                  placeholder="Пол"
                  value={genderFilter}
                  onChange={(e) => setGenderFilter(e.target.value)}
                />
              </th>
              <th>
                <input
                  type="text"
                  placeholder="Дата рождения"
                  value={birthdayFilter}
                  onChange={(e) => setBirthdayFilter(e.target.value)}
                />
              </th>
              <th>
                <input
                  type="text"
                  placeholder="Тренер"
                  value={coachFilter}
                  onChange={(e) => setCoachFilter(e.target.value)}
                />
              </th>
              <th></th>
            </tr>
          </thead>
          <tbody>
            {filteredParticipants.length > 0 ? (
              filteredParticipants.map((participant) => (
                <tr key={participant.id}>
                  <td>{participant.FIO}</td>
                  <td>{participant.Category}</td>
                  <td>{participant.Gender}</td>
                  <td>{participant.Birthday}</td>
                  <td>{participant.Coach}</td>
                  <td>
                    <button
                      onClick={() => {
                        setSelectedParticipant(participant);
                        setShowParticipantModal(true);
                      }}
                    >
                      Редактировать
                    </button>
                    <button
                      onClick={() => handleDeleteParticipant(participant.id)}
                    >
                      Удалить
                    </button>
                  </td>
                </tr>
              ))
            ) : (
              <tr>
                <td colSpan={6}>Нет участников</td>
              </tr>
            )}
          </tbody>
        </table>
      </div>
      <button onClick={handleAddParticipantClick}>Добавить участника</button>

      {/* Модалки для добавления/редактирования */}
      {showParticipantModal && (
        <ParticipantModal
          participant={{
            id: selectedParticipant ? selectedParticipant.id : "",
            FIO: selectedParticipant ? selectedParticipant.FIO : "",
            Category: selectedParticipant ? selectedParticipant.Category : "",
            Gender: selectedParticipant ? selectedParticipant.Gender : "",
            Birthday: selectedParticipant ? selectedParticipant.Birthday : "",
            Coach: selectedParticipant ? selectedParticipant.Coach : "",
          }}
          type={selectedParticipant ? "update" : "create"}
          onClose={handleParticipantModalClose}
        />
      )}
    </div>
  );
};

export default ParticipantTable;
