package util

import "github.com/axgle/mahonia"

const (
	UTF8    = "utf-8"
	GBK     = "gbk"
	GB18030 = "gb18030"
)

/*
  字符串编码转换
*/
func Encode(input string, oriCode string, tarCode string) (output string) {
	oriEnc := mahonia.NewEncoder(oriCode)
	oriStr := oriEnc.ConvertString(input)
	tarEnc := mahonia.NewEncoder(tarCode)
	output = tarEnc.ConvertString(oriStr)
	return output
}

