package common

/* 2022/01/12 leo AES-CBC加解密函式
* 	此為AES-CBC 加解密功能
*	若使用aes-128-cbc演算法 則 common.AES_CBC_Key 須設定16字元長(要剛好 不能多不能少)
*	若使用aes-256-cbc演算法 則 common.AES_CBC_Key 須設定32字元長(要剛好 不能多不能少)
*   common.AES_CBC_Iv 長度固定為16字元長
*	AES_CBC_Key與AES_CBC_Iv 需與其他加解密端相同，否則會失敗
* 	Aes_encode 傳入字串即可獲得加密後結果
* 	Aes_decode 傳入字串即可獲得解密後結果
 */
import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"strings"
)

func Aes_encode(in string) string {
	data_str := []byte(in)
	key := AES_CBC_Key
	iv := []byte(AES_CBC_Iv)
	plaintext := pkcs7Padding(data_str)
	ciphertext := make([]byte, len(plaintext))
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		LogError(fmt.Sprintln(err))
		return ""
	}

	cipher.NewCBCEncrypter(block, iv).CryptBlocks(ciphertext, plaintext)
	result := base64.StdEncoding.EncodeToString(ciphertext)
	return result
}

func Aes_decode(in string) string {
	key := AES_CBC_Key
	iv := []byte(AES_CBC_Iv)
	data_str, err := base64.StdEncoding.DecodeString(strings.TrimSpace(in))

	if err != nil {
		LogError(fmt.Sprintln(in, err))
		return ""
	}
	if len(data_str)%aes.BlockSize != 0 {
		LogError(fmt.Sprintln(in))
		return ""
	}
	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		LogError(fmt.Sprintln(in, err))
		return ""
	}
	cipher.NewCBCDecrypter(block, iv).CryptBlocks(data_str, data_str)
	result := string(pkcs7UnPadding(data_str))
	return result
}

func pkcs7Padding(ciphertext []byte) []byte {
	padding := aes.BlockSize - len(ciphertext)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	if length < unpadding {
		return plantText[:]
	}
	return plantText[:(length - unpadding)]
}
