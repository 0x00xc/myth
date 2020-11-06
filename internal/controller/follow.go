/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/30 16:33
 */
package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"myth/internal/dao"
)

type FollowRequest struct {
	FollowedID int    `json:"followed_id" binding:"required"`
	Mark       string `json:"mark"`
	Group      string `json:"group"`
}

type FollowResponse struct{}

func Follow(c *gin.Context) {
	ctx := With(c)
	req := new(FollowRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		logrus.WithContext(ctx).Errorln(ctx.uid, req.FollowedID, err)
		ctx.Code(CodeParamsError, err.Error())
		return
	}
	err := dao.Follow(ctx.uid, req.FollowedID, req.Mark, req.Group)
	if err != nil {
		logrus.WithContext(ctx).Errorln("follow failed", ctx.uid, req.FollowedID, err)
		ctx.Code(CodeFailed, err.Error())
		return
	}
	ctx.Data(FollowResponse{})
}

type UnfollowRequest struct {
	FollowedID int `json:"followed_id" binding:"required"`
}
type UnfollowResponse struct{}

func Unfollow(c *gin.Context) {
	ctx := With(c)
	req := new(UnfollowRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		logrus.WithContext(ctx).Errorln(ctx.uid, req.FollowedID, err)
		ctx.Code(CodeParamsError, err.Error())
		return
	}
	err := dao.Unfollow(ctx.uid, req.FollowedID)
	if err != nil {
		logrus.WithContext(ctx).Errorln("follow failed", ctx.uid, req.FollowedID, err)
		ctx.Code(CodeFailed, err.Error())
		return
	}
	ctx.Data(UnfollowResponse{})
}

type EditFollowRequest struct {
	FollowedID int    `json:"followed_id" binding:"required"`
	Key        string `json:"key"         binding:"required"`
	Value      string `json:"value"       binding:"required"`
}

type EditFollowResponse struct{}

func EditFollowed(c *gin.Context) {
	ctx := With(c)
	req := new(EditFollowRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		logrus.WithContext(ctx).Errorln(ctx.uid, req.FollowedID, err)
		ctx.Code(CodeParamsError, err.Error())
		return
	}
	var err error
	switch req.Key {
	case "mark":
		err = dao.EditFollowMark(req.FollowedID, req.Value)
	case "group":
		err = dao.EditFollowGroup(req.FollowedID, req.Value)
	default:
		ctx.Code(CodeParamsError, "invalid key, only supported 'mark', 'group'")
		return
	}
	if err != nil {
		logrus.WithContext(ctx).Errorln("follow failed", ctx.uid, req.FollowedID, req.Key, err)
		ctx.Code(CodeFailed, err.Error())
		return
	}
	ctx.Data(EditFollowResponse{})
}

type FollowedListRequest struct {
	LastID int `json:"last_id"`
	Count  int `json:"count"`
}

type FollowedListResponse struct {
	List   []FollowInfo `json:"list"`
	LastID int          `json:"last_id"`
}

type FollowInfo struct {
	UID        int    `json:"uid"`
	FollowedID int    `json:"followed_id"`
	Both       bool   `json:"both"`
	Group      string `json:"group"`
	Mark       string `json:"mark"`
}

func FollowedList(c *gin.Context) {
	ctx := With(c)
	req := new(FollowedListRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		logrus.WithContext(ctx).Errorln(ctx.uid, err)
		ctx.Code(CodeParamsError, err.Error())
		return
	}
	data, err := dao.FollowList(ctx.uid, req.LastID, req.Count)
	if err != nil {
		logrus.WithContext(ctx).Errorln(ctx.uid, err)
		ctx.Code(CodeFailed, err.Error())
		return
	}
	re := new(FollowedListResponse)
	if len(data) > 0 {
		re.LastID = data[len(data)-1].ID
	}
	for _, v := range data {
		re.List = append(re.List, FollowInfo{
			UID:        v.UID,
			FollowedID: v.FollowedID,
			Both:       v.Both,
			Group:      v.Group,
			Mark:       v.Mark,
		})
	}
	ctx.Data(re)
}

func FollowerList(c *gin.Context) {
	ctx := With(c)
	req := new(FollowedListRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		logrus.WithContext(ctx).Errorln(ctx.uid, err)
		ctx.Code(CodeParamsError, err.Error())
		return
	}
	data, err := dao.FollowerList(ctx.uid, req.LastID, req.Count)
	if err != nil {
		logrus.WithContext(ctx).Errorln(ctx.uid, err)
		ctx.Code(CodeFailed, err.Error())
		return
	}
	re := new(FollowedListResponse)
	if len(data) > 0 {
		re.LastID = data[len(data)-1].ID
	}
	for _, v := range data {
		re.List = append(re.List, FollowInfo{
			UID:        v.UID,
			FollowedID: v.FollowedID,
			Both:       v.Both,
			Group:      v.Group,
			Mark:       v.Mark,
		})
	}
	ctx.Data(re)
}
