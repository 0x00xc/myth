/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/30 16:09
 */
package app

import (
	"github.com/gin-gonic/gin"
	"myth/internal/chat"
	"myth/internal/controller"
	"myth/internal/dao"
	"net/http"
)

func (s *HTTPServer) route() {
	router := gin.New()
	router.Use(controller.RequestID, controller.Logger)
	router.Use(controller.Auth(
		"/api/v1/captcha/email",
		"/api/v1/account/signin",
		"/api/v1/account/signup",
		"/api/v1/article/list",
		"/api/v1/article/detail",
	))

	v1 := router.Group("/api/v1/")
	v1.POST("/captcha/email", controller.EmailCaptcha)

	acc := v1.Group("/account")

	acc.POST("/signin", controller.SignIn)
	acc.POST("/signup", controller.SignUp)
	acc.POST("/follow", controller.Follow)
	acc.POST("/unfollow", controller.Unfollow)
	acc.POST("/followed/edit", controller.EditFollowed)
	acc.POST("/followed/list", controller.FollowedList)
	acc.POST("/follower/list", controller.FollowerList)

	art := v1.Group("/article")
	art.POST("/new", controller.NewArticle)
	art.POST("/edit", controller.EditArticle)
	art.POST("/list", controller.ArticleList(dao.ArticleStatusPublish))
	art.POST("/draft", controller.ArticleList(dao.ArticleStatusDraft))
	art.POST("/recycled", controller.ArticleList(dao.ArticleStatusRecycle))
	art.POST("/detail", controller.ArticleDetail)

	ichat := v1.Group("/chat")
	ichat.GET("/ws", wrap(chat.Handler()))
	ichat.POST("/ws", wrap(chat.Handler())) //实际上一般是通过GET方法请求，可能用不到POST
	s.srv.Handler = router
}

func wrap(h http.Handler) func(c *gin.Context) {
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
