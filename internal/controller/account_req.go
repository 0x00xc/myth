/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/11/5 17:11
 */
package controller

type SignInRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	CaptchaID string `json:"captcha_id"`
	Captcha   string `json:"captcha"`
}

type SignInResponse struct {
	Token  string `json:"token"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	BIO    string `json:"bio"`
	Gender int8   `json:"gender"`
}

type SignUpRequest struct {
	Username  string `json:"username"   binding:"required"` //
	Password  string `json:"password"   binding:"required"` //
	Email     string `json:"email"      binding:"required"` //
	EmailID   string `json:"email_id"   binding:"required"` //
	EmailCode string `json:"email_code" binding:"required"` //
}

type SignUpResponse struct {
	UID    int    `json:"uid"`
	OpenID string `json:"open_id"`
	Token  string `json:"token"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	BIO    string `json:"bio"`
	Gender int8   `json:"gender"`
}

type EmailCaptchaRequest struct {
	Email     string `json:"email"      binding:"required"` //
	CaptchaID string `json:"captcha_id" binding:"required"` //
	Captcha   string `json:"captcha"    binding:"required"` //
}

type EmailCaptchaResponse struct {
	EmailID string `json:"email_id"`
}
