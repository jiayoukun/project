package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"test10/api"
	"test10/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("api/v1")
	{
		//用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{

			authed.POST("user/task", api.CreateTask)
			authed.GET("user/task/:id", api.ShowTask)
			authed.POST("user/tasks", api.ListTask)
			authed.PUT("user/task/:id", api.UpdateTask)
			authed.POST("user/task/search", api.SearchTask)
			authed.DELETE("user/task/delete/:id", api.DeleteTask)
		}
	}

	return r
}
