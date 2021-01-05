/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2021/1/5 16:34
 */
package cry

import "testing"

func TestAESEncrypt(t *testing.T) {
	text := []byte("hello")
	key := RandIV()
	res, err := AESEncryptRandIV(text, key)
	t.Log(res, err)
	b, err := AESDecryptRandIV(res, key)
	t.Log(string(b), err)
}

func TestAESDecrypt(t *testing.T) {
	text := []byte{12, 11, 214, 51, 164, 202, 242, 221, 83, 55, 242, 151, 24, 51, 40, 249, 246, 168, 60, 166, 23, 240, 224, 29, 41, 51, 113, 66, 233, 190, 234, 134}
	key := []byte{4, 49, 25, 203, 139, 141, 138, 97, 245, 2, 85, 24, 110, 132, 111, 174}
	b, err := AESDecryptRandIV(text, key)
	t.Log(string(b))
	t.Log(err)
}
