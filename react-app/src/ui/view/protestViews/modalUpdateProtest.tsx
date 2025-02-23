import React, { useState } from "react";
import { ProtestInput } from "../../models/protestModel";
import { useUpdateProtest } from "../../controllers/protestControllers/updateProtestController";
import {useParams} from "react-router-dom";

interface UpdateProtestModalProps {
  onClose: () => void;
  protest: ProtestInput;
}

const UpdateProtestModal: React.FC<UpdateProtestModalProps> = ({
  onClose,
  protest,
}) => {
  const { updateProtest, loading } = useUpdateProtest();
  const [formData, setFormData] = useState<ProtestInput>(protest);
  const { ratingID, raceID, protestID } = useParams<{
    ratingID: string;
    raceID: string;
    protestID: string;
  }>();

  const handleChange = (
    e: React.ChangeEvent<
      HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement
    >,
  ) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: name === "status" || name === "ruleNum" ? Number(value) : value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    await updateProtest(ratingID || '', raceID || '', protestID || '', formData);
    onClose();
  };

  return (
    <div className="modal-overlay">
      <div className="modal-content">
        <h2>Обновить протест</h2>
        <form onSubmit={handleSubmit}>
          <label>
            Номер правила:
            <input
              type="number"
              name="ruleNum"
              value={formData.ruleNum}
              onChange={handleChange}
              required
            />
          </label>

          <label>
            Дата рассмотрения:
            <input
              type="date"
              name="reviewDate"
              value={formData.reviewDate}
              onChange={handleChange}
              required
            />
          </label>

          <label>
            Комментарий:
            <textarea
              name="comment"
              value={formData.comment}
              onChange={handleChange}
              required
            />
          </label>

          <div className="modal-buttons">
            <button type="submit" disabled={loading}>
              {loading ? "Обновление..." : "Обновить"}
            </button>
            <button type="button" onClick={onClose} className="cancel-btn">
              Отмена
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default UpdateProtestModal;
