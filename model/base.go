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

// 尋找
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

// 啟用
func Insert(collectionName string, data interface{}, result interface{}) error {
	util.GreenLog("Insert(%s) data(%+v)", collectionName, data)

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := config.GetCollection(config.GetDB(), collectionName).InsertOne(c, data)
	if err != nil {
		util.RedLog("Insert err: %s", err.Error())
		return err
	}

	return nil
}

// 啟用
func Enable(collectionName string, objectId *primitive.ObjectID) error {
	util.GreenLog("Enable(%s) objectId(%+v)", collectionName, objectId)

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := config.GetCollection(config.GetDB(), collectionName).UpdateMany(c, bson.M{"id": objectId}, bson.D{{Key: "$status", Value: NormalStatus}})
	if err != nil {
		util.RedLog("Enable err: %s", err.Error())
		return err
	}

	return nil
}

// 刪除
func Delete(collectionName string, objectId *primitive.ObjectID, forceDelete bool) error {
	util.GreenLog("Delete(%s) objectId(%+v)", collectionName, objectId)

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	if forceDelete {
		_, err = config.GetCollection(config.GetDB(), collectionName).DeleteMany(c, bson.M{"id": objectId})
	} else {
		_, err = config.GetCollection(config.GetDB(), collectionName).UpdateMany(c, bson.M{"id": objectId}, bson.D{{Key: "$status", Value: NormalStatus}})
	}

	if err != nil {
		util.RedLog("Delete err: %s", err.Error())
		return err
	}

	return nil
}
