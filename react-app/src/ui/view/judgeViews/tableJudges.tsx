import React, { useState, useEffect } from "react";
import { JudgeFormData } from "../../models/judgeModel";
import { useGetJudges } from "../../controllers/judgeControllers/getJudgesController";
import { useDeleteJudge } from "../../controllers/judgeControllers/deleteJudgeController"; // Хук для удаления судьи
import JudgeModal from "./modalInputJudge";

const JudgeTable: React.FC = () => {
  const [showJudgeModal, setShowJudgeModal] = useState(false);
  const [selectedJudge, setSelectedJudge] = useState<JudgeFormData | null>(
    null,
  );

  const {
    judges,
    loading: loadingJudges,
    error: judgesError,
    fetchJudges,
  } = useGetJudges();
  const {
    deleteJudge,
  } = useDeleteJudge(); // Хук для удаления судьи

  const [fioFilter, setFioFilter] = useState("");
  const [loginFilter, setLoginFilter] = useState("");
  const [roleFilter, setRoleFilter] = useState("");
  const [postFilter, setPostFilter] = useState("");

  const handleAddJudgeClick = () => {
    setSelectedJudge(null);
    setShowJudgeModal(true);
  };

  const handleJudgeModalClose = () => {
    setShowJudgeModal(false);
  };

  useEffect(() => {
    if (judgesError) {
      alert(`Ошибка загрузки судей: ${judgesError}`);
    }
  }, [judgesError]);

  useEffect(() => {
    if (judges.length === 0 && !loadingJudges) {
      fetchJudges();
    }
  }, [judges, fetchJudges, loadingJudges]);

  const handleDeleteJudge = (id: string) => {
    if (window.confirm("Вы уверены, что хотите удалить судью?")) {
      try {
        deleteJudge(id);
      } catch  {
        alert("Ошибка при удалении судьи");
      }
    }
  };

  const filteredJudges = judges.filter((judge) => {
    return (
      judge.fio.toLowerCase().includes(fioFilter.toLowerCase()) &&
      judge.login.toLowerCase().includes(loginFilter.toLowerCase()) &&
      judge.role.toLowerCase().includes(roleFilter.toLowerCase()) &&
      judge.post.toLowerCase().includes(postFilter.toLowerCase())
    );
  });

  if (loadingJudges) {
    return <div>Загрузка...</div>;
  }

  return (
    <div>
      {/* Таблица судей */}
      <h2>Судьи</h2>
      <div className={"tableContent"}>
        <table>
          <thead>
            <tr>
              <th className={"stickyHeader"}>ФИО</th>
              <th className={"stickyHeader"}>Логин</th>
              <th className={"stickyHeader"}>Роль</th>
              <th className={"stickyHeader"}>Должность</th>
              <th className={"stickyHeader"}>Действия</th>
            </tr>
            <tr>
              <th>
                <input
                  type="text"
                  value={fioFilter}
                  onChange={(e) => setFioFilter(e.target.value)}
                  placeholder="ФИО"
                />
              </th>
              <th>
                <input
                  type="text"
                  value={loginFilter}
                  onChange={(e) => setLoginFilter(e.target.value)}
                  placeholder="Логин"
                />
              </th>
              <th>
                <input
                  type="text"
                  value={roleFilter}
                  onChange={(e) => setRoleFilter(e.target.value)}
                  placeholder="Роль"
                />
              </th>
              <th>
                <input
                  type="text"
                  value={postFilter}
                  onChange={(e) => setPostFilter(e.target.value)}
                  placeholder="Должность"
                />
              </th>
              <th></th>
            </tr>
          </thead>
          <tbody>
            {filteredJudges.map((judge) => (
              <tr key={judge.id}>
                <td>{judge.fio}</td>
                <td>{judge.login}</td>
                <td>{judge.role}</td>
                <td>{judge.post}</td>
                <td>
                  <button
                    onClick={() => {
                      setSelectedJudge(judge);
                      setShowJudgeModal(true);
                    }}
                  >
                    Редактировать
                  </button>
                  <button onClick={() => handleDeleteJudge(judge.id)}>
                    Удалить
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
      <button onClick={handleAddJudgeClick}>Добавить судью</button>

      {/* Модалки для добавления/редактирования */}
      {showJudgeModal && (
        <JudgeModal
          judge={{
            id: selectedJudge ? selectedJudge.id : "",
            fio: selectedJudge ? selectedJudge.fio : "",
            login: selectedJudge ? selectedJudge.login : "",
            password: "",
            role: selectedJudge
              ? selectedJudge.role
                ? parseInt(selectedJudge.role, 10)
                : undefined
              : undefined,
            post: selectedJudge ? selectedJudge.post : "",
          }}
          type={selectedJudge ? "update" : "create"}
          onClose={handleJudgeModalClose}
        />
      )}
    </div>
  );
};

export default JudgeTable;
