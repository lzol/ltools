package util

import (
	"github.com/axgle/mahonia"
	"golang.org/x/text/encoding/simplifiedchinese"
	"errors"
	"strings"
	"strconv"
)

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

/*
 * 填充字符串至指定位数
 * inputStr string 待填充字符串
 * iType int 填充类型，0为左填充 1为右填充
 * fillStr string 填充的字符串，左/右填充的字符串内容
 * destLength int 指定长度
*/
func StringFill(inputStr string,iType int,fillStr string,destLength int)(outStr string,err error)  {

	if GetRealLength(inputStr)==0{
		inputStr = " "
	}
	outStr = inputStr
	if iType!=0 && iType!=1{
		return outStr,errors.New("填充类型为0或1")
	}
	oriLen := GetRealLength(inputStr)
	fillStrLen :=GetRealLength(fillStr)
	if fillStrLen == 0{
		return outStr,errors.New("待填充字符长度为0")
	}

	if fillStrLen>1{
		m := (destLength-oriLen)%fillStrLen
		if m!=0{
			return outStr,errors.New("填充需截断字符")
		}
	}

	for i:=0;i<(destLength-oriLen)/fillStrLen;i++{
		if iType == 0{
			outStr = fillStr + outStr
		}else{
			outStr = outStr + fillStr
		}
	}
	return outStr,nil
}

//获取字符串真实长度，Go中string为UTF-8存储，len（）获取的长度不准确(按照汉字占2位计算)
func GetRealLength(inputStr string)(length int){
	sl := 0
	rs := []rune(inputStr)
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			sl++
		} else {
			sl += 2
		}
	}
	return sl
}

func SubString(s string, l int) string {
	if len(s) <= l {
		return s
	}
	ss, sl, rl, rs := "", 0, 0, []rune(s)
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			rl = 1
		} else {
			rl = 2
		}
		if sl + rl > l {
			break
		}
		sl += rl
		ss += string(r)
	}
	return ss
}

//将数字字符串转换为银联数据金额类型（12位，无小数点）
func ConvertToUDAmount(input string)(sAmount string,err error){
	//sAmount = strings.Replace(input, ".", "", -1)
	sAmount = strings.Trim(input," ")
	fAmount, err := strconv.ParseFloat(input,32)
	if err != nil {
		return input,errors.New("输入字符非数值型："+input)
	}
	fAmount = fAmount * 100
	sAmount = strconv.FormatFloat(fAmount,'f',-1,32)
	//如果*100，再格式化成字符串后，还有小数点，说明输入值小数点后不止2位
	if strings.Contains(sAmount,"."){
		return input,errors.New("输入字符非数值型："+input)
	}
	sAmount,err = StringFill(sAmount,0,"0",12)
	return sAmount,nil
}

