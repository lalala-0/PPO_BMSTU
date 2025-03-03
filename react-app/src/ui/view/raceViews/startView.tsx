import React, { useState, useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { useStartProcedure } from "../../controllers/raceControllers/startProcedureController";
import { SpecCircumstanceMap } from "../../models/raceModel";
import { useGetCrewsByRatingID } from "../../controllers/crewControllers/getCrewsController";
import { useGetRace } from "../../controllers/raceControllers/getRaceController";
import { useGetRating } from "../../controllers/ratingControllers/getRatingController";

const StartProcedure: React.FC = () => {
  const { ratingID, raceID } = useParams<{
    ratingID: string;
    raceID: string;
  }>();
  const navigate = useNavigate();

  const { ratingInfo, loading: ratingLoading } = useGetRating();
  const { raceInfo, loading: raceLoading } = useGetRace();
  const { loading, error, success, startProcedure } = useStartProcedure(
      ratingID!,
      raceID!,
  );
  const {
    loading: crewsLoading,
    error: crewsError,
    crews,
    getCrewsByRatingID,
  } = useGetCrewsByRatingID();

  const [specCircumstance, setSpecCircumstance] = useState<number>(() => {
    return parseInt(localStorage.getItem("specCircumstance") || "0", 10);
  });
  const [search, setSearch] = useState<string>(
      localStorage.getItem("search") || ""
  );
  const [falseStartList, setFalseStartList] = useState<number[]>(() => {
    const savedFalseStartList = localStorage.getItem("falseStartList");
    return savedFalseStartList ? JSON.parse(savedFalseStartList) : [];
  });
  const [availableCrews, setAvailableCrews] = useState<number[]>([]);

  useEffect(() => {
    if (ratingID) {
      getCrewsByRatingID(ratingID);
    }
  }, [ratingID, getCrewsByRatingID]);

  useEffect(() => {
    if (success) {
      navigate(`/ratings/${ratingID}/races/${raceID}`);
    }
  }, [success, navigate, ratingID, raceID]);

  useEffect(() => {
    if (crews) {
      setAvailableCrews(crews.map((crew) => crew.SailNum));
    }
  }, [crews]);

  useEffect(() => {
    localStorage.setItem("specCircumstance", specCircumstance.toString());
  }, [specCircumstance]);

  useEffect(() => {
    localStorage.setItem("search", search);
  }, [search]);

  useEffect(() => {
    localStorage.setItem("falseStartList", JSON.stringify(falseStartList));
  }, [falseStartList]);

  const handleSpecCircumstanceChange = (
      e: React.ChangeEvent<HTMLSelectElement>
  ) => {
    setSpecCircumstance(parseInt(e.target.value));
  };

  const handleSearchChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setSearch(e.target.value);
  };

  const toggleFalseStart = (yachtNumber: number) => {
    if (falseStartList.includes(yachtNumber)) {
      setFalseStartList(falseStartList.filter((num) => num !== yachtNumber));
      setAvailableCrews([...availableCrews, yachtNumber]);
    } else {
      setAvailableCrews(availableCrews.filter((num) => num !== yachtNumber));
      setFalseStartList([...falseStartList, yachtNumber]);
    }
  };

  const handleSubmit = () => {
    startProcedure({ specCircumstance, falseStartList });
    localStorage.removeItem("falseStartList");
    localStorage.removeItem("search");
    localStorage.removeItem("specCircumstance");
  };

  if (ratingLoading || raceLoading || crewsLoading)
    return <p>Загрузка данных...</p>;
  if (!ratingInfo || !raceInfo)
    return <p className="error">Ошибка загрузки данных</p>;
  if (crewsError)
    return <p className="error">Ошибка загрузки яхт: {crewsError}</p>;

  return (
      <div className="start-procedure">
        <h1 className={"headerH1"}>
          {ratingInfo.Name}
        </h1>
        <h2 className={"headerH2"}>
          Гонка {raceInfo.number}
        </h2>
        <h3 className={"headerH3"}>
          Стартовая процедура
        </h3>
        <label>
          Наказание за фальшстарт:
          <select
              value={specCircumstance}
              onChange={handleSpecCircumstanceChange}
          >
            {Object.entries(SpecCircumstanceMap).map(([key, value]) => (
                <option key={key} value={key}>
                  {value}
                </option>
            ))}
          </select>
        </label>

        <div className="finish-controls">
          <div className="yacht-icons">
            <label>
              <input
                  type="text"
                  placeholder="Поиск по номеру яхты"
                  value={search}
                  onChange={handleSearchChange}
                  className="search-input"
              />
            </label>
            {availableCrews
                .filter((num) => num.toString().includes(search))
                .map((num) => (
                    <div
                        key={num}
                        className="yacht-button"
                        onClick={() => toggleFalseStart(num)}
                    >
                      {num}
                    </div>
                ))}
          </div>
          <div className="false-start-container">
            <h3>Фальшстарты</h3>
            <div className="false-start-area">
              {falseStartList.map((num) => (
                  <div
                      key={num}
                      className="selected-yacht-button"
                      onClick={() => toggleFalseStart(num)}
                  >
                    {num}
                  </div>
              ))}
            </div>
          </div>
        </div>

        <button onClick={handleSubmit} disabled={loading}>
          Завершить стартовую процедуру
        </button>

        {error && <p className="error">Ошибка: {error}</p>}
        {success && <p className="success">{success}</p>}
      </div>
  );
};

export default StartProcedure;
