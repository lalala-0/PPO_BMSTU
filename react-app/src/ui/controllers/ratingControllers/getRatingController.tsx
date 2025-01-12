// src/controllers/RatingController.ts
import { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import {
  Rating,
  RatingInput,
  RatingTableLine,
  RaceInfo,
} from "../../models/ratingModel";
import { fetchRatingInfo, fetchRankingData } from "./rankingRatingController";
import { useUpdateRatingController } from "./updateRatingController";

export const useGetRating = () => {
  const { ratingID } = useParams<{ ratingID: string }>();
  const navigate = useNavigate();

  const { success, handleUpdate } = useUpdateRatingController();

  const [ratingInfo, setRatingInfo] = useState<Rating | null>(null);
  const [ratingInfoEditable, setRatingInfoEditable] =
    useState<RatingInput | null>(null);
  const [rankingTable, setRankingTable] = useState<RatingTableLine[]>([]);
  const [races, setRaces] = useState<RaceInfo[]>([]); // Указан тип для races
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    if (ratingID) {
      const fetchData = async () => {
        try {
          const ratingData = await fetchRatingInfo(ratingID);
          setRatingInfo(ratingData);
          setRatingInfoEditable({
            name: ratingData.Name,
            class: Number(ratingData.Class), // Приведение к числу
            blowout_cnt: ratingData.BlowoutCnt,
          });

          const rankingData = await fetchRankingData(ratingID);
          setRankingTable(rankingData.RankingTable);
          setRaces(rankingData.Races);
        } catch (err) {
          console.error("Ошибка:", err);
        } finally {
          setLoading(false);
        }
      };

      fetchData();
    }
  }, [ratingID]);

  const handleInputChange = (
    e: React.ChangeEvent<HTMLInputElement>,
    field: keyof RatingInput,
  ) => {
    if (ratingInfoEditable) {
      setRatingInfoEditable({
        ...ratingInfoEditable,
        [field]: e.target.value,
      });
    }
  };

  const handleSaveUpdate = async () => {
    if (ratingInfoEditable) {
      const { name, class: ratingClass, blowout_cnt } = ratingInfoEditable;
      const updatedData: RatingInput = {
        name,
        class: ratingClass, // Используем class как есть
        blowout_cnt,
      };

      try {
        if (ratingID) {
          await handleUpdate(ratingID, updatedData);
          setRatingInfoEditable(null); // Очистка данных после сохранения
        }
      } catch (error) {
        console.error("Ошибка при обновлении рейтинга", error);
      }
    }
  };

  const handleGoBack = () => {
    navigate("/ratings");
  };

  return {
    ratingInfo,
    ratingInfoEditable,
    rankingTable,
    races,
    loading,
    success,
    handleInputChange,
    handleSaveUpdate,
    handleGoBack,
  };
};
