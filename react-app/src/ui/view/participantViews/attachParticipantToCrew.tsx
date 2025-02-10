import React, { useEffect, useState } from "react";
import { useAttachCrewMember } from "../../controllers/crewMemberControllers/attachCrewMemberController";
import { ParticipantFormData } from "../../models/participantModel";
import {
  HelmsmanMap,
  CrewParticipantAttachInput,
} from "../../models/crewModel";
import { useGetAllParticipants } from "../../controllers/participantControllers/getParticipantsController";

interface AttachCrewMemberModalProps {
  ratingID: string;
  crewID: string;
  onClose: () => void;
  onSuccess: () => void; // Функция для обновления списка после добавления
}

const AttachCrewMemberModal: React.FC<AttachCrewMemberModalProps> = ({
  ratingID,
  crewID,
  onClose,
  onSuccess,
}) => {
  const { getParticipants, participants, loading, error } =
    useGetAllParticipants({});
  const {
    attachCrewMember,
    success,
    loading: attaching,
    error: attachError,
  } = useAttachCrewMember(ratingID, crewID);

  const [selectedParticipantID, setSelectedParticipantID] =
    useState<string>("");
  const [helmsman, setHelmsman] = useState<number>(0);

  useEffect(() => {
    getParticipants(); // Загружаем всех участников при открытии модального окна
  }, []);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!selectedParticipantID) {
      alert("Выберите участника!");
      return;
    }

    const data: CrewParticipantAttachInput = {
      participantID: selectedParticipantID,
      helmsman,
    };

    await attachCrewMember(data);
    if (success) {
      onSuccess(); // Обновляем список в родительском компоненте
      onClose(); // Закрываем модальное окно
    }
  };

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal-content" onClick={(e) => e.stopPropagation()}>
        <h2>Добавить участника в команду</h2>

        {loading ? (
          <p>Загрузка участников...</p>
        ) : error ? (
          <p className="error">Ошибка при загрузке участников</p>
        ) : (
          <form onSubmit={handleSubmit}>
            <label>
              Участник:
              <select
                value={selectedParticipantID}
                onChange={(e) => setSelectedParticipantID(e.target.value)}
                required
              >
                <option value="" disabled>
                  Выберите участника
                </option>
                {participants.map((participant: ParticipantFormData) => (
                  <option key={participant.id} value={participant.id}>
                    {participant.FIO}
                  </option>
                ))}
              </select>
            </label>

            <label>
              Роль:
              <select
                value={helmsman}
                onChange={(e) => setHelmsman(Number(e.target.value))}
              >
                {Object.entries(HelmsmanMap).map(([key, value]) => (
                  <option key={key} value={key}>
                    {value}
                  </option>
                ))}
              </select>
            </label>

            <button type="submit" disabled={attaching}>
              {attaching ? "Добавление..." : "Добавить"}
            </button>
          </form>
        )}

        {attachError && <p className="error">Ошибка: {attachError}</p>}
        {success && <p className="success">{success}</p>}

        <button className="close-button" onClick={onClose}>
          Закрыть
        </button>
      </div>
    </div>
  );
};

export default AttachCrewMemberModal;
