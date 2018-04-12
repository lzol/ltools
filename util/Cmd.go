package util

import (
	"bufio"
	"os/exec"
	"bytes"
	"fmt"
	"io"
)

func ExecCommand(commandName string, params []string, needResult bool) (result string, err error) {
	cmd := exec.Command(commandName, params...)
	stdout, err := cmd.StdoutPipe()
	defer stdout.Close()
	if err != nil {
		return result, err
	}
	err = cmd.Start()
	if err!=nil{
		fmt.Println(err)
		return result,err
	}
	reader := bufio.NewReader(stdout)
	var buffer bytes.Buffer

	//实时循环读取输出流中的一行内容
	if needResult {
		for {
			line, _,err2 := reader.ReadLine()
			fmt.Println("line",string(line),err2)
			buffer.WriteString(string(line))
			if err2 != nil && err2!=io.EOF{
				return result, err2
				break
			}
			//if io.EOF == err2 {
			//	break
			//}
		}
	}
	cmd.Wait()
	result = buffer.String()
	return result, nil
}
