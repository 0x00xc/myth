/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/30 17:26
 */
package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"myth/internal/dao"
)

func NewArticle(c *gin.Context) {
	ctx := With(c)
	req := new(NewArticleRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		logrus.WithContext(ctx).Errorln(err)
		ctx.Code(CodeParamsError, err.Error())
		return
	}
	art := new(dao.Article)
	art.UID = ctx.uid
	art.Title = req.Title
	art.Thumb = req.Thumb
	art.Content = req.Content
	art.Format = req.Format
	art.Status = req.Status
	art.Visible = req.Visible
	err := art.New()
	if err != nil {
		logrus.WithContext(ctx).Errorln(err)
		ctx.Code(CodeFailed, err.Error())
		return
	}
	ctx.Data(NewArticleResponse{ID: art.ID})
}

func ArticleList(status int) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := With(c)
		req := new(ArticleListRequest)
		if err := ctx.ShouldBindJSON(req); err != nil {
			ctx.Code(CodeParamsError, err.Error())
			logrus.WithContext(ctx).Errorln(err)
			return
		}

		relation := dao.RelationStranger
		if req.UID == ctx.uid {
			relation = dao.RelationOwn
		}

		if req.Limit == 0 {
			req.Limit = 10
		}
		var data []*dao.Article
		var err error
		switch status {
		case dao.ArticleStatusPublish:
			data, err = dao.ArticleList(req.UID, relation, req.LastID, req.Limit)
		case dao.ArticleStatusDraft:
			data, err = dao.ArticleDraftList(req.UID, req.LastID, req.Limit)
		case dao.ArticleStatusRecycle:
			data, err = dao.ArticleRecycledList(req.UID, req.LastID, req.Limit)
		default:
			ctx.Code(CodeFailed, "invalid article status")
			return
		}
		if err != nil {
			logrus.WithContext(ctx).Errorln(err)
			ctx.Code(CodeFailed, err.Error())
			return
		}
		re := new(ArticleListResponse)
		re.Set(data)
		ctx.Data(re)
	}
}

func EditArticle(c *gin.Context) {
	ctx := With(c)
	req := new(EditArticleRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.Code(CodeParamsError, err.Error())
		logrus.WithContext(ctx).Errorln(err)
		return
	}
	edit := req.M()
	if len(edit) == 0 {
		ctx.Code(CodeParamsError, "empty params")
		return
	}
	err := dao.EditArticle(req.ID, ctx.uid, edit)
	if err != nil {
		ctx.Code(CodeFailed, "")
		logrus.WithContext(ctx).Errorln(err)
		return
	}
	ctx.Data(EditArticleResponse{})
}
