/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/30 13:52
 */
package email

import (
	"github.com/go-gomail/gomail"
	"github.com/sirupsen/logrus"
	"myth/conf"
	"myth/pkg/cache"
)

var storage *cache.MStorage
var dialer = gomail.NewDialer(conf.C.SMTPHost, conf.C.SMTPPort, conf.C.SMTPUser, conf.C.SMTPPass)

func Init(st cache.Storage) {
	storage = &cache.MStorage{Storage: st}
	c, err := dialer.Dial()
	if err != nil {
		logrus.Errorln("connect smtp server failed")
	} else {
		c.Close()
	}
}

func Send(m *gomail.Message) error {
	return dialer.DialAndSend(m)
}

func SendText(to string, content string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", dialer.Username)
	m.SetHeader("To", to)
	m.SetBody("text/plain", content)
	return Send(m)
}
