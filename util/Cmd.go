package util

import (
	"bufio"
	"os/exec"
	"bytes"
	"io"
	"github.com/pkg/errors"
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
		return result,err
	}
	reader := bufio.NewReader(stdout)
	var buffer bytes.Buffer

	//实时循环读取输出流中的一行内容
	if needResult {
		for {
			line,err2 := reader.ReadString('\n')
			buffer.WriteString(string(line))
			if err2 != nil && err2!=io.EOF{
				return result, err2
				break
			}
			if io.EOF == err2 {
				break
			}
		}
	}
	cmd.Wait()
	//执行某些命令时，没有输出，Error也是nil，但是实际上没执行成功，所以判断一下执行状态
	if !cmd.ProcessState.Success(){
		return result,errors.New("执行失败，退出状态："+cmd.ProcessState.String())
	}

	result = buffer.String()
	return result, nil
}
