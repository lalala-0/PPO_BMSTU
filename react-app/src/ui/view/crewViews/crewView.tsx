import React, { useState } from "react";
import { useParams } from "react-router-dom";
import CrewMembersTable from "./tableMembers";
import CrewModal from "./modalInputCrew";
// import MemberModal from '../crewMemberViews/modalInputMember';
import { useGetRating } from "../../controllers/ratingControllers/getRatingController";
import { useGetCrew } from "../../controllers/crewControllers/getCrewController";
import "../../controllers/crewControllers/updateCrewController";
import "../../styles/styles.css";

const CrewView: React.FC = () => {
  const { ratingID, crewID } = useParams<{
    ratingID: string;
    crewID: string;
  }>(); // Явно указываем типы для ratingID и crewID

  // Хуки вызываются всегда, независимо от условий
  const { ratingInfo, loading: ratingLoading } = useGetRating();
  const { crewInfo, loading: crewLoading } = useGetCrew();

  const [isCrewModalOpen, setIsCrewModalOpen] = useState(false);
  const [isMemberModalOpen, setIsMemberModalOpen] = useState(false);

  const handleCrewModalOpen = () => setIsCrewModalOpen(true);
  const handleCrewModalClose = () => setIsCrewModalOpen(false);

  const handleMemberModalOpen = () => setIsMemberModalOpen(true);
  const handleMemberModalClose = () => setIsMemberModalOpen(false);

  // Проверка на наличие параметров URL и возвращение ошибки, если они отсутствуют
  if (!ratingID || !crewID) {
    return <div>Неверные параметры URL!</div>;
  }

  // Теперь можно проверять загрузку и отображать соответствующие сообщения
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
        <button onClick={handleMemberModalOpen}>Добавить участника</button>
        <button onClick={handleCrewModalOpen}>
          Изменить информацию о команде
        </button>
      </div>

      {isCrewModalOpen && crewID && (
        <CrewModal
          crew={crewInfo}
          type={"update"}
          onClose={handleCrewModalClose}
        />
      )}

      {/*{isMemberModalOpen && (*/}
      {/*    <MemberModal*/}
      {/*        member={{ id: '', crewId: crewID, name: '', role: '' }}*/}
      {/*        crewID={crewID}*/}
      {/*        type={'create'}*/}
      {/*        onClose={handleMemberModalClose}*/}
      {/*    />*/}
      {/*)}*/}
    </div>
  );
};

export default CrewView;
