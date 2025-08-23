package main

import (
	"errors"
	"fmt"
	"net"
	"os/exec"
	"strings"
)

//工厂方法模式

/* 工厂方法模式是定义了一个创建对象的接口，但将具体的创建逻辑延迟到子类中。每个子类负责创建一种具体的产品。
特点：
    每个产品对应一个工厂类。
    符合开闭原则，新增产品时只需增加新的工厂类，无需修改现有代码。
适用场景：
    产品种类较多，且创建逻辑复杂的场景。

优点：
    符合开闭原则，扩展性强。
    每个工厂只负责一种产品的创建，职责单一。
缺点：
    类的数量会增加，系统复杂度提高。
*/

var (
	//命令为空
	ErrCommandEmpty = errors.New("命令为空")
	//解析IP错误
	ErrParseIP = errors.New("解析IP错误")
	//命令不存在或错误
	ErrCommandNotExist = errors.New("命令不存在或错误")
)

const (
	// 微软系统
	WindowsOS = "windows"
	// linux系统
	LinuxOS = "linux"
	// 苹果系统
	DarwinOS = "darwin"
)

type Commander interface {
	Execute(string) ([]net.IP, error)
}

// CommandW windows类
type CommandW struct {
}

func (c *CommandW) Execute(command string) ([]net.IP, error) {
	if command == "" {
		return nil, ErrCommandEmpty
	}
	cmd := exec.Command("powershell", "/C", command)
	return ExecuteCommand(cmd)
}

// CommandL linux类
type CommandL struct {
}

func (c *CommandL) Execute(command string) ([]net.IP, error) {
	if command == "" {
		return nil, ErrCommandEmpty
	}
	cmd := exec.Command("bash", "-c", command)
	return ExecuteCommand(cmd)
}

// CommandD darwin类
type CommandD struct {
}

func (c *CommandD) Execute(command string) ([]net.IP, error) {
	if command == "" {
		return nil, ErrCommandEmpty
	}
	cmd := exec.Command("zsh", "-c", command)
	return ExecuteCommand(cmd)
}

// CommandFactory 工厂接口
type CommandFactory interface {
	CreateCommand() Commander
}

// WindowsFactory windows工厂
type WindowsFactory struct {
}

func (w *WindowsFactory) CreateCommand() Commander {
	return &CommandW{}
}

// LinuxFactory linux工厂
type LinuxFactory struct {
}

func (l *LinuxFactory) CreateCommand() Commander {
	return &CommandL{}
}

// DarwinFactory darwin工厂
type DarwinFactory struct {
}

func (d *DarwinFactory) CreateCommand() Commander {
	return &CommandD{}
}

// 使用commander执行命令
func UseCommander(factory CommandFactory, command string) ([]net.IP, error) {
	commander := factory.CreateCommand()
	return commander.Execute(command)
}

// ExecuteCommand 执行命令并返回结果
func ExecuteCommand(cmd *exec.Cmd) (ips []net.IP, err error) {
	// 返回标准输出
	var output []byte
	output, err = cmd.Output()
	if err != nil || len(output) == 0 {
		return nil, errors.Join(err, ErrCommandNotExist)
	}
	strs := strings.Split(string(output), "\n")
	for _, str := range strs {
		str = strings.TrimSpace(str)
		ip := net.ParseIP(str)
		if ip != nil {
			ips = append(ips, ip)
		}
	}
	if len(ips) == 0 {
		return nil, ErrParseIP
	}
	return
}

func main() {
	// // windows命令
	// windowsCommand := `Get-NetRoute -AddressFamily IPv6 | Where-Object { $_.DestinationPrefix.StartsWith("240") -and $_.DestinationPrefix.endsWith("/64") } | ForEach-Object { ($_.DestinationPrefix -split '::/')[0] + ":suffix of other mac" }`
	// // linux命令
	// linuxCommand := `ip -6 route | awk '{print $1}' | awk '/240:?/' | awk -F:: '{print $1 ":9209:d0ff:fe09:781d"}'`
	// // darwin命令
	// darwinCommand := `ifconfig | awk '/240:?/' |awk '{print $2}'`

	darwinFactory := &DarwinFactory{}
	ips, err := UseCommander(darwinFactory, "ifconfig | awk '/240:?/' |awk '{print $2}'")
	if err != nil {
		fmt.Println(err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}

}
