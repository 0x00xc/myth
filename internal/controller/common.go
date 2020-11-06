/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/29 17:30
 */
package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	CodeOK            = 1
	CodeParamsError   = 20000
	CodeUsernameError = 20010
	CodePasswordError = 20011
	CodeCaptchaError  = 20020
	CodeNotFound      = 20030
	//CodeFailed        = 30000
	CodeServiceError  = 50000
)

type Response struct {
	RequestID string      `json:"request_id"`
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
}

type Context struct {
	*gin.Context
	uid       int
	requestID string
}

func With(c *gin.Context) *Context {
	ctx := new(Context)
	ctx.Context = c
	ctx.uid = ctx.GetInt("uid")
	ctx.requestID = ctx.GetString("request_id")
	return ctx
}

func (c *Context) out(re *Response) {
	var code = http.StatusOK
	//switch re.Code {
	//case CodeOK:
	//	code = http.StatusOK
	//case CodeParamsError:
	//	code = http.StatusBadRequest
	//case CodeServiceError:
	//	code = http.StatusInternalServerError
	//}
	re.RequestID = c.requestID
	c.JSON(code, re)
}

func (c *Context) Code(code int, message ...string) {
	re := new(Response)
	re.Code = code
	if len(message) > 0 {
		re.Message = message[0]
	}
	c.out(re)
}

func (c *Context) Data(data interface{}) {
	c.out(&Response{Data: data})
}
