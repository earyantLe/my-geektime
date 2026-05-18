package service

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/smtp"
	"time"

	"github.com/zkep/my-geektime/internal/global"
	"go.uber.org/zap"
)

// EmailService 邮件服务
type EmailService struct{}

// NewEmailService 创建邮件服务实例
func NewEmailService() *EmailService {
	return &EmailService{}
}

// GenerateVerificationCode 生成6位验证码
func (s *EmailService) GenerateVerificationCode() string {
	const chars = "0123456789"
	code := ""
	for i := 0; i < 6; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		code += string(chars[n.Int64()])
	}
	return code
}

// SendVerificationEmail 发送验证邮件
func (s *EmailService) SendVerificationEmail(to, code string) error {
	// 从配置中读取SMTP配置
	smtpConfig := global.CONF.Site.Email.SMTP

	// 检查配置是否完整
	if smtpConfig.Host == "" || smtpConfig.Port == "" || smtpConfig.Username == "" || smtpConfig.Password == "" || smtpConfig.From == "" {
		global.LOG.Error("SMTP配置不完整，请检查配置文件")
		return fmt.Errorf("SMTP配置不完整")
	}

	host := smtpConfig.Host
	port := smtpConfig.Port
	username := smtpConfig.Username
	password := smtpConfig.Password
	from := smtpConfig.From

	subject := global.CONF.Site.Register.Email.Subject
	if subject == "" {
		subject = "【我的极客时间】邮箱验证码"
	}

	body := global.CONF.Site.Register.Email.Body
	if body == "" {
		body = fmt.Sprintf("您的验证码是：%s，有效期为10分钟。请勿将验证码告知他人。", code)
	}

	msg := []byte(fmt.Sprintf("To: %s\r\nFrom: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/plain; charset=UTF-8\r\n\r\n%s",
		to, from, subject, body))

	addr := fmt.Sprintf("%s:%s", host, port)
	auth := smtp.PlainAuth("", username, password, host)

	err := smtp.SendMail(addr, auth, from, []string{to}, msg)
	if err != nil {
		global.LOG.Error("发送邮件失败", zap.Error(err), zap.String("to", to))
		return fmt.Errorf("发送邮件失败: %w", err)
	}

	global.LOG.Info("验证码邮件发送成功", zap.String("to", to))
	return nil
}

// VerificationCodeStore 验证码存储（简单内存实现，生产环境应使用Redis）
type VerificationCodeStore struct {
	codes map[string]*verificationCode
}

type verificationCode struct {
	Code      string
	ExpiresAt time.Time
}

var verificationStore = &VerificationCodeStore{
	codes: make(map[string]*verificationCode),
}

// StoreVerificationCode 存储验证码
func StoreVerificationCode(email, code string) {
	verificationStore.codes[email] = &verificationCode{
		Code:      code,
		ExpiresAt: time.Now().Add(10 * time.Minute), // 10分钟有效期
	}
}

// VerifyCode 验证验证码
func VerifyCode(email, code string) bool {
	vc, exists := verificationStore.codes[email]
	if !exists {
		return false
	}

	// 检查是否过期
	if time.Now().After(vc.ExpiresAt) {
		delete(verificationStore.codes, email)
		return false
	}

	// 验证成功后删除验证码（一次性使用）
	if vc.Code == code {
		delete(verificationStore.codes, email)
		return true
	}

	return false
}
