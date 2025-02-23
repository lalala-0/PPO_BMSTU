import { useState } from "react";
import { useCompleteProtest } from "../../controllers/protestControllers/completeProtestController";
import React from "react";
import {useParams} from "react-router-dom";

interface CompleteProtestModalProps {
  isOpen: boolean;
  onClose: () => void;
}

export const CompleteProtestModal: React.FC<CompleteProtestModalProps> = ({
  isOpen,
  onClose,
}) => {
  const { completeProtest, loading, error } = useCompleteProtest();
  const [resPoints, setResPoints] = useState<number>(0);
  const [comment, setComment] = useState<string>("");
  const { ratingID, raceID, protestID } = useParams<{
    ratingID: string;
    raceID: string;
    protestID: string;
  }>();
  const handleSubmit = async () => {
    await completeProtest(ratingID || '', raceID || '', protestID || '', { resPoints, comment });
    onClose(); // Закрываем модалку после завершения
  };

  if (!isOpen) return null;

  return (
    <div className="modal-overlay">
      <div className="modal-content">
        <h2>Завершение рассмотрения протеста</h2>
        {error && <p className="error">{error}</p>}
        <label>Результирующие очки:</label>
        <input
          type="number"
          value={resPoints}
          onChange={(e) => setResPoints(Number(e.target.value))}
        />

        <label>Комментарий:</label>
        <textarea
          value={comment}
          onChange={(e) => setComment(e.target.value)}
        />

        <button onClick={handleSubmit} disabled={loading}>
          {loading ? "Отправка..." : "Завершить"}
        </button>
        <button onClick={onClose}>Отмена</button>
      </div>
    </div>
  );
};
