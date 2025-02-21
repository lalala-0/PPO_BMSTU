import React, { useState, useEffect } from "react";
import { useGetJudge } from "../../controllers/judgeControllers/getJudgeController";
import JudgeModal from "./modalInputJudge";

const JudgeProfile: React.FC = () => {
  const [judgeID, setJudgeID] = useState<string | null>(null); // Хранение ID судьи
  const { judge, loading, error, fetchJudge } = useGetJudge(judgeID || ""); // Получаем данные судьи через хук
  const [isModalOpen, setIsModalOpen] = useState<boolean>(false); // Для управления состоянием модального окна
  const [modalType, setModalType] = useState<"create" | "update">("update"); // Тип модального окна: создание или редактирование

  // При монтировании компонента, извлекаем judgeID из sessionStorage
  useEffect(() => {
    const storedJudgeID = sessionStorage.getItem("judgeID");
    if (storedJudgeID) {
      setJudgeID(storedJudgeID); // Устанавливаем judgeID в состояние
    } else {
      console.error("ID судьи не найден в sessionStorage");
    }
  }, []);

  // Эффект для загрузки данных судьи, если judgeID изменился
  useEffect(() => {
    if (judgeID && !judge) {
      // Добавляем проверку, чтобы fetchJudge не вызывался, если данные уже загружены
      fetchJudge(); // Загружаем данные судьи при изменении judgeID
    }
  }, [judgeID, judge, fetchJudge]); // Мы добавляем judge в зависимости, чтобы не делать лишних запросов

  const handleUpdate = () => {
    setModalType("update");
    setIsModalOpen(true); // Открываем модальное окно для обновления информации
  };

  const handleCloseModal = () => {
    setIsModalOpen(false); // Закрываем модальное окно
  };

  if (loading) return <div>Загрузка...</div>;
  if (error) return <div>Ошибка: {error}</div>;
  if (!judge) return <div>Судья не найден</div>;

  return (
    <div>
      <h1>Профиль судьи</h1>
      <div>
        <p>
          <strong>ФИО:</strong> {judge.fio}
        </p>
        <p>
          <strong>Логин:</strong> {judge.login}
        </p>
        <p>
          <strong>Роль:</strong> {judge.role}
        </p>
        <p>
          <strong>Должность:</strong> {judge.post}
        </p>
      </div>
      <button onClick={handleUpdate}>Редактировать информацию</button>

      {/* Здесь будет модальное окно для обновления информации */}
      {isModalOpen && (
        <JudgeModal
          judge={{
            id: judge.id,
            fio: judge.fio,
            login: judge.login,
            password: "", // Если пароля нет, передаем пустую строку
            role: judge.role ? parseInt(judge.role, 10) : undefined, // Преобразуем role в число, если оно не пустое
            post: judge.post,
          }}
          type={modalType}
          onClose={handleCloseModal}
        />
      )}
    </div>
  );
};

export default JudgeProfile;
