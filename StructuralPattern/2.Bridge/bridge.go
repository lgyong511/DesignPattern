package main

import "fmt"

// 桥接模式
// 将抽象部分与实现部分分离，使它们都可以独立地变化。

/* 🔹 结构
桥接模式有 4 个主要角色：
	1.	Abstraction（抽象类）
	•	定义对外的高层接口，包含一个 实现类接口引用。
	2.	RefinedAbstraction（扩展抽象类）
	•	在 Abstraction 的基础上扩展功能。
	3.	Implementor（实现类接口）
	•	定义底层的实现方法（不对外直接暴露）。
	4.	ConcreteImplementor（具体实现类）
	•	真正的实现逻辑。 */

// 实现类接口
type Printer interface {
	Print()
}

// 具体实现类
type Espone struct {
}

func (e *Espone) Print() {
	fmt.Println("Espone 打印")
}

type Lenovo struct {
}

func (l *Lenovo) Print() {
	fmt.Println("Lenovo 打印")
}

// 抽象类
type Computer interface {
	Print()
	SetPrinter(p Printer)
}

// 具体抽象类
type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	m.printer.Print()
}
func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

// 具体抽象类
type Windows struct {
	printer Printer
}

func (w *Windows) Print() {
	w.printer.Print()
}
func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}

func main() {
	// 创建具体实现类
	espone := &Espone{}
	lenovo := &Lenovo{}

	// 创建具体抽象类
	mac := &Mac{}
	windows := &Windows{}

	// 桥接
	mac.SetPrinter(espone)
	mac.Print()

	mac.SetPrinter(lenovo)
	mac.Print()

	windows.SetPrinter(espone)
	windows.Print()

	windows.SetPrinter(lenovo)
	windows.Print()
}
