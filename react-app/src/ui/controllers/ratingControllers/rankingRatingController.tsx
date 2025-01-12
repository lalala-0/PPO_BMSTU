// src/services/RatingService.ts
import axios from "axios";
import { Rating, RankingResponse } from "../../models/ratingModel";

const BASE_URL = "/api/ratings";

export const fetchRatingInfo = async (ratingID: string): Promise<Rating> => {
  const response = await axios.get<Rating>(`${BASE_URL}/${ratingID}`);
  return response.data;
};

export const fetchRankingData = async (
  ratingID: string,
): Promise<RankingResponse> => {
  const response = await axios.get<RankingResponse>(
    `${BASE_URL}/${ratingID}/rankings`,
  );
  return response.data;
};
