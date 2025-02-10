import { useState, useEffect } from "react";
import axios from "axios";
import { CrewFormData } from "../../models/crewModel";
import { useParams } from "react-router-dom";
import { API_URL } from "../../config";
import { handleError } from "../errorHandler";
import api from "../api";

export const useGetCrew = () => {
  const { ratingID, crewID } = useParams<{
    ratingID: string;
    crewID: string;
  }>();
  const [crewInfo, setCrewInfo] = useState<CrewFormData | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!ratingID || !crewID) {
      setError("Недостаточно данных для загрузки");
      setLoading(false);
      return;
    }

    const fetchCrewData = async () => {
      try {
        const response = await api.get<CrewFormData>(
          `/ratings/${ratingID}/crews/${crewID}`,
        );
        setCrewInfo(response.data);
        setLoading(false);
      } catch (err: any) {
        handleError(err, setError);
      }
    };

    fetchCrewData();
  }, [ratingID, crewID]);

  return {
    crewInfo,
    loading,
    error,
  };
};
