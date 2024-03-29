package controller

import (
	"errors"
	"fmt"
	"permission-api/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateTaskOpts struct {
	ToUserOId    *primitive.ObjectID `json:"toUserOId" binding:"omitempty" example:"abd1234"`   // 被指派帳號id
	Deadline     string              `json:"deadline" binding:"omitempty" example:"2024-01-03"` // 結束日期
	PriorityType *model.PriorityType `json:"priorityType" binding:"omitempty" example:"HIGH"`   // 任務優先度 HIGH=高, MEDIUM=中, LOW=低
	Title        string              `json:"title" binding:"omitempty" example:"完成表格"`          // 標題
	Content      string              `json:"content" binding:"omitempty" example:"將1到10欄的內容補齊"` // 內容
}

// 建立任務
func CreateTask(opts CreateTaskOpts, fromUserId *primitive.ObjectID, result *model.Task) error {
	priorityType := model.PriorityTypeMedium
	if opts.PriorityType != nil {
		priorityType = *opts.PriorityType
	}

	task := model.Task{
		FromUserOId:  fromUserId,
		ToUserOId:    opts.ToUserOId,
		Deadline:     opts.Deadline,
		PriorityType: priorityType,
		Title:        opts.Title,
		Content:      opts.Content,
	}

	return model.Insert(model.TaskCollName, &task, &result)
}

type FindTaskOpts struct {
	FromUserOId *primitive.ObjectID `json:"fromUserOId" bson:"fromUserOId,omitempty" binding:"omitempty" example:"623850247ea4cca15cd55303"` // 指派帳號id
	ToUserOId   *primitive.ObjectID `json:"toUserOId" bson:"toUserOId,omitempty" binding:"omitempty" example:"623850247ea4cca15cd55303"`     // 被指派帳號id
}

// 取得任務
func FindTask(opts FindTaskOpts, result *[]*model.Task) error {
	filter := bson.D{}

	if opts.FromUserOId != nil {
		filter = append(filter, bson.E{Key: "fromUserOId", Value: opts.FromUserOId})
	}

	if opts.ToUserOId != nil {
		filter = append(filter, bson.E{Key: "toUserOId", Value: opts.ToUserOId})
	}

	return model.Find(model.TaskCollName, filter, result)
}

// 更新任務進度
func UpdateTaskProgressType(objectId *primitive.ObjectID, userOId *primitive.ObjectID, progressType *model.ProgressType) error {
	task := model.Task{}
	if err := model.Get(model.TaskCollName, objectId, &task); err != nil {
		return err
	}

	if *task.ToUserOId != *userOId && *task.FromUserOId != *userOId {
		return errors.New("task not assigned to this user or created by this user")
	}

	if *progressType == model.DONE && (task.Checked == nil || !*task.Checked) {
		return errors.New("task cannot be moved to DONE before checked")
	}

	if *progressType == model.DELETE && *task.FromUserOId != *userOId {
		return errors.New("only creator can delete the task")
	}

	if err := model.Update(model.TaskCollName, objectId, bson.D{{Key: "progressType", Value: *progressType}}); err != nil {
		return err
	}

	// 如果移到測試, 就新增通知給指派者確認驗收
	if *progressType == model.TEST {
		err := createNotification(createNotificationOpts{
			ToUserOId: task.FromUserOId,
			Content:   fmt.Sprintf("Task '%s' has been completed and is ready for checking.", task.Title),
		}, nil)

		return err
	}

	return nil
}

// 更新任務驗收
func CheckTask(objectId *primitive.ObjectID, userOId *primitive.ObjectID, checked *bool) error {
	task := model.Task{}
	if err := model.Get(model.TaskCollName, objectId, &task); err != nil {
		return err
	}

	if *task.FromUserOId != *userOId {
		return errors.New("task not created by this user")
	}

	if err := model.Update(model.TaskCollName, objectId, bson.D{{Key: "checked", Value: *checked}}); err != nil {
		return err
	}

	// 如果驗收完成, 通知被指派者已完成
	if *checked {
		err := createNotification(createNotificationOpts{
			ToUserOId: task.ToUserOId,
			Content:   fmt.Sprintf("Task '%s' has been checked successfully.", task.Title),
		}, nil)

		return err
	}

	return nil
}

// 刪除任務
func DeleteTask(objectId *primitive.ObjectID, userOId *primitive.ObjectID) error {
	task := model.Task{}
	if err := model.Get(model.TaskCollName, objectId, &task); err != nil {
		return err
	}

	if *task.FromUserOId != *userOId {
		return errors.New("task not created by this user")
	}

	return model.Delete(model.TaskCollName, objectId, false)
}
