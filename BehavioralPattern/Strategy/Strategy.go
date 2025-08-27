package main

import (
	"errors"
	"fmt"
	"os/exec"
)

// 策略模式

var (
	ErrParamsEmpty = errors.New("params is empty")
)

// Commander 策略接口
type Commander interface {
	// Execute 执行命令
	//参数：要执行的命令
	//返回值：命令执行结果或者错误信息
	Execute(string) ([]byte, error)
}

// Windows 具体策略windows执行命令
type Windows struct {
}

// Execute 执行命令
func (w *Windows) Execute(command string) ([]byte, error) {
	if command == "" {
		return nil, ErrParamsEmpty
	}
	return exec.Command("powershell", "/C", command).Output()
}

// linux 具体策略linux执行命令
type Linux struct {
}

// Execute 执行命令
func (l *Linux) Execute(command string) ([]byte, error) {
	if command == "" {
		return nil, ErrParamsEmpty
	}
	return exec.Command("sh", "-c", command).Output()
}

// darwin 具体策略darwin（MAC）执行命令
type Darwin struct {
}

// Execute 执行命令
func (d *Darwin) Execute(command string) ([]byte, error) {
	if command == "" {
		return nil, ErrParamsEmpty
	}
	return exec.Command("zsh", "-c", command).Output()
}

// 上下文
type Context struct {
	command Commander
}

// NewContext 创建上下文
func NewContext(commander Commander) *Context {
	return &Context{
		command: commander,
	}
}

// Execute 执行命令
func (c *Context) Execute(command string) ([]byte, error) {
	return c.command.Execute(command)
}

func main() {
	// 创建上下文
	ctx := NewContext(&Darwin{})
	// 执行命令
	output, err := ctx.Execute("ifconfig")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
}
