/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/12/4 16:06
 */
package dao

type RPGRoom struct {
	Model
	Title     string
	Describe  string
	Password  string
	Status    int
	CreatedBy int
	KP        int
	//RoleList  []RPGRole
}

func (r *RPGRoom) Create() error {
	return db.Save(r).Error
}

func EditRPGRoom(id int, uid int, edit M) error {
	edit.fix("title", "describe", "password", "status", "kp")
	err := db.Model(&RPGRoom{}).Where("id = ? AND kp = ?", id, uid).Updates(edit).Error
	return err
}
