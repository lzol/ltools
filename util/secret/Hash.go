package secret

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func Md5(data string) string {
	md5Ctx := md5.New()                            //md5 init
	md5Ctx.Write([]byte(data))                     //md5 updata
	cipherStr := md5Ctx.Sum(nil)                   //md5 final
	encryptedData := hex.EncodeToString(cipherStr) //hex_digest
	return encryptedData
}

func Md5WithSalt(data,salt string) string {
	md5Ctx := md5.New()                            //md5 init
	md5Ctx.Write([]byte(data))                     //md5 updata
	md5Ctx.Write([]byte(salt))
	cipherStr := md5Ctx.Sum(nil)                   //md5 final
	encryptedData := hex.EncodeToString(cipherStr) //hex_digest
	return encryptedData
}

func Sha256(data string) string{
	sha := sha256.New()
	sha.Write([]byte(data))
	sum := sha.Sum(nil)
	return string(sum)
}

func Sha256WithSalt(data,salt string) string{
	sha := sha256.New()
	sha.Write([]byte(data))
	sha.Write([]byte(salt))
	sum := sha.Sum(nil)
	return string(sum)
}



