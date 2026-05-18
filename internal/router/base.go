package router

import (
	"github.com/gin-gonic/gin"
	v2 "github.com/zkep/my-geektime/internal/api/v2"
)

func base(public, _ *gin.RouterGroup) {
	api := v2.NewBase()
	{
		public.GET("/base/config", api.Config)
		public.POST("/base/login", api.Login)
		public.POST("/base/register", api.Register)
		public.POST("/base/email/code", api.SendEmailCode)
	}
}
