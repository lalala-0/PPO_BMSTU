import { useState, useEffect } from "react";
import axios from "axios";
import { Rating } from "../../models/ratingModel";

export const useGetRatingsController = () => {
  const [ratings, setRatings] = useState<Rating[]>([]);
  const [error, setError] = useState<string | null>(null);

  const getRatings = (filters: Record<string, string>) => {
    const params = new URLSearchParams(filters).toString();
    const url = params ? `/api/ratings?${params}` : "/api/ratings";

    axios
      .get(url)
      .then((response) => {
        setRatings(response.data);
      })
      .catch((err) => {
        setError("Ошибка при получении данных");
      });
  };

  useEffect(() => {
    getRatings({});
  }, []);

  return { ratings, error, setRatings };
};

export const filterAndGroupRatings = (
  ratings: Rating[],
  filters: { name: string; class: string; blowoutCnt: string },
): Rating[] => {
  const filteredRatings = ratings.filter((rating) => {
    return (
      (filters.name === "" ||
        rating.Name.toLowerCase().includes(filters.name.toLowerCase())) &&
      (filters.class === "" ||
        rating.Class.toLowerCase() === filters.class.toLowerCase()) &&
      (filters.blowoutCnt === "" ||
        rating.BlowoutCnt.toString() === filters.blowoutCnt)
    );
  });

  // Группировка рейтингов по имени
  return filteredRatings.sort((a, b) => a.Name.localeCompare(b.Name));
};
