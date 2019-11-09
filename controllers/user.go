package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"github.com/pengchujin/subscribe_go/models"
	"github.com/pengchujin/subscribe_go/util"
	"github.com/pengchujin/subscribe_go/database"
)

type User struct {
	Basic
}

type CreateRequest struct {
	UserName	string `form:"username" json:"username" binding:"required"`
	Email		string `form:"email" json:"email" binding:"required"`
	Password	string `form:"password" json:"password" binding:"required"`
}

type GetRequest struct {
	Email		string `form:"email" json:"email" binding:"required"`
	Password	string `form:"password" json:"password" binding:"required"`
}

func (u *User) Store(c *gin.Context) {
	var request CreateRequest
	if err := c.ShouldBind(&request); err == nil {
		var count int
		database.DB.Model(&models.User{}).Where("Email = ?", request.Email).Count(&count)
		if (count > 0) {
			u.JsonFail(c, http.StatusBadRequest, "该邮箱已注册")
			return
		}
		EncryptedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
		if err != nil {
			u.JsonFail(c, http.StatusBadRequest, "注册出现错误")
			return
		}
		user := models.User {
			Email: request.Email,
			UserName: request.UserName,
			EncryptedPassword: string(EncryptedPassword),
		}
		if err := database.DB.Create(&user).Error; err != nil {
			u.JsonFail(c, http.StatusBadRequest, err.Error())
			return
		}
		u.JsonSuccess(c, http.StatusCreated, gin.H{"message": "创建成功"})
	} else {
		u.JsonFail(c, http.StatusBadRequest, err.Error())
	}
}


func (u *User) Get(c *gin.Context) {
	var request GetRequest
	if err := c.ShouldBind(&request); err == nil {
		var user models.User
		if err := database.DB.Model(&models.User{}).Where("Email = ?", request.Email).Find(&user).Error; err != nil {
			u.JsonFail(c, http.StatusBadRequest, "用户名或密码错误")
			return
		}
	  if err := bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword),[]byte(request.Password)); err != nil {
			u.JsonFail(c, http.StatusBadRequest, "用户名或密码错误")
			return
		} 
		 token, err := util.GenerateToken(user.Email, user.UserName, user.UUID.String())
			if err != nil {
				u.JsonFail(c, http.StatusBadRequest, "用户名或密码错误")
				return
			}
			u.JsonSuccess(c, http.StatusCreated, gin.H{
				"jwt": token,
				"email": user.Email,
				"uuid": user.UUID,
			})
	} else {
		u.JsonFail(c, http.StatusBadRequest, err.Error())
	}
}