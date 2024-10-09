package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func getNextSequence(db *mongo.Database, sequenceName string) (int, error) {
	collection := db.Collection("counters")

	filter := bson.M{"_id": sequenceName}
	update := bson.M{"$inc": bson.M{"seq": 1}}

	var result struct {
		Seq int `bson:"seq"`
	}

	err := collection.FindOneAndUpdate(context.Background(), filter, update).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			_, err := collection.InsertOne(context.Background(), bson.M{"_id": sequenceName, "seq": 1})
			if err != nil {
				return 0, err
			}
			return 1, nil
		}
		return 0, err
	}

	return result.Seq, nil
}
