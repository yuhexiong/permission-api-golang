package controller

import (
	"permission-api/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateTaskOpts struct {
	ToUserId     *primitive.ObjectID `json:"toUserId" binding:"omitempty" example:"abd1234"`    // 被指派帳號id
	Deadline     string              `json:"deadline" binding:"omitempty" example:"2024-01-03"` // 結束日期
	PriorityType *model.PriorityType `json:"priorityType" binding:"omitempty" example:"HIGH"`   // 任務優先度 HIGH=高, MEDIUM=中, LOW=低
	Title        string              `json:"title" binding:"omitempty" example:"完成表格"`          // 標題
	Content      string              `json:"content" binding:"omitempty" example:"將1到10欄的內容補齊"` // 內容
}

// 建立新任務
func CreateTask(opts CreateTaskOpts, fromUserId *primitive.ObjectID, result *model.Task) error {
	priorityType := model.PriorityTypeMedium
	if opts.PriorityType != nil {
		priorityType = *opts.PriorityType
	}

	task := model.Task{
		FromUserId:   fromUserId,
		ToUserId:     opts.ToUserId,
		Deadline:     opts.Deadline,
		PriorityType: priorityType,
		Title:        opts.Title,
		Content:      opts.Content,
	}

	return model.Insert(model.TaskCollName, &task, &result)
}
