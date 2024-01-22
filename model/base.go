package model

import (
	"context"
	"permission-api/config"
	"permission-api/util"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Find(collectionName string, filter interface{}, result interface{}) error {
	util.GreenLog("doFind(%s) filter(%+v)", collectionName, filter)
	var pipeline = mongo.Pipeline{}
	if filter != nil {
		pipeline = append(pipeline, bson.D{{Key: "$match", Value: filter}})
	}

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coll := config.GetCollection(config.GetDB(), collectionName)
	cursor, err := coll.Aggregate(c, pipeline)

	if err != nil {
		util.RedLog("Find err: %s", err.Error())
		return err
	}

	if result != nil && cursor.TryNext(context.Background()) {
		c, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := cursor.All(c, result); err != nil {
			util.RedLog("Find - decode err:  %s", err.Error())
			return err
		}
	}
	return nil
}
