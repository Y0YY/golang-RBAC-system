package controller

import (
	. "Assignment/model"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

// 得到所有用户
func (uc *UserController) GetAllUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, GetAllUsers())
}

// 得到一个用户的所有角色
func (uc *UserController) GetAllRolesOfUser(ctx *gin.Context) {
	userName := ctx.Param("userName")
	u, err1 := FindUserByName(userName)
	if err1 != nil {
		ctx.JSON(403, gin.H{
			"err":      err1.Error(),
			"userName": userName,
		})
	} else {
		rs, err2 := GetAllRolesOfUser(*u)
		if err2 != nil {
			ctx.JSON(403, gin.H{
				"err":      err2.Error(),
				"userName": userName,
			})
		} else {
			ctx.JSON(http.StatusOK, rs)
		}
	}

}

// 得到一个用户的所有权限
func (uc *UserController) GetAllPermsOfUser(ctx *gin.Context) {
	userName := ctx.Param("userName")
	u, err1 := FindUserByName(userName)

	if err1 != nil {
		ctx.JSON(403, gin.H{
			"err": err1.Error(),
		})
	} else {
		rs, err2 := GetAllPermsOfUser(*u)
		if err2 != nil {
			ctx.JSON(403, gin.H{
				"err": err2.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, rs)
		}
	}
}

// 查询用户是否有指定权限
func (uc *UserController) IsPrmitted(ctx *gin.Context) {
	userName := ctx.Param("userName")
	permName := ctx.Param("permName")
	u, err1 := FindUserByName(userName)
	p, err2 := FindPermByName(permName)
	if err1 != nil {
		ctx.JSON(403, gin.H{
			"err": err1.Error(),
		})
	} else if err2 != nil {
		ctx.JSON(403, gin.H{
			"err": err2.Error(),
		})
	} else if IsPrmitted(*u, *p) {
		ctx.JSON(http.StatusOK, gin.H{
			"result": "success",
		})
	} else {
		ctx.JSON(403, gin.H{
			"error": "access denied",
		})
	}
}

// 创建用户
func (uc *UserController) Create(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	var NewUser User
	_ = json.Unmarshal(data, &NewUser)
	err := CreatUser(NewUser)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": "success",
		})
	} else {
		ctx.JSON(403, gin.H{
			"error": err.Error(),
		})
	}
}

// 为一个用户添加角色
func (uc *UserController) AddRole(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	var body map[string]string
	_ = json.Unmarshal(data, &body)
	userName := body["userName"]
	roleName := body["roleName"]
	u, err1 := FindUserByName(userName)
	r, err2 := FindRoleByName(roleName)
	if err1 != nil {
		ctx.JSON(403, gin.H{
			"userName": userName,
			"err":      err1.Error(),
		})
	} else if err2 != nil {
		ctx.JSON(403, gin.H{
			"roleName": roleName,
			"err":      err2.Error(),
		})
	} else {
		err := AddRole(*u, *r)
		if err != nil {
			ctx.JSON(403, gin.H{
				"err": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"result": "success",
			})
		}
	}
}

// 通过ID删除一个用户
func (uc *UserController) DeleteById(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	var body map[string]string
	_ = json.Unmarshal(data, &body)
	id := body["userId"]
	err := DeleteUserById(id)
	if err != nil {
		ctx.JSON(403, gin.H{
			"err": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"result": "delete success",
		})
	}
}

// 通过名字删除一个用户
func (uc *UserController) DeleteByName(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	var body map[string]string
	_ = json.Unmarshal(data, &body)
	userName := body["userName"]
	err := DeleteUserByName(userName)
	if err != nil {
		ctx.JSON(403, gin.H{
			"err": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"result": "delete success",
		})
	}
}

// 删除一个用户的指定角色
func (uc *UserController) DeleteRole(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	var body map[string]string
	_ = json.Unmarshal(data, &body)
	userName := body["userName"]
	roleName := body["roleName"]
	u, err1 := FindUserByName(userName)
	r, err2 := FindRoleByName(roleName)
	if err1 != nil {
		ctx.JSON(403, gin.H{
			"userName": userName,
			"err":      err1.Error(),
		})
	} else if err2 != nil {
		ctx.JSON(403, gin.H{
			"roleName": roleName,
			"err":      err2.Error(),
		})
	} else {
		err := DeleteRoleOfUser(*u, *r)
		if err != nil {
			ctx.JSON(403, gin.H{
				"err": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"result": "delete success",
			})
		}
	}
}
