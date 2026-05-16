package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zkep/my-geektime/internal/global"
)

func AccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie := global.CONF.Site.Cookie.Geektime
		if cookie == "" {
			c.Abort()
			global.JSON(c, http.StatusOK, nil, "product.no_cookie", "")
			return
		}
		c.Set(global.AccessToken, cookie)
		c.Next()
	}
}
