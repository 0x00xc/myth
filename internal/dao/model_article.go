/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/30 17:27
 */
package dao

const (
	ArticleStatusRecycle = -1 //
	ArticleStatusDraft   = 0  //
	ArticleStatusPublish = 1  //

	FormatText = 0 //
	FormatMD   = 1 //
	FormatHTML = 2 //
	FormatJSON = 3 //

	RelationStranger = 0
	RelationFollower = 1
	RelationFollowed = 2
	RelationOwn      = 100
)

type Article struct {
	Model
	UID     int
	Title   string
	Thumb   string
	Content string
	Format  int
	Status  int
	Visible int
	Type    int
}

func (a *Article) New() error {
	return db.Save(a).Error
}

func FindArticle(id int, relation int) (*Article, error) {
	var a = new(Article)
	err := db.First(a, "id = ? AND visible <= ?", id, relation).Error
	return a, err
}

func articleList(uid, status, relation, lastID, limit int) ([]*Article, error) {
	if limit == 0 {
		limit = 10
	}
	var data []*Article
	var where = db.Model(&Article{})
	if lastID > 0 {
		where = db.Where("uid = ? AND status = ? AND visible <= ? AND id < ?", uid, status, relation, lastID)
	} else {
		where = db.Where("uid = ? AND status = ? AND visible <= ?", uid, status, relation)
	}
	err := where.Select("id,created_at,updated_at,uid,title,thumb,content,format,status,visible").
		Order("created_at DESC").Limit(limit).Find(&data).Error
	return data, err
}

func ArticleList(uid int, relation int, lastID int, limit int) ([]*Article, error) {
	return articleList(uid, ArticleStatusPublish, relation, lastID, limit)
}

func ArticleDraftList(uid int, lastID int, limit int) ([]*Article, error) {
	return articleList(uid, ArticleStatusDraft, RelationOwn, lastID, limit)
}

func ArticleRecycledList(uid int, lastID int, limit int) ([]*Article, error) {
	return articleList(uid, ArticleStatusRecycle, RelationOwn, lastID, limit)
}

func EditArticle(id int, uid int, edit M) error {
	edit.fix("title", "thumb", "content", "format", "status", "visible")
	err := db.Model(&Article{}).Where("id = ? AND uid = ?", id, uid).Updates(edit).Error
	return err
}

func DeleteArticle(id int) error {
	a := &Article{}
	a.ID = id
	return db.Delete(a).Error
}
