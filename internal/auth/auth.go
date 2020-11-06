/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/29 17:47
 */
package auth

import (
	"errors"
	"fmt"
	"myth/internal/dao"
	"myth/pkg/cache"
	"myth/pkg/rand"
	"time"
)

var storage *cache.MStorage

func Init(st cache.Storage) {
	storage = &cache.MStorage{Storage: st}
}

func getSecret(uid int) (string, error) {
	return storage.GetString(fmt.Sprintf("auth_%d", uid))
}

func setSecret(uid int, secret string) error {
	return storage.Put(fmt.Sprintf("auth_%d", uid), []byte(secret))
}

func Auth(username, password string) (*dao.Account, string, error) {
	acc, err := dao.FindAccount(username)
	if err != nil {
		return nil, "", err
	}
	if err := acc.Verify(password); err != nil {
		return nil, "", err
	}
	secret, _ := getSecret(acc.ID)
	if secret == "" {
		secret = rand.String(16)
		err := setSecret(acc.ID, secret)
		if err != nil {
			return nil, "", err
		}
	}
	token := NewToken(acc.ID, 86400*7, secret)
	return acc, token, nil
}

func Verify(token string) (*Data, error) {
	data, err := parse(token)
	if err != nil {
		return nil, err
	}
	if data.ExpireAt < time.Now().Unix() {
		return nil, errors.New("invalid token: expired")
	}
	secret, _ := getSecret(data.UID)
	if secret == "" {
		return nil, errors.New("invalid token: secret not found")
	}
	return data, data.valid(secret)
}
