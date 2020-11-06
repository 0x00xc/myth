/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/29 16:59
 */
package dao

import (
	"errors"
	"fmt"
	"myth/pkg/hash"
	"myth/pkg/rand"
	"strconv"
	"strings"
)

type Account struct {
	Model
	Username string `gorm:"not null;unique"`
	Password string
	Secret   string
	Email    string

	Name     string
	Avatar   string
	BIO      string
	Gender   int8
	Followed int
	Follower int
}

func (a *Account) Update() error {
	return db.Update(a).Error
}

func FindAccountByID(id int) (*Account, error) {
	acc := new(Account)
	acc.ID = id
	err := db.First(acc, "id = ?", id).Error
	return acc, err
}

func FindAccountByUsername(username string) (*Account, error) {
	acc := new(Account)
	err := db.First(acc, "username = ?", username).Error
	return acc, err
}
func FindAccountByEmail(email string) (*Account, error) {
	acc := new(Account)
	err := db.First(acc, "email = ?", email).Error
	return acc, err
}

func FindAccount(query string) (*Account, error) {
	if id, err := strconv.Atoi(query); err == nil {
		return FindAccountByID(id)
	}
	if n := strings.Index(query, "@"); n > 0 && n < len(query)-3 {
		return FindAccountByEmail(query)
	}
	return FindAccountByUsername(query)
}

func CreateAccount(username, pwd, email string) (*Account, error) {
	secret := rand.String(128)
	password := hash.Sha256(fmt.Sprintf("%s%s%s", secret, pwd, secret))
	acc := new(Account)
	acc.Username = username
	acc.Password = password
	acc.Secret = secret
	acc.Email = email
	acc.Name = "用户_" + rand.String(8)
	err := db.Save(acc).Error
	return acc, err
}

func UpdateAccount(id int, m M) error {
	acc := new(Account)
	acc.ID = id
	return db.Model(acc).Updates(m).Error
}

func (a *Account) Verify(pwd string) error {
	if pwd == "" || a.Secret == "" || a.Password == "" {
		return errors.New("verify failed")
	}
	temp := hash.Sha256(fmt.Sprintf("%s%s%s", a.Secret, pwd, a.Secret))
	if a.Password != temp {
		return errors.New("verify failed")
	}
	a.Secret = rand.String(128)
	a.Password = hash.Sha256(fmt.Sprintf("%s%s%s", a.Secret, pwd, a.Secret))
	UpdateAccount(a.ID, M{"secret": a.Secret, "password": a.Password})
	return nil
}
