package api

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"test10/pkg/utils"
	"test10/service"
)

//新增一条备忘录
func CreateTask(c *gin.Context) {
	var createTask service.CreateTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createTask); err == nil {
		res := createTask.Create(claim.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}
}

//展示一条备忘录
func ShowTask(c *gin.Context) {
	var showTask service.ShowTaskService
	//claim,_:=utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showTask); err == nil {
		res := showTask.Show(c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}
}

//展示所有备忘录
func ListTask(c *gin.Context) {
	var listTask service.ListTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listTask); err == nil {
		res := listTask.List(claim.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}
}

//更新备忘录
func UpdateTask(c *gin.Context) {
	var updateTask service.UpdateTaskService
	//claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&updateTask); err == nil {
		res := updateTask.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}
}

//查询备忘录
func SearchTask(c *gin.Context) {
	var searchTask service.SearchTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&searchTask); err == nil {
		res := searchTask.Search(claim.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}
}

//删除备忘录
func DeleteTask(c *gin.Context) {
	var deleteTask service.DeleteTaskService
	//claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteTask); err == nil {
		res := deleteTask.Delete(c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}
}
