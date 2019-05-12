package secret

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func AesEncrypt(orig, key, iv string) (ret string, err error) {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)

	// 分组秘钥
	block, err := aes.NewCipher(k)
	if err != nil {
		return ret, err
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)
	ret = base64.URLEncoding.EncodeToString(cryted)
	return ret, nil
}

func AesDecrypt(cryted, key, iv string) (ret string, err error) {
	// 转成字节数组
	crytedByte, err := base64.URLEncoding.DecodeString(cryted)
	if err != nil {
		return ret, err
	}
	k := []byte(key)
	// 分组秘钥
	block, err := aes.NewCipher(k)
	if err != nil {
		return ret, err
	}
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	ret = string(orig)
	return ret, nil
}
