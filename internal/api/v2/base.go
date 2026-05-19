package v2

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zkep/my-geektime/internal/global"
	"github.com/zkep/my-geektime/internal/model"
	"github.com/zkep/my-geektime/internal/service"
	"github.com/zkep/my-geektime/internal/types/base"
	"github.com/zkep/my-geektime/internal/types/geek"
	"github.com/zkep/my-geektime/internal/types/user"
	"github.com/zkep/my-geektime/libs/utils"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
)

type Base struct{}

func NewBase() *Base {
	return &Base{}
}

func (b *Base) Login(c *gin.Context) {
	var r base.LoginRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
	var info model.User
	switch r.Type {
	case base.LoginWithName:
		var req base.LoginWithNameRequest
		if err := json.Unmarshal(r.Data, &req); err != nil {
			global.FAIL(c, "fail.msg", err.Error())
			return
		}
		if err := binding.Validator.ValidateStruct(req); err != nil {
			global.FAIL(c, "fail.msg", err.Error())
			return
		}
		if err := global.DB.
			Where(&model.User{
				UserName: req.Account,
				Status:   user.UserStatusActive,
			}).
			First(&info).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				global.FAIL(c, "base.login.error")
				return
			}
			global.FAIL(c, "fail.msg", err.Error())
			return
		}
		if !utils.BcryptCheck(req.Password, info.Password) {
			global.FAIL(c, "base.login.error")
			return
		}
	default:
		global.FAIL(c, "base.login.type")
		return
	}
	token, expire, err := global.JWT.TokenGenerator(func(claims jwt.MapClaims) {
		claims[global.Identity] = info.Uid
		claims[global.Role] = info.RoleId
	})
	if err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "OK",
		"token":  token,
		"user": base.User{
			Uid:      info.Uid,
			UserName: info.UserName,
			NickName: info.NickName,
			Avatar:   info.Avatar,
			GeekAuth: len(info.AccessToken) > 0,
			RoleId:   info.RoleId,
		},
		"expire": expire.Format(time.RFC3339),
	})
	c.SetCookie(global.Analogjwt, token, int(expire.Unix()), "/", "", false, false)
}

func (b *Base) Register(c *gin.Context) {
	var r base.RegisterRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
	switch r.Type {
	case base.LoginWithName:
		var req base.RegisterWithNameRequest
		if err := json.Unmarshal(r.Data, &req); err != nil {
			global.FAIL(c, "fail.msg", err.Error())
			return
		}
		if err := binding.Validator.ValidateStruct(req); err != nil {
			global.FAIL(c, "fail.msg", err.Error())
			return
		}
		var info model.User
		if err := global.DB.
			Where(&model.User{
				UserName: req.Account,
				Status:   user.UserStatusActive,
			}).
			First(&info).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			global.FAIL(c, "fail.msg", err.Error())
			return
		}
		if len(info.UserName) > 0 {
			global.FAIL(c, "base.register.exists")
			return
		}
		var count int64
		if err := global.DB.Model(&model.User{}).Count(&count).Error; err != nil {
			global.FAIL(c, "fail.msg", err.Error())
			return
		}
		// first user is admin
		if count == 0 {
			info.RoleId = user.AdminRoleId
		}
		info.Uid = utils.HalfUUID()
		info.UserName = req.Account
		info.NickName = req.Account
		info.Password = utils.BcryptHash(req.Password)
		if err := global.DB.Create(&info).Error; err != nil {
			global.FAIL(c, "fail.msg", err.Error())
			return
		}
	case base.LoginWithEmail:
		var req base.RegisterWithEmailRequest
		if err := json.Unmarshal(r.Data, &req); err != nil {
			global.FAIL(c, "fail.msg", err.Error())
			return
		}
		if err := binding.Validator.ValidateStruct(req); err != nil {
			global.FAIL(c, "fail.msg", err.Error())
			return
		}
		// 验证邮箱验证码
		if !service.VerifyCode(req.Email, req.Code) {
			global.FAIL(c, "base.register.code_invalid")
			return
		}
		// 检查邮箱是否已注册
		var info model.User
		if err := global.DB.
			Where(&model.User{
				Email:  req.Email,
				Status: user.UserStatusActive,
			}).
			First(&info).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			global.FAIL(c, "fail.msg", err.Error())
			return
		}
		if len(info.Email) > 0 {
			global.FAIL(c, "base.register.exists")
			return
		}
		var count int64
		if err := global.DB.Model(&model.User{}).Count(&count).Error; err != nil {
			global.FAIL(c, "fail.msg", err.Error())
			return
		}
		// first user is admin
		if count == 0 {
			info.RoleId = user.AdminRoleId
		}
		info.Uid = utils.HalfUUID()
		info.Email = req.Email
		info.UserName = req.Email // 使用邮箱作为用户名
		info.NickName = req.Email
		info.Password = utils.BcryptHash(req.Password)
		if err := global.DB.Create(&info).Error; err != nil {
			global.FAIL(c, "fail.msg", err.Error())
			return
		}
	default:
		global.FAIL(c, "base.register.type")
		return
	}
	global.OK(c, nil)
}

func (b *Base) Config(c *gin.Context) {
	ret := base.Config{
		RegisterTypes: global.CONF.Site.Register.Types,
		LoginTypes:    global.CONF.Site.Login.Types,
		LoginGuests:   []base.Guest{},
	}

	// 根据配置的登录类型，添加对应的访客账号
	for _, loginType := range global.CONF.Site.Login.Types {
		guest := base.Guest{Type: loginType}
		switch loginType {
		case "name":
			guest.Account = global.CONF.Site.Login.Guest.Name.Account
			guest.Password = global.CONF.Site.Login.Guest.Name.Password
		case "email":
			guest.Account = global.CONF.Site.Login.Guest.Email.Account
			guest.Password = global.CONF.Site.Login.Guest.Email.Password
		}
		if guest.Account != "" {
			ret.LoginGuests = append(ret.LoginGuests, guest)
		}
	}

	global.OK(c, ret)
}

func (b *Base) SendEmailCode(c *gin.Context) {
	var req base.SendEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
	if err := binding.Validator.ValidateStruct(req); err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}

	// 生成验证码
	emailService := service.NewEmailService()
	code := emailService.GenerateVerificationCode()

	// 发送验证邮件
	if err := emailService.SendVerificationEmail(req.Email, code); err != nil {
		global.FAIL(c, "base.email.send_failed")
		return
	}

	// 存储验证码
	service.StoreVerificationCode(req.Email, code)

	global.OK(c, nil)
}

func (b *Base) RefreshCookie(c *gin.Context) {
	var req base.RefreshCookieRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
	if err := binding.Validator.ValidateStruct(req); err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
	// 验证cookie有效性
	var auth geek.AuthResponse
	if err := service.Authority(req.Cookie, func(r *http.Response) error {
		_, err := service.GetGeekUser(r, &auth)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		global.FAIL(c, "product.no_valid_cookie")
		return
	}

	global.CONF.Site.Cookie.Geektime = req.Cookie

	// 保存配置到文件
	raw, err := yaml.Marshal(global.CONF)
	if err != nil {
		global.FAIL(c, "fail.msg", err)
		return
	}
	customConfPath := global.CustomConfigFile
	if len(global.CONFPath) > 0 {
		customConfPath = filepath.Join(filepath.Dir(global.CONFPath), global.CustomConfigFile)
	}
	if err = os.WriteFile(customConfPath, raw, os.ModePerm); err != nil {
		global.FAIL(c, "fail.msg", err)
		return
	}
	global.OK(c, nil)
}
