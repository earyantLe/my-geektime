package v2

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zkep/my-geektime/internal/global"
	"github.com/zkep/my-geektime/internal/model"
	"github.com/zkep/my-geektime/internal/types/user"
	"github.com/zkep/my-geektime/libs/utils"
)

type User struct{}

func NewUser() *User {
	return &User{}
}

func (u *User) List(c *gin.Context) {
	var req user.UserListRequest
	if err := c.ShouldBind(&req); err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
	if req.PerPage <= 0 || (req.PerPage > 200) {
		req.PerPage = 10
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	roleId := c.GetFloat64(global.Role)
	if roleId != user.AdminRoleId {
		global.FAIL(c, "fail.msg", "no auth")
		return
	}
	ret := user.UserListResponse{
		Rows: make([]user.User, 0, 10),
	}
	var ls []*model.User
	tx := global.DB.Model(&model.User{})
	if req.Status > 0 {
		tx = tx.Where("status = ?", req.Status)
	}
	tx = tx.Where("role_id = ?", user.MemeberRoleId)
	tx = tx.Where("deleted_at = 0")
	tx = tx.Order("id DESC")
	if err := tx.Count(&ret.Count).
		Offset((req.Page - 1) * req.PerPage).
		Limit(req.PerPage).
		Find(&ls).Error; err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
	for _, l := range ls {
		ret.Rows = append(ret.Rows, user.User{
			Uid:         l.Uid,
			UserName:    l.UserName,
			NickName:    l.NickName,
			Avatar:      l.Avatar,
			Status:      l.Status,
			AccessToken: l.AccessToken,
			RoleId:      l.RoleId,
			CreatedAt:   l.CreatedAt,
			UpdatedAt:   l.UpdatedAt,
		})
	}
	global.OK(c, ret)
}

func (u *User) Status(c *gin.Context) {
	var req user.UserStatusRequest
	if err := c.ShouldBind(&req); err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
	if err := global.DB.Model(&model.User{}).
		Where(&model.User{Uid: req.Uid}).
		Updates(&model.User{Status: req.Status}).Error; err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
	global.OK(c, nil)
}

func (u *User) Create(c *gin.Context) {
	var req user.UserCreateRequest
	if err := c.ShouldBind(&req); err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}

	// 校验登录名长度
	if len(req.UserName) < 6 {
		global.FAIL(c, "fail.msg", "登录名至少需要6个字符")
		return
	}

	// 校验密码长度
	if len(req.Password) < 6 {
		global.FAIL(c, "fail.msg", "密码至少需要6个字符")
		return
	}

	// 检查用户名是否已存在
	var existingUser model.User
	err := global.DB.Where("user_name = ?", req.UserName).First(&existingUser).Error
	if err == nil {
		// 用户已存在，更新状态、密码、昵称并清除删除时间
		hashedPassword := utils.BcryptHash(req.Password)
		now := time.Now().Unix()
		if err := global.DB.Model(&model.User{}).
			Where("user_name = ?", req.UserName).
			Updates(map[string]interface{}{
				"nick_name":  req.NickName,
				"password":   hashedPassword,
				"status":     user.UserStatusActive,
				"deleted_at": 0,
				"updated_at": now,
			}).Error; err != nil {
			global.FAIL(c, "fail.msg", err.Error())
			return
		}
		global.OK(c, nil)
		return
	}

	// 生成 UID
	uid := utils.HalfUUID()

	// 密码加密
	hashedPassword := utils.BcryptHash(req.Password)

	now := time.Now().Unix()
	newUser := &model.User{
		Uid:       uid,
		UserName:  req.UserName,
		NickName:  req.NickName,
		Password:  hashedPassword,
		RoleId:    user.MemeberRoleId,
		Status:    user.UserStatusActive,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := global.DB.Create(newUser).Error; err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}

	global.OK(c, nil)
}

func (u *User) Delete(c *gin.Context) {
	var req user.UserDeleteRequest
	if err := c.ShouldBind(&req); err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}

	// 检查用户是否存在
	var userModel model.User
	if err := global.DB.Where("uid = ?", req.Uid).First(&userModel).Error; err != nil {
		global.FAIL(c, "fail.msg", "用户不存在")
		return
	}

	// 软删除
	if err := global.DB.Model(&model.User{}).
		Where("uid = ?", req.Uid).
		Update("deleted_at", time.Now().Unix()).Error; err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}

	global.OK(c, nil)
}
