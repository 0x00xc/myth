/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2021/1/5 16:00
 */
package cry

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"math/rand"
	"time"
)

func pkcsPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcsUnpadding(data []byte) ([]byte, error) {
	length := len(data)
	n := length - int(data[length-1])
	if n > length || n < 0 {
		return data, errors.New("invalid ciphertext")
	}
	return data[:n], nil
}

func AESEncrypt(origin, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	padded := pkcsPadding(origin, block.BlockSize())
	cipherText := make([]byte, len(padded))
	cipher.NewCBCEncrypter(block, iv).CryptBlocks(cipherText, padded)
	return cipherText, nil
}

func AESDecrypt(cipherText, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(cipherText)%block.BlockSize() != 0 {
		return nil, errors.New("input not full blocks")
	}
	out := make([]byte, len(cipherText))
	cipher.NewCBCDecrypter(block, iv).CryptBlocks(out, cipherText)
	return pkcsUnpadding(out)
}

func AESEncryptRandIV(origin, key []byte) ([]byte, error) {
	iv := RandIV()
	b, err := AESEncrypt(origin, key, iv)
	if err != nil {
		return b, err
	}
	return append(iv, b...), err
}

func AESDecryptRandIV(text, key []byte) ([]byte, error) {
	if len(text) <= aes.BlockSize {
		return nil, errors.New("invalid cipher text")
	}
	iv := text[:aes.BlockSize]
	cipherText := text[aes.BlockSize:]
	return AESDecrypt(cipherText, key, iv)
}

func init() {
	rand.Seed(time.Now().Unix())
}

func RandIV() []byte {
	var b = make([]byte, aes.BlockSize)
	for i := 0; i < aes.BlockSize; i++ {
		b[i] = byte(rand.Intn(256))
	}
	return b
}
