import React, { useState, useEffect } from "react";
import { CrewFormData } from "../../models/crewModel";
import { useCreateCrew } from "../../controllers/crewControllers/createCrewController";
import { useUpdateCrew } from "../../controllers/crewControllers/updateCrewController";

interface CrewModalProps {
  ratingID: string;
  crew: CrewFormData;
  type: "update" | "create";
  onClose: () => void;
}

const CrewModal: React.FC<CrewModalProps> = ({ ratingID, crew, type, onClose }) => {
  const [localCrew, setLocalCrew] = useState<CrewFormData>(crew); // Используем CrewInput для состояния
  const [errorMessage, setErrorMessage] = useState<string | null>(null); // Состояние для отображения ошибки
  const { handleSubmit } = useCreateCrew();
  const { handleUpdate } = useUpdateCrew();

  // Восстановление состояния из localStorage при открытии модалки
  useEffect(() => {
    const storedCrew = localStorage.getItem(`crew-${ratingID}`);
    if (storedCrew) {
      setLocalCrew(JSON.parse(storedCrew));
    } else if (type === "create") {
      setLocalCrew({
        ...crew,
        SailNum: 1,
      });
    } else {
      setLocalCrew(crew); // Для обновления используем переданный рейтинг
    }
  }, [type, crew, ratingID]);

  // Обработчик изменения значения номера паруса
  const handleSailNumChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setLocalCrew((prev) => {
      const newCrew = {
        ...prev,
        SailNum: parseInt(e.target.value) || 0,
      };
      // Сохраняем изменённое состояние в localStorage
      localStorage.setItem(`crew-${ratingID}`, JSON.stringify(newCrew));
      return newCrew;
    });
  };

  // Сохранение данных
  const handleSave = async () => {
    setErrorMessage(null); // Сбрасываем сообщение об ошибке
    const crewData = {
      sailNum: localCrew.SailNum,
    };

    try {
      if (type === "update") {
        await handleUpdate(ratingID, crew.id, crewData);
      } else {
        await handleSubmit(ratingID, crewData);
      }
      onClose();
      // После закрытия очищаем данные из localStorage
      localStorage.removeItem(`crew-${ratingID}`);
    } catch (error: any) {
      if (error.response && error.response.data) {
        setErrorMessage(
            error.response.data.error || "Ошибка при сохранении данных"
        );
      } else {
        setErrorMessage("Произошла ошибка. Попробуйте позже.");
      }
    }
  };

  // Обработчик закрытия модалки
  const handleClose = () => {
    onClose();
    // После закрытия очищаем данные из localStorage
    localStorage.removeItem(`crew-${ratingID}`);
  };

  return (
      <div className="modal-overlay">
        <div className="modal-content">
          <h3>{type === "update" ? "Обновить команду" : "Создать новую команду"}</h3>
          {errorMessage && (
              <div className="error-message">
                <p>{errorMessage}</p>
                <button onClick={() => setErrorMessage(null)}>Закрыть</button>
              </div>
          )}
          <input
              type="number"
              value={localCrew.SailNum}
              onChange={handleSailNumChange}
              placeholder="Номер паруса"
          />
          <div className="buttons-container">
            <button onClick={handleSave}>Сохранить изменения</button>
            <button onClick={handleClose}>Отмена</button>
          </div>
        </div>
      </div>
  );
};

export default CrewModal;
