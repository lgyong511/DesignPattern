package main

import (
	"errors"
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"strings"
)

//简单工厂模式

/* 简单工厂并不是一个正式的设计模式，而是一种编程习惯。它通过一个工厂类来封装对象的创建逻辑，客户端只需要传递参数给工厂类，由工厂类决定创建哪种对象。
特点：
    只有一个工厂类，负责创建所有产品。
    通过条件判断（如 switch 或 if-else）来决定创建哪种产品。
适用场景：
    产品种类较少，且创建逻辑简单的场景。
优点：
    简单易用，适合小型项目。
缺点：
    不符合开闭原则（OCP），新增产品时需要修改工厂类。
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

// Commander
type Commander interface {
	GetIP(string) ([]net.IP, error)
}

// Windows 实现 Commander 接口
type Windows struct {
}

// GetIP 获取 Windows 系统的 IP 地址
func (w *Windows) GetIP(command string) (ips []net.IP, err error) {
	if command == "" {
		return nil, ErrCommandEmpty
	}
	cmd := exec.Command("powershell", "/C", command)
	return ExecuteCommand(cmd)
}

// Linux 实现 Commander 接口
type Linux struct {
}

// GetIP 获取 Linux 系统的 IP 地址
func (l *Linux) GetIP(command string) (ips []net.IP, err error) {
	if command == "" {
		return nil, ErrCommandEmpty
	}
	cmd := exec.Command("bash", "-c", command)
	return ExecuteCommand(cmd)
}

// Darwin 实现 Commander 接口
type Darwin struct {
}

// GetIP 获取 Darwin 系统的 IP 地址
func (d *Darwin) GetIP(command string) (ips []net.IP, err error) {
	if command == "" {
		return nil, ErrCommandEmpty
	}
	cmd := exec.Command("zsh", "-c", command)
	return ExecuteCommand(cmd)
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

func NewCommander() Commander {
	switch runtime.GOOS {
	case WindowsOS:
		return &Windows{}
	case LinuxOS:
		return &Linux{}
	case DarwinOS:
		return &Darwin{}
	default:
		return nil
	}
}

func main() {
	Commander := NewCommander()
	if Commander == nil {
		fmt.Println("not support os")
		return
	}
	ips, err := Commander.GetIP("ifconfig | awk '/240:?/' |awk '{print $2}'")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
