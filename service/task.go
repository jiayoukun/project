package service

import (
	"test10/model"
	"test10/serializer"
	"time"
)

type CreateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` //0是未做 1是已做
}

type UpdateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` //0是未做 1是已做
}
type ShowTaskService struct {
}
type ListTaskService struct {
	PageNumber int `json:"page_number" form:"page_number"`
	PageMax    int `json:"page_max" form:"page_max"`
}

type SearchTaskService struct {
	Info       string `json:"info" form:"info"`
	PageNumber int    `json:"page_number" form:"page_number"`
	PageMax    int    `json:"page_max" form:"page_max"`
}

type DeleteTaskService struct {
}

func (service *CreateTaskService) Create(id uint) serializer.Response {
	var user model.User
	var code int
	model.DB.First(&user, id)
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Status:    0,
		StartTime: time.Now().Unix(),
		EndTime:   0,
	}
	err := model.DB.Create(&task).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "创建备忘录失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "创建备忘录成功",
	}
}

func (service *ShowTaskService) Show(tid string) serializer.Response {
	var task model.Task
	code := 200
	err := model.DB.First(&task, tid).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "查询失败",
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
		Msg:    "查询成功",
	}
}

func (service *ListTaskService) List(uid uint) serializer.Response {
	var tasks []model.Task
	count := 0
	if service.PageMax == 0 {
		service.PageMax = 15
	}
	model.DB.Model(&model.Task{}).Preload("User").Where("Uid=?", uid).Count(&count).Limit(service.PageMax).Offset((service.PageNumber - 1) * service.PageMax).Find(&tasks)
	return serializer.BuildListResponse(tasks, uint(count))
}

func (service *UpdateTaskService) Update(tid string) serializer.Response {
	var task model.Task
	model.DB.First(&task, tid)
	task.Content = service.Content
	task.Title = service.Title
	task.Status = service.Status
	model.DB.Save(&task)
	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildTask(task),
		Msg:    "修改成功",
	}
}

func (service *SearchTaskService) Search(uid uint) serializer.Response {
	var tasks []model.Task
	count := 0
	if service.PageMax == 0 {
		service.PageMax = 15
	}
	err := model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).Where("title LIKE ? OR content LIKE ?", "%"+service.Info+"%", "%"+service.Info+"%").Count(&count).
		Limit(service.PageMax).Offset((service.PageNumber - 1) * service.PageMax).Find(&tasks).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "数据错误",
		}
	} else {

		return serializer.BuildListResponse(tasks, uint(count))
	}
}

func (service *DeleteTaskService) Delete(tid string) serializer.Response {
	var task model.Task
	model.DB.First(&task, tid)
	err := model.DB.Delete(&task, tid).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "删除失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "删除成功",
	}
}
