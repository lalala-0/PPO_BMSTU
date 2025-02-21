import React, { useState, useEffect } from "react";
import { useUpdateJudge } from "../../controllers/judgeControllers/updateJudgeController";
import { useCreateJudge } from "../../controllers/judgeControllers/createJudgeController";
import { JudgeInput } from "../../models/judgeModel";

interface JudgeModalProps {
  judge?: JudgeInput | null; // Если передан, значит редактирование
  type: "create" | "update";
  onClose: () => void;
}

const JudgeModal: React.FC<JudgeModalProps> = ({ judge, type, onClose }) => {
  const validJudgeID = sessionStorage.getItem("judgeID") ?? "";

  const { updateJudge, loading: updateLoading } = useUpdateJudge(validJudgeID);
  const { createJudge, loading: createLoading } = useCreateJudge();

  const [formData, setFormData] = useState<JudgeInput>({
    id: validJudgeID,
    fio: judge?.fio || "",
    login: judge?.login || "",
    password: "",
    role: judge?.role || undefined,
    post: judge?.post || "",
  });

  useEffect(() => {
    if (judge) {
      setFormData({
        id: validJudgeID,
        fio: judge.fio,
        login: judge.login,
        password: "",
        role: judge.role,
        post: judge.post,
      });
    }
  }, [judge, validJudgeID]);

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>,
  ) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: name === "role" ? Number(value) : value, // Преобразуем `role` в число
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (type === "update") {
      await updateJudge(formData);
    } else {
      await createJudge(formData);
    }
    onClose();
  };

  return (
    <div className="modal-overlay">
      <div className="modal-content">
        <h2>{type === "update" ? "Редактировать судью" : "Создать судью"}</h2>
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
            Логин:
            <input
              type="text"
              name="login"
              value={formData.login}
              onChange={handleChange}
              required
            />
          </label>
          <label>
            Пароль:
            <input
              type="password"
              name="password"
              value={formData.password}
              onChange={handleChange}
            />
          </label>
          <label>
            Роль:
            <select
              name="role"
              value={formData.role ?? ""}
              onChange={handleChange}
            >
              <option value="">Выберите роль</option>
              <option value="1">Главный судья</option>
              <option value="2">Ассистент</option>
            </select>
          </label>
          <label>
            Должность:
            <input
              type="text"
              name="post"
              value={formData.post}
              onChange={handleChange}
            />
          </label>

          <div className="modal-buttons">
            <button type="submit" disabled={updateLoading || createLoading}>
              {type === "update" ? "Сохранить" : "Создать"}
            </button>
            <button type="button" onClick={onClose}>
              Отмена
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default JudgeModal;
