import React, { useState } from "react";
import {
  ParticipantFormData,
  ParticipantInput,
  CategoryMap,
  GenderMap,
} from "../../models/participantModel";
import { useCreateParticipant } from "../../controllers/participantControllers/createParticipantController";
import { useUpdateParticipant } from "../../controllers/participantControllers/updateParticipantController";

interface ParticipantModalProps {
  participant: ParticipantFormData;
  type: "create" | "update";
  onClose: () => void;
}

const ParticipantModal: React.FC<ParticipantModalProps> = ({
  participant,
  type,
  onClose,
}) => {
  const [formData, setFormData] = useState<ParticipantInput>({
    fio: participant.FIO,
    category: participant.Category ? parseInt(participant.Category) : 1,
    gender: participant.Gender ? parseInt(participant.Gender) : 1,
    birthday: participant.Birthday,
    coach: participant.Coach,
  });
  const { createParticipant } = useCreateParticipant();
  const { updateParticipant } = useUpdateParticipant(
    participant.id,
  );

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>,
  ) => {
    const { name, value } = e.target;
    setFormData((prevState) => ({
      ...prevState,
      [name]:
        name === "category" || name === "gender" ? parseInt(value) : value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (type === "create") {
      await createParticipant(formData);
    } else {
      await updateParticipant(formData);
    }
    onClose();
  };

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal-content" onClick={(e) => e.stopPropagation()}>
        <h2>
          {type === "create" ? "Добавить участника" : "Изменить участника"}
        </h2>
        <form onSubmit={handleSubmit}>
          <label>
            ФИО:
            <input
              type="text"
              name="fio"
              value={formData.fio}
              onChange={handleChange}
              required
            />
          </label>
          <label>
            Категория:
            <select
              name="category"
              value={formData.category}
              onChange={handleChange}
              required
            >
              <option value="" disabled>
                Выберите категорию
              </option>
              {Object.entries(CategoryMap).map(([key, value]) => (
                <option key={key} value={key}>
                  {value}
                </option>
              ))}
            </select>
          </label>
          <label>
            Пол:
            <select
              name="gender"
              value={formData.gender}
              onChange={handleChange}
              required
            >
              <option value="" disabled>
                Выберите пол
              </option>
              {Object.entries(GenderMap).map(([key, value]) => (
                <option key={key} value={key}>
                  {value}
                </option>
              ))}
            </select>
          </label>
          <label>
            Дата рождения:
            <input
              type="date"
              name="birthday"
              value={formData.birthday}
              onChange={handleChange}
              required
            />
          </label>
          <label>
            Тренер:
            <input
              type="text"
              name="coach"
              value={formData.coach}
              onChange={handleChange}
              required
            />
          </label>
          <button type="submit">
            {type === "create" ? "Создать" : "Сохранить"}
          </button>
        </form>
        <button className="close-button" onClick={onClose}>
          Закрыть
        </button>
      </div>
    </div>
  );
};

export default ParticipantModal;
