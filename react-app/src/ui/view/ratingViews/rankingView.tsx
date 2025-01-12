import React, { useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { RatingTableLine, RaceInfo } from "../../models/ratingModel";
import { useDeleteCrew } from "../../controllers/crewControllers/deleteCrewController";
import { useUpdateCrew } from "../../controllers/crewControllers/updateCrewController";

interface RatingTableProps {
  rankingTable: RatingTableLine[];
  races: RaceInfo[];
}

const RatingTable: React.FC<RatingTableProps> = ({ rankingTable, races }) => {
  const [tableData, setTableData] = useState<RatingTableLine[]>(rankingTable);
  const [filters, setFilters] = useState({
    sailNum: "",
    participantName: "",
    rank: "",
    coachName: "",
  });
  const navigate = useNavigate(); // Хук для программного редиректа
  const { ratingID } = useParams<{ ratingID: string }>(); // Получаем ratingID из текущего пути

  const deleteCrewController = useDeleteCrew();
  const onDelete = async (sailNum: number) => {
    const crewToDelete = tableData.find((line) => line.SailNum === sailNum);

    if (!crewToDelete) {
      console.error("Команда не найдена");
      return;
    }

    const { CrewID } = crewToDelete;

    try {
      await deleteCrewController.deleteCrewByID(CrewID); // Передаем ratingID и crewID
      // Если удаление прошло успешно, обновляем таблицу
      const updatedTable = tableData.filter((line) => line.SailNum !== sailNum);
      setTableData(updatedTable); // Удаляем строку из таблицы
    } catch (err) {
      console.error("Ошибка при удалении записи:", err);
      alert("Не удалось удалить запись. Попробуйте снова.");
    }
  };

  const handleCrewNavigation = (sailNum: number) => {
    const crew = rankingTable.find((line) => line.SailNum === sailNum);

    if (crew) {
      const CrewID = crew.CrewID; // Предположим, что ID команды хранится в поле TeamID

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

  const filteredData = tableData.filter((line) => {
    return (
      (filters.sailNum === "" ||
        line.SailNum.toString().includes(filters.sailNum)) &&
      (filters.participantName === "" ||
        line.ParticipantNames[0]
          .toLowerCase()
          .includes(filters.participantName.toLowerCase())) &&
      (filters.rank === "" || line.Rank.toString().includes(filters.rank)) &&
      (filters.coachName === "" ||
        line.CoachNames[0]
          .toLowerCase()
          .includes(filters.coachName.toLowerCase()))
    );
  });

  return (
    <div style={{ padding: "20px" }}>
      <div
        style={{
          maxWidth: "100%",
          overflowX: "auto",
          border: "1px solid #ccc",
        }}
      >
        <div style={{ maxHeight: "400px", overflowY: "auto" }}>
          <table
            style={{
              tableLayout: "auto",
              width: "100%",
              borderCollapse: "collapse",
            }}
          >
            <thead>
              <tr>
                <th>
                  <input
                    type="text"
                    placeholder="Поиск по номеру паруса"
                    value={filters.sailNum}
                    onChange={(e) =>
                      handleFilterChange("sailNum", e.target.value)
                    }
                    style={{ width: "100%" }}
                  />
                  Номер паруса
                </th>
                <th>
                  <input
                    type="text"
                    placeholder="Поиск по имени"
                    value={filters.participantName}
                    onChange={(e) =>
                      handleFilterChange("participantName", e.target.value)
                    }
                    style={{ width: "100%" }}
                  />
                  Имя участника
                </th>
                <th>Дата рождения</th>
                {races.map((race) => (
                  <th key={race.RaceID}>
                    <button onClick={() => handleRaceNavigation(race.RaceID)}>
                      {`Гонка ${race.RaceNum}`}
                    </button>
                  </th>
                ))}
                <th>
                  <input
                    type="text"
                    placeholder="Поиск по баллам"
                    value={filters.rank}
                    onChange={(e) => handleFilterChange("rank", e.target.value)}
                    style={{ width: "100%" }}
                  />
                  Сумма баллов
                </th>
                <th>
                  <input
                    type="text"
                    placeholder="Поиск по рангу"
                    value={filters.rank}
                    onChange={(e) => handleFilterChange("rank", e.target.value)}
                    style={{ width: "100%" }}
                  />
                  Ранг
                </th>
                <th>
                  <input
                    type="text"
                    placeholder="Поиск по тренеру"
                    value={filters.coachName}
                    onChange={(e) =>
                      handleFilterChange("coachName", e.target.value)
                    }
                    style={{ width: "100%" }}
                  />
                  Имя тренера
                </th>
                <th>Действия</th>
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
                  <td>
                    <button onClick={() => onDelete(line.SailNum)}>
                      Удалить
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
};

export default RatingTable;
