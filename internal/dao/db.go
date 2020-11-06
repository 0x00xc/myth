/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/29 16:39
 */
package dao

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/jinzhu/gorm"

	"myth/conf"
)

var db *gorm.DB

func Init() error {
	name := conf.C.DatabaseName
	source := conf.C.DatabaseSource
	temp, err := gorm.Open(name, source)
	if err != nil {
		return err
	}
	db = temp
	db.SingularTable(true)
	//createTable()
	//db.LogMode(true)
	return nil
}

func createTable() {
	db.CreateTable(&Account{})
	db.CreateTable(&Article{})
	db.CreateTable(&Follower{})
	//db.CreateTable(&Moment{})
	//db.CreateTable(&Timeline{})
}
