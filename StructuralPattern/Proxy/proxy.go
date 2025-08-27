package main

import "fmt"

// 代理模式
//为某个对象提供一个代理对象，由代理对象来控制对原对象的访问。
// 适用场景：
//     权限控制、日志记录、缓存、延迟加载等。

/* 🔹 结构
代理模式通常包含三个角色：
	1.	Subject（抽象主题）
	•	定义真实对象和代理对象的公共接口。
	2.	RealSubject（真实主题）
	•	真正执行逻辑的对象。
	3.	Proxy（代理对象）
	•	持有对真实对象的引用，在调用真实对象方法前后加上额外逻辑（权限控制、日志、延迟加载、缓存等）。
*/

// 抽象接口
type Command interface {
	Execute(cmd string) (string, error)
}

// 真实对象
type RealCommand struct {
}

func (r *RealCommand) Execute(cmd string) (string, error) {
	return "执行命令" + cmd, nil
}

// 代理对象
type ProxyCommand struct {
	realCommand *RealCommand
}

func (p *ProxyCommand) Execute(cmd string) (string, error) {
	// 可以添加一些额外的逻辑
	fmt.Println("代理对象执行命令")
	result, err := p.realCommand.Execute(cmd)
	fmt.Println("代理对象执行命令完成")
	return result, err
}

func main() {
	realCommand := &RealCommand{}
	proxyCommand := &ProxyCommand{realCommand: realCommand}
	result, err := proxyCommand.Execute("ls")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
