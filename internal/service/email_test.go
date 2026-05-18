package service

import (
	"testing"
)

func TestGenerateVerificationCode(t *testing.T) {
	service := NewEmailService()

	// 测试生成验证码
	code := service.GenerateVerificationCode()

	// 验证验证码长度为6
	if len(code) != 6 {
		t.Errorf("Expected code length 6, got %d", len(code))
	}

	// 验证验证码只包含数字
	for _, c := range code {
		if c < '0' || c > '9' {
			t.Errorf("Expected digit, got %c", c)
		}
	}

	t.Logf("Generated code: %s", code)
}

func TestVerificationCodeStore(t *testing.T) {
	email := "test@example.com"
	code := "123456"

	// 存储验证码
	StoreVerificationCode(email, code)

	// 验证正确的验证码
	if !VerifyCode(email, code) {
		t.Error("Expected valid code to be verified")
	}

	// 验证错误的验证码
	if VerifyCode(email, "wrong") {
		t.Error("Expected wrong code to fail verification")
	}

	// 验证一次性使用（第二次应该失败）
	if VerifyCode(email, code) {
		t.Error("Expected code to be used only once")
	}
}

func TestVerificationCodeExpiration(t *testing.T) {
	email := "test2@example.com"
	code := "654321"

	// 存储验证码
	StoreVerificationCode(email, code)

	// 立即验证应该成功
	if !VerifyCode(email, code) {
		t.Error("Expected valid code to be verified immediately")
	}
}
