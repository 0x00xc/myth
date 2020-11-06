/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/30 16:09
 */
package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myth/conf"
	"myth/internal/controller"
	"myth/internal/dao"
	"net/http"
)

func (s *HTTPServer) route() {
	s.srv = new(http.Server)
	s.srv.Addr = fmt.Sprintf(":%d", conf.C.AppPort)

	router := gin.New()
	router.Use(controller.RequestID, controller.Logger)

	v1 := router.Group("/api/v1/")

	v1.POST("/captcha/email", controller.EmailCaptcha)
	v1.POST("/account/signin", controller.SignIn)
	v1.POST("/account/signup", controller.SignUp)

	fo := v1.Group("", controller.Auth)

	fo.POST("/account/follow", controller.Follow)
	fo.POST("/account/unfollow", controller.Unfollow)
	fo.POST("/account/followed/edit", controller.EditFollowed)
	fo.POST("/account/followed/list", controller.FollowedList)
	fo.POST("/account/follower/list", controller.FollowerList)

	art := v1.Group("/article", controller.Auth)
	art.POST("/new", controller.NewArticle)
	art.POST("/edit", controller.EditArticle)
	art.POST("/list", controller.ArticleList(dao.ArticleStatusPublish))
	art.POST("/draft", controller.ArticleList(dao.ArticleStatusDraft))
	art.POST("/recycled", controller.ArticleList(dao.ArticleStatusRecycle))

	s.srv.Handler = router
}
