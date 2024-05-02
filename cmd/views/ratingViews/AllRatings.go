package ratingViews

import (
	"PPO_BMSTU/cmd/modelTables"
	"PPO_BMSTU/internal/registry"
)

func AllRatings(services registry.Services) error {
	ratings, err := services.RatingService.GetAllRatings()
	if err != nil {
		return err
	}
	return modelTables.Ratings(ratings)
}
