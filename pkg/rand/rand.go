/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/23 16:37
 */
package rand

import (
	crand "crypto/rand"
	"encoding/hex"
	"math/big"
	"math/rand"
	"time"
)

const (
	SEED_NUM          = "0123456789"
	SEED_LETTER_LOWER = "abcdefghijklmnopqrstuvwxyz"
	SEED_LETTER_UPPER = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	SEED_LETTER       = SEED_LETTER_LOWER + SEED_LETTER_UPPER
	SEED_WORD         = SEED_NUM + SEED_LETTER
)

func init() {
	rand.Seed(time.Now().Unix())
}

//String 生成随机字符串
func String(n int, seeds ...string) string {
	seed := SEED_NUM + SEED_LETTER_LOWER
	if len(seeds) > 0 {
		seed = seeds[0]
	}
	buffer := make([]byte, n)
	max := big.NewInt(int64(len(seed)))

	for i := 0; i < n; i++ {
		r, err := crand.Int(crand.Reader, max)
		if err != nil {
			return ""
		}
		buffer[i] = seed[r.Int64()]
	}
	return string(buffer)
}

//Int 生成随机int数
func Intn(n int) int {
	return rand.Intn(n)
}

//Float64 生成随机浮点数
func Float64(n int) float64 {
	if n == 1 {
		return rand.Float64()
	}
	return float64(rand.Intn(n-1)) + rand.Float64()
}

//Bytes 生成随机字节
func Bytes(n int) []byte {
	data := make([]byte, n)
	for i := 0; i < n; i++ {
		data[i] = byte(rand.Intn(256))
	}
	return data
}

//Bytes 生成随机字节并转为16进制
func HexBytes(n int) string {
	return hex.EncodeToString(Bytes(n))
}
