/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/11/5 17:11
 */
package controller

import "myth/internal/dao"

type NewArticleRequest struct {
	Title   string `json:"title"   binding:"required"`
	Thumb   string `json:"thumb"`
	Content string `json:"content" binding:"required"`
	Format  int    `json:"format"`
	Status  int    `json:"status"`
	Visible int    `json:"visible"`
}

type NewArticleResponse struct {
	ID int `json:"id"`
}

type ArticleListRequest struct {
	UID    int `json:"uid"`
	LastID int `json:"last_id"`
	Limit  int `json:"limit"`
}

type ArticleListResponse struct {
	List   []Article `json:"list"`
	LastID int       `json:"last_id"`
}

type Article struct {
	ID        int    `json:"id"`
	UID       int    `json:"uid"`
	Title     string `json:"title"`
	Thumb     string `json:"thumb"`
	Format    int    `json:"format"`
	Status    int    `json:"status"`
	Visible   int    `json:"visible"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Content   string `json:"content,omitempty"`
}

func (a *Article) Set(v *dao.Article) {
	a.ID = v.ID
	a.UID = v.UID
	a.Title = v.Title
	a.Thumb = v.Thumb
	a.Format = v.Format
	a.Status = v.Status
	a.Visible = v.Visible
	a.CreatedAt = v.CreatedAt.Format("2006-01-02 15:04:05")
	a.UpdatedAt = v.UpdatedAt.Format("2006-01-02 15:04:05")
	a.Content = v.Content
}

func (r *ArticleListResponse) Set(data []*dao.Article) {
	r.List = make([]Article, 0)
	for _, v := range data {
		r.List = append(r.List, Article{
			ID:        v.ID,
			UID:       v.UID,
			Title:     v.Title,
			Thumb:     v.Thumb,
			Format:    v.Format,
			Status:    v.Status,
			Visible:   v.Visible,
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: v.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	if len(r.List) > 0 {
		r.LastID = r.List[len(r.List)-1].ID
	}
}

type EditArticleRequest struct {
	ID      int     `json:"id"`
	Title   *string `json:"title"`
	Thumb   *string `json:"thumb"`
	Content *string `json:"content"`
	Format  *int    `json:"format"`
	Status  *int    `json:"status"`
	Visible *int    `json:"visible"`
}

func (r *EditArticleRequest) M() dao.M {
	var m = dao.M{}
	if r.Title != nil {
		m["title"] = *r.Title
	}
	if r.Thumb != nil {
		m["thumb"] = *r.Thumb
	}
	if r.Content != nil {
		m["content"] = *r.Content
	}
	if r.Format != nil {
		m["format"] = *r.Format
	}
	if r.Status != nil {
		m["status"] = *r.Status
	}
	if r.Visible != nil {
		m["visible"] = *r.Visible
	}
	return m
}

type EditArticleResponse struct {
}

type ArticleDetailRequest struct {
	UID       int `json:"uid"        binding:"required"` //
	ArticleID int `json:"article_id" binding:"required"` //
}

type ArticleDetailResponse struct {
	Article Article `json:"article"`
}
