package main

import (
	"fmt"
	"ltools/util/secret"
)

func main(){
	orig := "Hello World!"
	fmt.Println("原文：", orig)

	//声明秘钥,利用此秘钥实现明文的加密和密文的解密，长度必须为8
	key := "12345678"

	//加密
	encyptCode := secret.MyDesEncrypt(orig, key)
	fmt.Println("密文：", encyptCode)

	//解密
	decyptCode := secret.MyDESDecrypt(encyptCode, key)
	fmt.Println("解密结果：", decyptCode)

	orig = "hello world"
	// 3DES的秘钥长度必须为24位
	key = "123456781234567812345678"
	fmt.Println("原文：", orig)

	encryptCode := secret.TripleDesEncrypt(orig, key)
	fmt.Println("密文：", encryptCode)

	decryptCode := secret.TripleDesDecrypt(encryptCode, key)
	fmt.Println("解密结果：", decryptCode)

	orig = "1"
	key = "B31F2A75FBF94099"
	iv := "1234567890123456"
	encryptCode,_ = secret.AesEncrypt(orig, key,iv)
	fmt.Println("AesEncrypt密文：" , encryptCode)

	key = "B31F2A75FBF94099"
	iv = "1234567890123456"
	decryptCode,_ = secret.AesDecrypt("b7ox4yi5XW5PSg1uzge1F/hrIh/MWJGvPw7nur9r2w4=", key,iv)
	fmt.Println("解密结果：", decryptCode)

	// 第一种调用方法
	sum := secret.Sha256("hello world")
	fmt.Printf("%x\n", sum)

	// 第二种调用方法
	h := secret.Sha256WithSalt("hello world","aabbdssdfdsfdsssssssss")
	fmt.Printf("%x\n", h)
}
