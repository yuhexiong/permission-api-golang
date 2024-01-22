package model

import (
	"context"
	"permission-api/config"
	"permission-api/util"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	NormalStatus uint8 = iota // 0: 正常
	DeleteStatus uint8 = 9    // 9: 刪除
)

func Find(collectionName string, filter interface{}, result interface{}) error {
	util.GreenLog("Find(%s) filter(%+v)", collectionName, filter)
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

func Delete(collectionName string, objectId *primitive.ObjectID, forceDelete bool) error {
	util.GreenLog("Delete(%s) objectId(%+v)", collectionName, objectId)

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if forceDelete {
		config.GetCollection(config.GetDB(), collectionName).DeleteMany(c, bson.M{"id": objectId})
	} else {
		config.GetCollection(config.GetDB(), collectionName).UpdateMany(c, bson.M{"id": objectId}, bson.D{{Key: "$status", Value: NormalStatus}})
	}

	return nil
}
