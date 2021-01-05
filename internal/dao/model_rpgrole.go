/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/12/4 16:06
 */
package dao

import (
	"database/sql/driver"
	"encoding/json"
)

type Value struct {
	Key string
	Val int
}

type Values []Value

func (v Values) Value() (driver.Value, error) {
	return json.Marshal(v)
}

func (v *Values) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	b, _ := src.([]byte)
	return json.Unmarshal(b, v)
}

type RPGRole struct {
	Model
	CreatedBy  int
	RoomID     int
	Name       string
	Bio        string
	Health     int
	SAN        int
	Status     int
	Attributes Values
	Skills     Values
}

func (r *RPGRole) Create() error {
	return db.Save(r).Error
}

func EditRPGRole(id int, uid int, edit M) error {
	edit.fix("name", "bio", "health", "san", "status", "attributes", "skills")
	err := db.Model(&RPGRole{}).Where("id = ? AND created_by = ?", id, uid).Updates(edit).Error
	return err
}

func RPGRoleList(roomID int) ([]RPGRole, error) {
	var data []RPGRole
	err := db.Find(&data, "room_id = ?", roomID).Error
	return data, err
}
