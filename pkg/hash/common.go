/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/29 17:57
 */
package hash

import (
	"encoding/hex"
	"hash"
)

func getHash(h hash.Hash, s string) string {
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
