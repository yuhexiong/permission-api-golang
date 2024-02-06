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

type BaseData struct {
	Status    *uint8     `bson:"status,omitempty" json:"status" example:"0"` // 0: 正常, 9: 刪除
	CreatedAt *time.Time `bson:"createdAt,omitempty" json:"createdAt" example:"2022-03-21T10:30:17.711Z"`
	UpdatedAt *time.Time `bson:"updatedAt,omitempty" json:"updatedAt" example:"2022-03-21T10:30:17.711Z"`
}

func structToBsonM(data interface{}) (bson.M, error) {
	document, err := bson.Marshal(data)
	if err != nil {
		return nil, err
	}

	var bsonM bson.M
	err = bson.Unmarshal(document, &bsonM)
	if err != nil {
		return nil, err
	}

	return bsonM, nil
}

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

	if !util.IsInterfaceValueNil(result) && cursor.Next(context.Background()) {
		c, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := cursor.All(c, result); err != nil {
			util.RedLog("Find - decode err: %s", err.Error())
			return err
		}
	}
	return nil
}

// 依id尋找一位
func Get(collectionName string, objectId *primitive.ObjectID, result interface{}) error {
	util.BlueLog("Get(%s) objectId(%+v)", collectionName, objectId)

	filter := bson.D{{Key: "_id", Value: objectId}}
	return GetByFilter(collectionName, filter, result)
}

// 尋找一位
func GetByFilter(collectionName string, filter interface{}, result interface{}) error {
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

	if !util.IsInterfaceValueNil(result) {
		if err := cursor.Decode(result); err != nil {
			util.RedLog("Get - decode err:  %s", err.Error())
			return err
		}
	}

	return nil
}

// 新增
func Insert(collectionName string, rawData interface{}, result interface{}) error {
	util.BlueLog("Insert(%s) data(%+v)", collectionName, rawData)

	data, err := structToBsonM(rawData)
	if err != nil {
		return err
	}
	data["createdAt"] = time.Now()
	data["updatedAt"] = time.Now()
	data["status"] = 0 // 預設新增的資料就是啟用的

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	insertResult, err := config.GetCollection(config.GetDB(), collectionName).InsertOne(c, data)
	if err != nil {
		util.RedLog("Insert err: %s", err.Error())
		return err
	}

	if insertResult != nil && !util.IsInterfaceValueNil(result) {
		return GetByFilter(collectionName, bson.D{{Key: "_id", Value: insertResult.InsertedID}}, result)
	}

	return err
}

// 更新
func Update(collectionName string, objectId *primitive.ObjectID, rawData interface{}) error {
	util.BlueLog("Update(%s) objectId(%+v) data(%+v)", collectionName, objectId, rawData)

	data, err := structToBsonM(rawData)
	if err != nil {
		return err
	}
	data["updatedAt"] = time.Now()

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = config.GetCollection(config.GetDB(), collectionName).UpdateMany(c, bson.M{"_id": objectId}, bson.M{"$set": data})
	if err != nil {
		util.RedLog("Update err: %s", err.Error())
		return err
	}

	return nil
}

// 啟用
func Enable(collectionName string, objectId *primitive.ObjectID) error {
	util.BlueLog("Enable(%s) objectId(%+v)", collectionName, objectId)

	data := bson.M{
		"updatedAt": time.Now(),
		"status":    NormalStatus,
	}

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := config.GetCollection(config.GetDB(), collectionName).UpdateMany(c, bson.M{"_id": objectId}, bson.D{{Key: "$set", Value: data}})
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
		_, err = config.GetCollection(config.GetDB(), collectionName).DeleteMany(c, filter)
	} else {
		_, err = config.GetCollection(config.GetDB(), collectionName).UpdateMany(c,
			bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: DeleteStatus}}}},
			filter,
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
		_, err = config.GetCollection(config.GetDB(), collectionName).DeleteMany(c, bson.M{"_id": objectId})
	} else {
		_, err = config.GetCollection(config.GetDB(), collectionName).UpdateMany(c, bson.M{"_id": objectId}, bson.M{"$set": bson.M{"status": DeleteStatus}})
	}

	if err != nil {
		util.RedLog("Delete err: %s", err.Error())
		return err
	}

	return nil
}
