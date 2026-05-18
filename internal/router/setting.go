package router

import (
	"github.com/gin-gonic/gin"
	v2 "github.com/zkep/my-geektime/internal/api/v2"
)

func setting(_, private *gin.RouterGroup) {
	api := v2.NewSetting()
	p := private.Group("/setting")
	{
		p.GET("/query", api.Query)
		p.POST("/update", api.Update)
	}
}
