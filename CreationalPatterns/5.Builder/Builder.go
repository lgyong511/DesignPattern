package main

import "fmt"

// 建造者模式

/* 它用于分步构建复杂对象。建造者模式的核心思想是将一个复杂对象的构建过程与其表示分离，使得同样的构建过程可以创建不同的表示。
建造者模式特别适用于以下场景：
    对象的构建过程非常复杂，包含多个步骤。
    对象的构建过程需要支持不同的配置或表示。
*/

// computer 产品：电脑
type Computer struct {
	CPU      string
	GPU      string
	Memory   string
	HardDisk string
}

// builder 建造者
type Builder interface {
	BuildCPU()
	BuildGPU()
	BuildMemory()
	BuildHardDisk()
	GetComputer() *Computer
}

// 游戏电脑建造者
type GameComputerBuilder struct {
	computer *Computer
}

func NewGameComputerBuilder() *GameComputerBuilder {
	return &GameComputerBuilder{
		computer: &Computer{},
	}
}

func (g *GameComputerBuilder) BuildCPU() {
	g.computer.CPU = "游戏CPU"
}
func (g *GameComputerBuilder) BuildGPU() {
	g.computer.GPU = "游戏GPU"
}
func (g *GameComputerBuilder) BuildMemory() {
	g.computer.Memory = "游戏内存"
}
func (g *GameComputerBuilder) BuildHardDisk() {
	g.computer.HardDisk = "游戏硬盘"
}
func (g *GameComputerBuilder) GetComputer() *Computer {
	return g.computer
}

// 工作电脑建造者
type WorkComputerBuilder struct {
	computer *Computer
}

func NewWorkComputerBuilder() *WorkComputerBuilder {
	return &WorkComputerBuilder{
		computer: &Computer{},
	}
}

func (w *WorkComputerBuilder) BuildCPU() {
	w.computer.CPU = "工作CPU"
}
func (w *WorkComputerBuilder) BuildGPU() {
	w.computer.GPU = "工作GPU"
}
func (w *WorkComputerBuilder) BuildMemory() {
	w.computer.Memory = "工作内存"
}
func (w *WorkComputerBuilder) BuildHardDisk() {
	w.computer.HardDisk = "工作硬盘"
}
func (w *WorkComputerBuilder) GetComputer() *Computer {
	return w.computer
}

// director 指挥者
type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) Construct() {
	d.builder.BuildCPU()
	d.builder.BuildGPU()
	d.builder.BuildMemory()
	d.builder.BuildHardDisk()
}

func main() {
	var pc Builder
	pc = NewGameComputerBuilder()
	pc = NewWorkComputerBuilder()
	director := NewDirector(pc)
	director.Construct()
	gameComputer := pc.GetComputer()
	fmt.Println(gameComputer)
}
