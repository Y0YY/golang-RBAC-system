package controller

import (
	. "Assignment/model"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
}

// 得到所有角色
func (rc *RoleController) GetAllRoles(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, GetAllRoles())
}

// 得到一个角色的所有权限
func (rc *RoleController) GetAllPermsOfRole(ctx *gin.Context) {
	roleName := ctx.Param("roleName")
	r, err1 := FindRoleByName(roleName)
	if err1 != nil {
		ctx.JSON(403, gin.H{
			"err": err1.Error(),
		})
	} else {
		ps, err2 := GetAllPermsOfRole(*r)
		if err2 != nil {
			ctx.JSON(403, gin.H{
				"err": err2.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, ps)
		}
	}
}

// 创建一个角色
func (rc *RoleController) Create(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	var NewRole Role
	_ = json.Unmarshal(data, &NewRole)
	err := CreatRole(NewRole)
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

// 为一个角色增加一个权限
func (rc *RoleController) AddPerm(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	var body map[string]string
	_ = json.Unmarshal(data, &body)
	roleName := body["roleName"]
	permName := body["permName"]
	r, err1 := FindRoleByName(roleName)
	p, err2 := FindPermByName(permName)
	if err1 != nil {
		ctx.JSON(403, gin.H{
			"roleName": roleName,
			"err":      err1.Error(),
		})
	} else if err2 != nil {
		ctx.JSON(403, gin.H{
			"permName": permName,
			"err":      err2.Error(),
		})
	} else {
		err := AddPerm(*r, *p)
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

// 通过角色ID删除角色
func (rc *RoleController) DeleteById(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	var body map[string]string
	_ = json.Unmarshal(data, &body)
	id := body["roleId"]
	err := DeleteRoleById(id)
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

// 通过角色名称删除角色
func (rc *RoleController) DeleteByName(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	var body map[string]string
	_ = json.Unmarshal(data, &body)
	roleName := body["roleName"]
	err := DeleteRoleByName(roleName)
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

// 删除指定角色的权限
func (rc *RoleController) DeletePerm(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	var body map[string]string
	_ = json.Unmarshal(data, &body)
	roleName := body["roleName"]
	permName := body["permName"]
	r, err1 := FindRoleByName(roleName)
	p, err2 := FindPermByName(permName)
	if err1 != nil {
		ctx.JSON(403, gin.H{
			"userName": roleName,
			"err":      err1.Error(),
		})
	} else if err2 != nil {
		ctx.JSON(403, gin.H{
			"roleName": permName,
			"err":      err2.Error(),
		})
	} else {
		err := DeletePermOfRole(*r, *p)
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
