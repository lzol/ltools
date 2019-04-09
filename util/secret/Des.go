package secret

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
)

/**
 * DES加密方法
 */
func MyDesEncrypt(orig, key string) string{

	// 将加密内容和秘钥转成字节数组
	origData := []byte(orig)
	k := []byte(key)

	// 秘钥分组
	block, _ := des.NewCipher(k)

	//将明文按秘钥的长度做补全操作
	origData = PKCS5Padding(origData, block.BlockSize())

	//设置加密方式－CBC
	blockMode := cipher.NewCBCDecrypter(block, k)

	//创建明文长度的字节数组
	crypted := make([]byte, len(origData))

	//加密明文
	blockMode.CryptBlocks(crypted, origData)

	//将字节数组转换成字符串，base64编码
	return base64.StdEncoding.EncodeToString(crypted)

}

/**
 * DES解密方法
 */
func MyDESDecrypt(data string, key string) string {

	k := []byte(key)

	//将加密字符串用base64转换成字节数组
	crypted, _ := base64.StdEncoding.DecodeString(data)

	//将字节秘钥转换成block快
	block, _ := des.NewCipher(k)

	//设置解密方式－CBC
	blockMode := cipher.NewCBCEncrypter(block, k)

	//创建密文大小的数组变量
	origData := make([]byte, len(crypted))

	//解密密文到数组origData中
	blockMode.CryptBlocks(origData, crypted)

	//去掉加密时补全的部分
	origData = PKCS5UnPadding(origData)

	return string(origData)
}
