/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/29 17:29
 */
package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"myth/conf"
	"myth/internal/auth"
	"myth/internal/dao"
	"myth/internal/email"
)

func SignIn(c *gin.Context) {
	ctx := With(c)
	req := new(SignInRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.Code(CodeParamsError, err.Error())
		return
	}
	acc, token, err := auth.Auth(req.Username, req.Password)
	if err != nil {
		ctx.Code(CodePasswordError, err.Error())
		return
	}
	ctx.SetCookie("token", token, 86400*7, "", "", conf.C.Secure(), true)
	ctx.Data(SignInResponse{
		Token:  token,
		Name:   acc.Name,
		Avatar: acc.Avatar,
		BIO:    acc.BIO,
		Gender: acc.Gender,
	})
}

func SignUp(c *gin.Context) {
	ctx := With(c)
	req := new(SignUpRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.Code(CodeParamsError, err.Error())
		logrus.WithContext(ctx).Errorln(err)
		return
	}
	//TODO
	//if err := email.VerifyCaptcha(req.Email, req.EmailID, req.EmailCode); err != nil {
	//	ctx.Code(CodeCaptchaError, "invalid email captcha")
	//	return
	//}
	acc, err := dao.CreateAccount(req.Username, req.Password, req.Email)
	if err != nil {
		ctx.Code(CodeServiceError, err.Error())
		logrus.WithContext(ctx).Errorln(err)
		return
	}
	token := auth.NewToken(acc.ID, 86400*7, acc.Secret)
	ctx.Data(&SignUpResponse{
		UID:    acc.ID,
		Token:  token,
		Name:   acc.Name,
		Avatar: acc.Avatar,
		BIO:    acc.BIO,
		Gender: acc.Gender,
	})
}

func EmailCaptcha(c *gin.Context) {
	ctx := With(c)
	req := new(EmailCaptchaRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		logrus.WithContext(ctx).Errorln(err)
		ctx.Code(CodeParamsError, err.Error())
		return
	}
	//TODO verify image captcha
	emailID, err := email.NewCaptcha(req.Email)
	if err != nil {
		logrus.WithContext(ctx).Errorln(err)
		ctx.Code(CodeServiceError, err.Error())
		return
	}
	ctx.Data(&EmailCaptchaResponse{EmailID: emailID})
}
