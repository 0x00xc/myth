/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/29 17:42
 */
package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"myth/internal/auth"
	"myth/pkg/array"
	"myth/pkg/rand"
	"net/http"
	"strings"
	"time"
)

func RequestID(ctx *gin.Context) {
	ctx.Set("request_id", rand.HexBytes(12))
}

func Logger(c *gin.Context) {
	start := time.Now()
	path := c.Request.URL.Path
	c.Next()
	logrus.WithContext(c).Infof("%d | %s | %s | %v", c.Writer.Status(), path, c.ClientIP(), time.Since(start))
}

func Auth(skip ...string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		if array.InStrings(ctx.FullPath(), skip) {
			return
		}
		token := strings.TrimSpace(strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer "))
		if token == "" {
			token, _ = ctx.Cookie("token")
		}
		if token != "" {
			if info, err := auth.Verify(token); err == nil {
				ctx.Set("uid", info.UID)
				ctx.Set("open_id", info.OpenID)
				return
			}
		}
		ctx.AbortWithStatus(http.StatusForbidden)
	}
}
