import { useState, useEffect } from "react";
import axios from "axios";
import { RaceFormData } from "../../models/raceModel";
import { useParams } from "react-router-dom";
import { API_URL } from "../../config";
import { handleError } from "../errorHandler";
import api from "../api"; // Импортируем функцию для обработки ошибок

export const useGetRace = () => {
  const { ratingID, raceID } = useParams<{
    ratingID: string;
    raceID: string;
  }>();
  const [raceInfo, setRaceInfo] = useState<RaceFormData | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!ratingID || !raceID) {
      setError("Недостаточно данных для загрузки");
      setLoading(false);
      return;
    }

    const fetchRaceData = async () => {
      try {
        const response = await api.get<RaceFormData>(
          `/ratings/${ratingID}/races/${raceID}`,
        );
        setRaceInfo(response.data);
        setLoading(false);
      } catch (err: any) {
        handleError(err, setError);
      }
    };

    fetchRaceData();
  }, [ratingID, raceID]);

  return {
    raceInfo,
    loading,
    error,
  };
};
