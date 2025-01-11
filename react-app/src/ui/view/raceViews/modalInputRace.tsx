import React, { useState } from "react";
import { RaceFormData } from "../../models/raceModel";
import { useUpdateRace } from "../../controllers/raceControllers/updateRaceController";
import { useCreateRace } from "../../controllers/raceControllers/createRaceController";
import { classOptions } from "../../models/classOptions";

interface RaceModalProps {
  race: RaceFormData;
  type: "update" | "create";
  onClose: () => void;
}

const RaceModal: React.FC<RaceModalProps> = ({ race, type, onClose }) => {
  const [formData, setFormData] = useState({
    date: race.date || "",
    number: race.number || 1,
    class: race.class || "",
  });
  const [error, setError] = useState<string | null>(null);

  const { handleSubmit } = useCreateRace();
  const { handleUpdate } = useUpdateRace();

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>,
  ) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: name === "number" ? parseInt(value) : value,
    }));
  };

  const handleSave = async (e: React.FormEvent) => {
    e.preventDefault();
    const raceData = {
      date: formData.date,
      class:
        classOptions.find((option) => option.label === formData.class)?.value ||
        1,
      number: formData.number,
    };

    try {
      if (type === "update") {
        await handleUpdate(race.id, raceData);
      } else {
        await handleSubmit(race.ratingId);
      }
      onClose();
    } catch (error: any) {
      const message =
        error.response?.data?.message || "Произошла ошибка. Попробуйте позже.";
      setError(message);
    }
  };

  return (
    <div className="modal-overlay">
      <div className="modal-content">
        <h2>
          {type === "update" ? "Обновить информацию о гонке" : "Создать гонку"}
        </h2>
        {error && <div className="error-message">{error}</div>}
        <form onSubmit={handleSave}>
          <label>
            Дата:
            <input
              type="date"
              name="date"
              value={formData.date}
              onChange={handleChange}
              required
            />
          </label>
          <label>
            Номер:
            <input
              type="number"
              name="number"
              value={formData.number}
              onChange={handleChange}
              required
            />
          </label>
          <label>
            Класс:
            <select
              name="class"
              value={formData.class}
              onChange={handleChange}
              required
            >
              <option value="">Выберите класс</option>
              {classOptions.map((option) => (
                <option key={option.value} value={option.label}>
                  {option.label}
                </option>
              ))}
            </select>
          </label>
          <div className="modal-actions">
            <button type="submit" className="save-button">
              Сохранить
            </button>
            <button type="button" onClick={onClose} className="cancel-button">
              Отмена
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default RaceModal;
