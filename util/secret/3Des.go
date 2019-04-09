package secret

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
)

/**
 * 加密
 */
func TripleDesEncrypt(orig, key string) string {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)

	// 3DES的秘钥长度必须为24位
	block, _ := des.NewTripleDESCipher(k)
	// 补全码
	origData = PKCS5Padding(origData, block.BlockSize())
	// 设置加密方式
	blockMode := cipher.NewCBCEncrypter(block, k[:8])
	// 创建密文数组
	crypted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(crypted, origData)

	return base64.StdEncoding.EncodeToString(crypted)
}

/**
 * 解密
 */
func TipleDesDecrypt(crypted string, key string) string {
	// 用base64转成字节数组
	cryptedByte, _ := base64.StdEncoding.DecodeString(crypted)
	// key转成字节数组
	k := []byte(key)

	block, _ := des.NewTripleDESCipher(k)
	blockMode := cipher.NewCBCDecrypter(block, k[:8])
	origData := make([]byte, len(cryptedByte))
	blockMode.CryptBlocks(origData, cryptedByte)
	origData = PKCS5UnPadding(origData)

	return string(origData)
}
