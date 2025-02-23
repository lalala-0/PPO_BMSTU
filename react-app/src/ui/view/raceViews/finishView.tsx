import React, { useState, useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { useFinishProcedure } from "../../controllers/raceControllers/finishProcedureController";
import { useGetCrewsByRatingID } from "../../controllers/crewControllers/getCrewsController";
import { useGetRace } from "../../controllers/raceControllers/getRaceController";
import { useGetRating } from "../../controllers/ratingControllers/getRatingController";
import { FinishInput } from "../../models/raceModel";

// DnD Kit
import { DndContext, closestCenter } from "@dnd-kit/core";
import {
  SortableContext,
  arrayMove,
  verticalListSortingStrategy,
} from "@dnd-kit/sortable";
import { useSortable } from "@dnd-kit/sortable";
import { CSS } from "@dnd-kit/utilities";

interface SortableItemProps {
  yachtNumber: number;
}

const SortableItem: React.FC<SortableItemProps> = ({ yachtNumber }) => {
  const { attributes, listeners, setNodeRef, transform, transition } =
      useSortable({ id: yachtNumber });

  const style = {
    transform: CSS.Transform.toString(transform),
    transition,
  };

  return (
      <div
          ref={setNodeRef}
          style={style}
          {...attributes}
          {...listeners}
          className="sortable-item"
      >
        {yachtNumber}
      </div>
  );
};

const FinishProcedure: React.FC = () => {
  const { ratingID, raceID } = useParams<{
    ratingID?: string;
    raceID?: string;
  }>();
  const navigate = useNavigate();

  const { ratingInfo, loading: ratingLoading } = useGetRating();
  const { raceInfo, loading: raceLoading } = useGetRace();
  const { loading, error, success, finishProcedure } = useFinishProcedure(
      ratingID!,
      raceID!
  );
  const {
    loading: crewsLoading,
    crews,
    getCrewsByRatingID,
  } = useGetCrewsByRatingID();

  const [search, setSearch] = useState<string>("");
  const [finisherList, setFinisherList] = useState<number[]>([]);
  const [availableYachts, setAvailableYachts] = useState<number[]>([]);

  useEffect(() => {
    if (ratingID) {
      getCrewsByRatingID(ratingID);
    }
  }, [ratingID, getCrewsByRatingID]);

  useEffect(() => {
    if (crews) {
      setAvailableYachts(
          crews.map((crew) => crew.SailNum).sort((a, b) => a - b)
      );
    }
  }, [crews]);

  useEffect(() => {
    if (success) {
      navigate(`/ratings/${ratingID}/races/${raceID}`);
    }
  }, [success, navigate, ratingID, raceID]);

  const handleSearchChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setSearch(e.target.value);
  };

  const handleSelectYacht = (yachtNumber: number) => {
    setFinisherList((prevList) => [...prevList, yachtNumber]);
    setAvailableYachts((prevYachts) =>
        prevYachts.filter((num) => num !== yachtNumber)
    );
    setSearch("");
  };

  const handleSubmit = () => {
    const input: FinishInput = { finisherList };
    finishProcedure(input);
  };

  const handleDragEnd = (event: any) => {
    const { active, over } = event;

    if (active.id !== over.id) {
      setFinisherList((items) => {
        const oldIndex = items.indexOf(active.id);
        const newIndex = items.indexOf(over.id);
        return arrayMove(items, oldIndex, newIndex);
      });
    }
  };

  if (!ratingID || !raceID)
    return <p className="error">Ошибка: отсутствует ID рейтинга или гонки</p>;
  if (ratingLoading || raceLoading || crewsLoading)
    return <p>Загрузка данных...</p>;
  if (!ratingInfo || !raceInfo)
    return <p className="error">Ошибка загрузки данных</p>;

  return (
      <div className="finish-procedure">
        <h1 className={"headerH1"}>{ratingInfo.Name}</h1>
        <h2 className={"headerH2"}>Гонка {raceInfo.number}</h2>
        <h3 className={"headerH3"}>Финишная процедура</h3>

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
            {availableYachts
                .filter((num) => num.toString().includes(search))
                .map((num) => (
                    <div
                        key={num}
                        className="yacht-button"
                        onClick={() => handleSelectYacht(num)}
                    >
                      {num}
                    </div>
                ))}
          </div>

          <div className="finish-list">
            <h3>Номера яхт в порядке финиша</h3>
            <DndContext collisionDetection={closestCenter} onDragEnd={handleDragEnd}>
              <SortableContext
                  items={finisherList}
                  strategy={verticalListSortingStrategy}
              >
                {finisherList.map((num) => (
                    <SortableItem key={num} yachtNumber={num} />
                ))}
              </SortableContext>
            </DndContext>
          </div>
        </div>

        <button
            onClick={handleSubmit}
            disabled={loading}
            className="submit-button"
        >
          Завершить финишную процедуру
        </button>
        {error && <p className="error">Ошибка: {error}</p>}
        {success && <p className="success">{success}</p>}
      </div>
  );
};

export default FinishProcedure;
