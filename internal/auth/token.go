/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/30 11:49
 */
package auth

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"myth/pkg/hash"
	"myth/pkg/rand"
	"strings"
	"time"
)

type Data struct {
	UID       int    `json:"uid"`
	OpenID    string `json:"open_id"`
	ExpireAt  int64  `json:"expire_at"`
	Nonce     string `json:"nonce"`
	b         []byte
	signature string
}

func NewToken(uid int, exp int64, secret string) string {
	var a = Data{
		UID:      uid,
		ExpireAt: time.Now().Unix() + exp,
		Nonce:    rand.String(16),
	}
	b, _ := json.Marshal(a)
	signature := hash.MD5(string(b) + secret)
	content := base64.URLEncoding.EncodeToString(b)
	return fmt.Sprintf("%s.%s", content, signature)
}

func parse(token string) (*Data, error) {
	fields := strings.Split(token, ".")
	if len(fields) != 2 || fields[0] == "" || fields[1] == "" {
		return nil, errors.New("invalid token")
	}
	b, err := base64.URLEncoding.DecodeString(fields[0])
	if err != nil {
		return nil, errors.New("invalid token:" + err.Error())
	}
	var data = new(Data)
	if err := json.Unmarshal(b, data); err != nil {
		return nil, errors.New("invalid token:" + err.Error())
	}
	data.b = b
	data.signature = fields[1]
	return data, nil
}

func (d *Data) valid(secret string) error {
	if hash.MD5(string(d.b)+secret) != d.signature {
		return errors.New("invalid token signature")
	}
	return nil
}
