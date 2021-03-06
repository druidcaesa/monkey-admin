package router

import "github.com/gin-gonic/gin"
import v1 "monkey-admin/api/v1"

func initNoticeRouter(router *gin.RouterGroup) {
	v := new(v1.NoticeApi)
	group := router.Group("/notice")
	{
		group.GET("/list",v.List)
		//添加公告
		group.POST("/add",v.Add)
		//删除
		group.DELETE("/:ids",v.Delete)
		//查询
		group.GET("/:id",v.Get)
		//修改
		group.PUT("/edit",v.Edit)
	}
}
