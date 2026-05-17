package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zkep/my-geektime/internal/global"
	"github.com/zkep/my-geektime/internal/service"
	"github.com/zkep/my-geektime/internal/types/geek"
)

func AccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie := global.CONF.Site.Cookie.Geektime
		if cookie == "" {
			c.Abort()
			global.JSON(c, http.StatusOK, nil, "product.no_cookie", "")
			return
		}
		var auth geek.AuthResponse
		if err := service.Authority(cookie, func(r *http.Response) error {
			_, err := service.GetGeekUser(r, &auth)
			return err
		}); err != nil {
			global.JSON(c, http.StatusOK, nil, "product.no_cookie", "")
			return
		}
		c.Set(global.AccessToken, cookie)
		c.Next()
	}
}
