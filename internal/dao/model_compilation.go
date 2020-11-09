/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/11/6 16:18
 */
package dao

import "github.com/jinzhu/gorm"

type Compilation struct {
	Model
	Title    string
	Describe string
	Cover    string
	Type     int
}

type CompilationArticle struct {
	Model
	CompilationID int
	ArticleID     int
	Index         int
}

func NewCompilation(title string, desc string, cover string) (*Compilation, error) {
	c := &Compilation{Title: title, Describe: desc, Cover: cover}
	err := db.Save(c).Error
	return c, err
}

func CompilationIndex(compilationID int) (int, error) {
	c := new(CompilationArticle)
	c.CompilationID = compilationID
	err := db.Where(c, "compilation_id = ?", compilationID).Order("index DESC").First(c).Error
	if err == nil {
		return c.Index, nil
	}
	if err == gorm.ErrRecordNotFound {
		return 0, nil
	}
	return 0, err
}

func AddCompilation(compilationID int, articleIDs []int, indexOffset int) error {
	var data []*CompilationArticle
	for i, v := range articleIDs {
		data = append(data, &CompilationArticle{
			CompilationID: compilationID,
			ArticleID:     v,
			Index:         i + indexOffset + 1,
		})
	}
	return db.Save(data).Error
}

func RemoveCompilationArticle(compilationID, articleID int) error {
	return db.Delete(&CompilationArticle{}, "compilation_id = ? AND article_id = ?", compilationID, articleID).Error
}

func EditCompilationArticleIndex(compilationID, articleID, index int) error {
	err := db.Model(&CompilationArticle{}).
		Where("compilation_id = ? AND article_id = ?", compilationID, articleID).
		Update("index", index).Error
	return err
}

func DeleteCompilation(id int) error {
	c := &Compilation{}
	c.ID = id
	return db.Delete(c).Error
	//tx := db.Begin()
	//if err := tx.Delete(c).Error; err != nil {
	//	return err
	//}
	//if err := tx.Delete(&CompilationArticle{}, "compilation_id = ?", id).Error; err != nil {
	//	tx.Rollback()
	//	return err
	//}
	//return tx.Commit().Error
}
