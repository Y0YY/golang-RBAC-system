package main

import (
	"Assignment/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	//创建服务
	ginSever := gin.Default()

	userController := controller.UserController{}
	roleController := controller.RoleController{}
	permController := controller.PermController{}
	//路由组
	userGroup := ginSever.Group("/user")
	{
		userGroup.GET("/getAllUsers", userController.GetAllUsers)
		userGroup.GET("/getAllRolesOfUser/:userName", userController.GetAllRolesOfUser)
		userGroup.GET("/getAllPermsOfUser/:userName", userController.GetAllPermsOfUser)
		userGroup.GET("/isPrmitted/:userName/:permName", userController.IsPrmitted)
		userGroup.POST("/create", userController.Create)
		userGroup.POST("/addRole", userController.AddRole)
		userGroup.DELETE("/deleteById", userController.DeleteById)
		userGroup.DELETE("/deleteByName", userController.DeleteByName)
		userGroup.DELETE("/deleteRole/:userName/:roleName", userController.DeleteRole)
	}
	roleGroup := ginSever.Group("/role")
	{
		roleGroup.GET("/getAllRoles", roleController.GetAllRoles)
		roleGroup.GET("/getAllPermsOfRole/:roleName", roleController.GetAllPermsOfRole)
		roleGroup.POST("/creat", roleController.Create)
		roleGroup.POST("/addPerm", roleController.AddPerm)
		roleGroup.DELETE("/deleteById", roleController.DeleteById)
		roleGroup.DELETE("/deleteByName", roleController.DeleteByName)
		roleGroup.DELETE("/deletePerm/:roleName/:permName", roleController.DeletePerm)
	}
	permGroup := ginSever.Group("/perm")
	{
		permGroup.GET("/getAllperms", permController.ShowAllPerms)
		permGroup.POST("/create", permController.Create)
		permGroup.DELETE("/deleteById", permController.DeleteById)
		permGroup.DELETE("/deleteByName", permController.DeleteByName)
	}

	//服务器端口
	ginSever.Run("127.0.0.1:8080")

}
