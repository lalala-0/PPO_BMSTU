import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { useGetParticipant } from "../../controllers/participantControllers/getParticipantController";
import ParticipantModal from "./modalInputParticipant";
import "../../styles/styles.css";

const ParticipantView: React.FC = () => {
  const { participantID } = useParams<{ participantID: string }>();
  const id = participantID || "undefined";
  // Хуки для получения данных о рейтинге, команде и участнике
  const {
    getParticipant,
    participant,
    loading: participantLoading,
  } = useGetParticipant(id);

  const [isParticipantModalOpen, setIsParticipantModalOpen] = useState(false);

  const handleParticipantModalOpen = () => setIsParticipantModalOpen(true);
  const handleParticipantModalClose = () => setIsParticipantModalOpen(false);

  // Вызов getParticipant при загрузке компонента
  useEffect(() => {
    if (participantID) {
      getParticipant();
    }
  }, [participantID, getParticipant]);

  // Проверка на наличие параметров URL и возвращение ошибки, если они отсутствуют
  if (!participantID) {
    return <div>Неверные параметры URL!</div>;
  }

  // Проверка загрузки данных
  if (participantLoading) {
    return <div>Загрузка...</div>;
  }

  if (!participant) {
    return <div>Информация не найдена!</div>;
  }

  return (
    <div style={{ padding: "20px" }}>
      <h3>{`Участник: ${participant.fio}`}</h3>

      <div>
        <p>
          <strong>ФИО:</strong> {participant.fio}
        </p>
        <p>
          <strong>Категория:</strong> {participant.category}
        </p>
        <p>
          <strong>Пол:</strong> {participant.gender}
        </p>
        <p>
          <strong>Дата рождения:</strong> {participant.birthday}
        </p>
        <p>
          <strong>Тренер:</strong> {participant.coach}
        </p>
      </div>

      <div className="buttons-container">
        <button onClick={handleParticipantModalOpen}>
          Изменить информацию об участнике
        </button>
      </div>

      {isParticipantModalOpen && participantID && (
        <ParticipantModal
          participant={participant}
          type={"update"}
          onClose={handleParticipantModalClose}
        />
      )}
    </div>
  );
};

export default ParticipantView;