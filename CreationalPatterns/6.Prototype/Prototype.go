package main

import "fmt"

// 原型模式

/* 它通过复制现有对象来创建新对象，而不是通过新建类的方式。原型模式的核心思想是利用对象的克隆能力，避免重复初始化，特别适用于创建成本较高的对象。

适用场景：
    资源-intensive 对象的创建，如数据库连接、文件操作等。
    避免重复初始化，如配置对象的创建。

使用原型模式，如果有引用类型，则需要考虑深拷贝和浅拷贝的问题

浅拷贝只复制对象本身而不复制其引用的对象，深拷贝则会递归地复制整个对象图。

这需要根据需求选择适当的拷贝方式
*/

// 原型接口
// 定义一个克隆方法，用于创建新的对象
type Prototyper interface {
	Clone() Prototyper
}

type Prototype struct {
	Name string
	Age  int
}

// 实现克隆方法
func (p *Prototype) Clone() Prototyper {
	return &Prototype{
		Name: p.Name,
		Age:  p.Age,
	}
}

// 展示
func (p *Prototype) Show() {
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}

func main() {
	// 创建一个原型对象
	prototype := &Prototype{
		Name: "张三",
		Age:  18,
	}

	// 克隆原型对象
	clone := prototype.Clone().(*Prototype)

	// 展示克隆对象
	clone.Show()

	// 修改克隆对象的属性
	clone.Name = "李四"
	clone.Age = 20

	// 展示修改后的克隆对象
	clone.Show()
	// 展示原型对象
	prototype.Show()
}
