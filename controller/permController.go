package controller

import (
	. "Assignment/model"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PermController struct {
}

// 得到所有权限
func (pc *PermController) ShowAllPerms(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ShowAllPerms())
}

// 创建一个权限
func (pc *PermController) Create(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	var NewPerm Perm
	_ = json.Unmarshal(data, &NewPerm)
	if CreatPerm(NewPerm) == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": "success",
		})
	}
}

// 通过权限ID删除一个权限
func (pc *PermController) DeleteById(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	var body map[string]string
	_ = json.Unmarshal(data, &body)
	id := body["permId"]
	err := DeletePermById(id)
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

// 通过权限名称删除一个权限
func (pc *PermController) DeleteByName(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	var body map[string]string
	_ = json.Unmarshal(data, &body)
	permName := body["permName"]
	err := DeletePermByName(permName)
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
