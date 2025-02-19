import React, { useState } from "react";
import { useParams } from "react-router-dom";
import CrewMembersTable from "./tableMembers";
import CrewModal from "./modalInputCrew";
import { useGetRating } from "../../controllers/ratingControllers/getRatingController";
import { useGetCrew } from "../../controllers/crewControllers/getCrewController";
import "../../controllers/crewControllers/updateCrewController";

import AttachCrewMemberModal from "../participantViews/attachParticipantToCrew";

const CrewView: React.FC = () => {
  const { ratingID, crewID } = useParams<{
    ratingID: string;
    crewID: string;
  }>();

  const { ratingInfo, loading: ratingLoading } = useGetRating();
  const { crewInfo, loading: crewLoading } = useGetCrew();

  const [isCrewModalOpen, setIsCrewModalOpen] = useState(false);
  const [isMemberModalOpen, setIsMemberModalOpen] = useState(false);

  const handleCrewModalOpen = () => setIsCrewModalOpen(true);
  const handleCrewModalClose = () => setIsCrewModalOpen(false);

  const handleMemberModalOpen = () => setIsMemberModalOpen(true);
  const handleMemberModalClose = () => setIsMemberModalOpen(false);

  if (!ratingID || !crewID) {
    return <div>Неверные параметры URL!</div>;
  }

  if (ratingLoading || crewLoading) {
    return <div>Загрузка...</div>;
  }

  if (!ratingInfo || !crewInfo) {
    return <div>Информация не найдена!</div>;
  }

  return (
    <div style={{ padding: "20px" }}>
      <h1>{`Рейтинг: ${ratingInfo.Name}`}</h1>
      <h2>{`Команда: ${crewInfo.SailNum}`}</h2>

      <CrewMembersTable ratingID={ratingID} crewID={crewID} />

      <div className="buttons-container">
        <button className="auth-required" onClick={handleMemberModalOpen}>
          Добавить участника
        </button>
        <button className="auth-required" onClick={handleCrewModalOpen}>
          Изменить информацию о команде
        </button>
      </div>

      {isCrewModalOpen && crewID && (
        <CrewModal
          ratingID={ratingID}
          crew={crewInfo}
          type="update"
          onClose={handleCrewModalClose}
        />
      )}

      {isMemberModalOpen && (
        <AttachCrewMemberModal
          ratingID={ratingID}
          crewID={crewID}
          onClose={handleMemberModalClose}
          onSuccess={() => window.location.reload()} // Обновляем страницу после добавления участника
        />
      )}
    </div>
  );
};

export default CrewView;
