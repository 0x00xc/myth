/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/29 17:01
 */
package dao

import (
	"myth/pkg/array"
	"time"
)

type Model struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type M map[string]interface{}

func (m M) fix(col ...string) {
	for k := range m {
		if !array.InStrings(k, col) {
			delete(m, k)
		}
	}
}
