package service

import (
	"github.com/jinzhu/gorm"
	"test10/model"
	"test10/pkg/utils"
	"test10/serializer"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	PassWord string `form:"password" json:"password" binding:"required,min=5,max=16"`
}

func (service *UserService) Register() serializer.Response {
	var user model.User
	var count int
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).First(&user).Count(&count)
	if count == 1 {
		return serializer.Response{
			Status: 400,
			Msg:    "该账号已存在",
		}
	}
	user.UserName = service.UserName
	//密码加密

	if err := user.SetPassWord(service.PassWord); err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "密码错误",
		}
	}
	//创建用户

	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "数据库操作错误",
		}
	}

	return serializer.Response{
		Status: 200,
		Msg:    "用户注册成功",
	}
}

func (service *UserService) Login() serializer.Response {
	var user model.User
	if err := model.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: 400,
				Msg:    "用户不存在",
			}
		}
		return serializer.Response{
			Status: 500,
			Msg:    "数据库错误",
		}
	}
	if user.CheckPassWord(service.PassWord) == false {
		return serializer.Response{
			Status: 400,
			Msg:    "密码错误",
		}
	}

	//登陆成功  签发token
	token, err := utils.GenerateToken(user.ID, service.UserName, service.PassWord)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "token签发错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    "登陆成功",
	}
}
