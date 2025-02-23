import React from "react";
import JudgeTable from "./tableJudges";
import ParticipantTable from "./tableParticipants";

const JudgeDashboard: React.FC = () => {
  return (
    <div className="judge-dashboard">
      <h1 className={"headerH1"}>Панель судьи</h1>

      {/* Таблица судей */}
      <JudgeTable />
      {/* Таблица участников */}
      <ParticipantTable />
    </div>
  );
};

export default JudgeDashboard;
