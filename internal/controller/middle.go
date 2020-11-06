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

func Auth(ctx *gin.Context) {
	token := strings.TrimSpace(strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer "))
	if token == "" {
		token, _ = ctx.Cookie("token")
	}
	if token == "" {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}
	info, err := auth.Verify(token)
	if err != nil {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}
	ctx.Set("uid", info.UID)
	ctx.Set("open_id", info.OpenID)
}
