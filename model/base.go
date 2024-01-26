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
	util.BlueLog("Find(%s) filter(%+v)", collectionName, filter)
	var pipeline = mongo.Pipeline{}
	if filter != nil {
		pipeline = append(pipeline, bson.D{{Key: "$match", Value: filter}})
	}

	return FindByPipeline(collectionName, pipeline, result)
}

// 使用 pipeline 尋找
func FindByPipeline(collectionName string, pipeline mongo.Pipeline, result interface{}) error {
	util.BlueLog("Find(%s) pipeline(%+v)", collectionName, pipeline)

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coll := config.GetCollection(config.GetDB(), collectionName)
	cursor, err := coll.Aggregate(c, pipeline)

	if err != nil {
		util.RedLog("Find err: %s", err.Error())
		return err
	}

	if result != nil && cursor.Next(context.Background()) {
		c, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := cursor.All(c, result); err != nil {
			util.RedLog("Find - decode err: %s", err.Error())
			return err
		}
	}
	return nil
}

// 尋找一位
func Get(collectionName string, filter interface{}, result interface{}) error {
	util.BlueLog("Get(%s) filter(%+v)", collectionName, filter)

	var pipeline = mongo.Pipeline{}
	if filter != nil {
		pipeline = append(pipeline, bson.D{{Key: "$match", Value: filter}})
	}

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coll := config.GetCollection(config.GetDB(), collectionName)
	cursor, err := coll.Aggregate(c, pipeline)

	if err != nil {
		util.RedLog("Get err: %s", err.Error())
		return err
	}

	if !cursor.TryNext(context.TODO()) {
		return mongo.ErrNoDocuments
	}

	if result != nil {
		if err := cursor.Decode(result); err != nil {
			util.RedLog("Get - decode err:  %s", err.Error())
			return err
		}
	}

	return nil
}

// 啟用
func Insert(collectionName string, data interface{}, result interface{}) error {
	util.BlueLog("Insert(%s) data(%+v)", collectionName, data)

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	insertResult, err := config.GetCollection(config.GetDB(), collectionName).InsertOne(c, data)
	if err != nil {
		util.RedLog("Insert err: %s", err.Error())
		return err
	}

	if err == nil && insertResult != nil && result != nil {
		return Get(collectionName, bson.D{{Key: "_id", Value: insertResult.InsertedID}}, result)
	}

	return nil
}

// 啟用
func Enable(collectionName string, objectId *primitive.ObjectID) error {
	util.BlueLog("Enable(%s) objectId(%+v)", collectionName, objectId)

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := config.GetCollection(config.GetDB(), collectionName).UpdateMany(c, bson.M{"id": objectId}, bson.D{{Key: "$status", Value: NormalStatus}})
	if err != nil {
		util.RedLog("Enable err: %s", err.Error())
		return err
	}

	return nil
}

// 依條件刪除
func DeleteByFilter(collectionName string, filter interface{}, forceDelete bool) error {
	util.BlueLog("Delete(%s) filter(%+v)", collectionName, filter)

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	if forceDelete {
		_, err = config.GetCollection(config.GetDB(), collectionName).DeleteMany(c, mongo.Pipeline{bson.D{{Key: "$match", Value: filter}}})
	} else {
		_, err = config.GetCollection(config.GetDB(), collectionName).UpdateMany(c,
			bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: DeleteStatus}}}},
			mongo.Pipeline{bson.D{{Key: "$match", Value: filter}}},
		)
	}

	if err != nil {
		util.RedLog("Delete err: %s", err.Error())
		return err
	}

	return nil
}

// 依id刪除
func Delete(collectionName string, objectId *primitive.ObjectID, forceDelete bool) error {
	util.BlueLog("Delete(%s) objectId(%+v)", collectionName, objectId)

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	if forceDelete {
		_, err = config.GetCollection(config.GetDB(), collectionName).DeleteMany(c, bson.M{"id": objectId})
	} else {
		_, err = config.GetCollection(config.GetDB(), collectionName).UpdateMany(c, bson.M{"id": objectId}, bson.D{{Key: "$status", Value: DeleteStatus}})
	}

	if err != nil {
		util.RedLog("Delete err: %s", err.Error())
		return err
	}

	return nil
}
