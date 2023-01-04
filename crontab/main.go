package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var (
		cmd    *exec.Cmd
		output []byte
		err    error
	)
	//生成cmd
	cmd = exec.Command("C:\\cygwin64\\bin\\bash.exe", "-c", "echo hello")
	//执行命令
	if output, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		return
	}
	//打印字进程
	fmt.Printf(string(output))
}
