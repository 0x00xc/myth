/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/30 13:52
 */
package email

import (
	"errors"
	"fmt"
	"myth/pkg/rand"
	"strings"
)

func getCaptcha(email string, id string) (string, error) {
	return storage.GetString(fmt.Sprintf("email_captcha_%s_%s", email, id))
}

func setCaptcha(email string, id string, captcha string) error {
	return storage.Put(fmt.Sprintf("email_captcha_%s_%s", email, id), []byte(captcha), 1800)
}

func delCaptcha(email string, id string) {
	storage.Delete(fmt.Sprintf("email_captcha_%s_%s", email, id))
}

func getCaptchaCD(email string) error {
	_, err := storage.GetUint8(fmt.Sprintf("email_captcha_cd_%s", email))
	return err
}

func setCaptchaCD(email string) {
	storage.PutUint8(fmt.Sprintf("email_captcha_cd_%s", email), 1, 60)
}

func NewCaptcha(email string) (string, error) {
	if getCaptchaCD(email) == nil {
		return "", errors.New("cool down")
	}
	id := rand.String(16)
	captcha := rand.String(6, rand.SEED_NUM)
	err := SendText(email, fmt.Sprintf("captcha code: %s", captcha))
	if err != nil {
		return "", err
	}
	return id, setCaptcha(email, id, captcha)
}

func VerifyCaptcha(email, id string, captchaCode string) error {
	code, err := getCaptcha(email, id)
	if err != nil {
		return err
	}
	if code != strings.ToUpper(captchaCode) {
		return errors.New("invalid captcha")
	}
	delCaptcha(email, id)
	return nil
}
