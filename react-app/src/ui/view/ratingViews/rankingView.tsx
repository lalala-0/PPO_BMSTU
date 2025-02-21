import React, { useState, useMemo } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { RatingTableLine, RaceInfo } from "../../models/ratingModel";
import { useDeleteCrew } from "../../controllers/crewControllers/deleteCrewController";

interface RatingTableProps {
  rankingTable: RatingTableLine[];
  races: RaceInfo[];
}

const RatingTable: React.FC<RatingTableProps> = ({ rankingTable, races }) => {
  const [tableData, setTableData] = useState<RatingTableLine[]>(rankingTable);
  const [filters, setFilters] = useState({
    sailNum: "",
    participantName: "",
    pointsSum: "",
    rank: "",
    coachName: "",
  });
  const navigate = useNavigate();
  const { ratingID } = useParams<{ ratingID: string }>();

  const deleteCrewController = useDeleteCrew();
  const onDelete = async (sailNum: number) => {
    const crewToDelete = tableData.find((line) => line.SailNum === sailNum);

    if (!crewToDelete) {
      console.error("Команда не найдена");
      return;
    }

    const CrewID = crewToDelete.crewID;

    try {
      await deleteCrewController.deleteCrewByID(ratingID || '', CrewID);
      const updatedTable = tableData.filter((line) => line.SailNum !== sailNum);
      setTableData(updatedTable);
    } catch (err) {
      console.error("Ошибка при удалении записи:", err);
      alert("Не удалось удалить запись. Попробуйте снова.");
    }
  };

  const handleCrewNavigation = (sailNum: number) => {
    const crew = rankingTable.find((line) => line.SailNum === sailNum);
    if (crew) {
      const CrewID = crew.crewID;
      if (ratingID) {
        navigate(`/ratings/${ratingID}/crews/${CrewID}`);
      } else {
        console.error("Rating ID is missing in the current path.");
      }
    } else {
      console.error("Команда с таким номером паруса не найдена");
    }
  };

  const handleRaceNavigation = (raceID: string) => {
    if (ratingID) {
      navigate(`/ratings/${ratingID}/races/${raceID}`);
    } else {
      console.error("Rating ID is missing in the current path.");
    }
  };

  const handleFilterChange = (key: string, value: string) => {
    setFilters((prevFilters) => ({ ...prevFilters, [key]: value }));
  };

  // Использование useMemo для фильтрации данных
  const filteredData = useMemo(() => {
    return tableData.filter((line) => {
      return (
          (filters.sailNum === "" ||
              line.SailNum.toString().includes(filters.sailNum)) &&
          (filters.participantName === "" ||
              line.ParticipantNames[0]
                  .toLowerCase()
                  .includes(filters.participantName.toLowerCase())) &&
          (filters.pointsSum === "" ||
              line.PointsSum.toString().includes(filters.pointsSum)) &&
          (filters.rank === "" || line.Rank.toString().includes(filters.rank)) &&
          (filters.coachName === "" ||
              line.CoachNames[0]
                  .toLowerCase()
                  .includes(filters.coachName.toLowerCase()))
      );
    });
  }, [tableData, filters]);

  return (
      <div className={"table-container"}>
          <div className={"tableContent"}>
            <table className={"table"}>
              <thead>
              <tr>
                <th className={"stickyHeader"}>
                  <input
                      type="text"
                      placeholder="Поиск по номеру паруса"
                      value={filters.sailNum}
                      onChange={(e) => handleFilterChange("sailNum", e.target.value)}
                  />
                  Номер паруса
                </th>
                <th className={"stickyHeader"}>
                  <input
                      type="text"
                      placeholder="Поиск по имени"
                      value={filters.participantName}
                      onChange={(e) =>
                          handleFilterChange("participantName", e.target.value)
                      }
                  />
                  Имя участника
                </th>
                <th className={"stickyHeader"}>
                  Дата рождения
                </th>
                {races.map((race) => (
                    <th className={"stickyHeader"} key={race.RaceID}>
                      <button onClick={() => handleRaceNavigation(race.RaceID)}>
                        {`Гонка ${race.RaceNum}`}
                      </button>
                    </th>
                ))}
                <th className={"stickyHeader"}>
                  <input
                      type="text"
                      placeholder="Поиск по баллам"
                      value={filters.rank}
                      onChange={(e) => handleFilterChange("rank", e.target.value)}
                  />
                  Сумма баллов
                </th>
                <th className={"stickyHeader"}>
                  <input
                      type="text"
                      placeholder="Поиск по рангу"
                      value={filters.rank}
                      onChange={(e) => handleFilterChange("rank", e.target.value)}
                  />
                  Ранг
                </th>
                <th className={"stickyHeader"}>
                  <input
                      type="text"
                      placeholder="Поиск по тренеру"
                      value={filters.coachName}
                      onChange={(e) =>
                          handleFilterChange("coachName", e.target.value)
                      }
                  />
                  Имя тренера
                </th>
                <th className={"stickyHeader"}>
                  <th className={"auth-required"}>Действия</th>
                  </th>
              </tr>
              </thead>
              <tbody>
              {filteredData.map((line) => (
                  <tr key={line.SailNum}>
                    <td>
                      <button onClick={() => handleCrewNavigation(line.SailNum)}>
                        {line.SailNum}
                      </button>
                    </td>
                    <td>{line.ParticipantNames[0]}</td>
                    <td>{line.ParticipantBirthDates[0]}</td>
                    {races.map((race) => (
                        <td key={race.RaceID}>
                          {line.ResInRace[race.RaceNum] || "Нет данных"}
                        </td>
                    ))}
                    <td>{line.PointsSum}</td>
                    <td>{line.Rank}</td>
                    <td>{line.CoachNames[0]}</td>
                    <td className={"auth-required"}>
                      <button
                          onClick={() => onDelete(line.SailNum)}
                      >
                        <img
                            src="/delete-icon.svg"
                            alt="Удалить"
                            width="20"
                            height="20"
                        />
                      </button>
                    </td>
                  </tr>
              ))}
              </tbody>
            </table>
          </div>
      </div>
  );
};

export default RatingTable;
