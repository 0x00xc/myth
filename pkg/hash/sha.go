/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/29 17:56
 */
package hash

import (
	"crypto/sha256"
	"crypto/sha512"
)

func Sha256(s string) string {
	return getHash(sha256.New(), s)
}

func Sha512(s string) string {
	return getHash(sha512.New(), s)
}
