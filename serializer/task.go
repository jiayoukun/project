package serializer

import "test10/model"

type Task struct {
	Id        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	View      int64  `json:"view"`
	Status    int    `json:"status"` // 0未完成 1是已完成
	CreateAt  int64  `json:"create_at"`
	StartTime int64  `json:"start_time"` //备忘录开始时间
	EndTime   int64  `json:"end_time"`   //备忘录完成时间
}

func BuildTask(item model.Task) Task {
	return Task{
		Id:        item.ID,
		Title:     item.Title,
		Content:   item.Content,
		Status:    item.Status,
		StartTime: item.StartTime,
		EndTime:   item.EndTime,
	}
}

func BuildTasks(item model.Task) Task {
	return Task{
		Id:        item.ID,
		Title:     item.Title,
		Content:   item.Content,
		Status:    item.Status,
		StartTime: item.StartTime,
		EndTime:   item.EndTime,
	}
}
