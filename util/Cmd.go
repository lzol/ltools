package util

import (
	"bufio"
	"os/exec"
	"io"
	"bytes"
	"fmt"
)

func ExecCommand(commandName string, params []string, needResult bool) (result string, err error) {
	cmd := exec.Command(commandName, params...)
	fmt.Println(commandName)
	stdout, err := cmd.StdoutPipe()
	defer stdout.Close()
	if err != nil {
		return result, err
	}
	err = cmd.Start()
	fmt.Println(cmd.ProcessState)
	if err!=nil{
		fmt.Println(err)
		return result,err
	}
	reader := bufio.NewReader(stdout)
	var buffer bytes.Buffer

	//实时循环读取输出流中的一行内容
	if needResult {
		for {
			line, err2 := reader.ReadString('\n')
			buffer.WriteString(line)
			if err2 != nil && io.EOF != err2 {
				return result, err2
				break
			}
			if io.EOF == err2 {
				break
			}
		}
	}
	cmd.Wait()
	result = buffer.String()
	return result, nil
}
