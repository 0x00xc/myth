/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/29 17:57
 */
package hash

import "crypto/md5"

func MD5(s string) string {
	return getHash(md5.New(), s)
}
