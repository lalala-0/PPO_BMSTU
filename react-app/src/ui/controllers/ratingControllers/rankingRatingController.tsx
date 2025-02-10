// src/services/RatingService.ts
import { Rating, RankingResponse } from "../../models/ratingModel";
import api from "../api"; // Импортируем функцию для обработки ошибок

export const fetchRatingInfo = async (ratingID: string): Promise<Rating> => {
  const response = await api.get<Rating>(`/ratings/${ratingID}`);
  return response.data;
};

export const fetchRankingData = async (
  ratingID: string,
): Promise<RankingResponse> => {
  const response = await api.get<RankingResponse>(
    `/ratings/${ratingID}/rankings`,
  );
  return response.data;
};
