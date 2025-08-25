package main

import "fmt"

// 抽象工厂模式

/* 抽象工厂模式提供一个创建一系列相关或相互依赖对象的接口，而无需指定它们的具体类。它适用于需要创建一组相关产品的场景。
特点：
    每个工厂类可以创建多个相关产品。
    强调产品族的概念，例如 GUI 库中的不同风格组件（Windows 风格、Mac 风格）。
适用场景：
    需要创建一组相关对象的场景。

优点：
    可以创建一组相关对象，保证对象之间的兼容性。
    符合开闭原则，扩展性强。
缺点：
    类的数量会增加，系统复杂度提高。
    新增产品族或产品等级结构时，需要修改抽象工厂接口及其所有实现类。

*/

// DBConnecter 产品A，接口
type DBConnecter interface {
	Connect() string
}

// DBCommander 产品B，接口
type DBCommander interface {
	Command(string) string
}

// 具体产品：MySQL 数据库连接器
type MySQLConnect struct {
}

func (m *MySQLConnect) Connect() string {
	return "MySQL 连接成功"
}

// 具体产品：MySQL 数据库命令执行器
type MySQLCommander struct {
}

func (m *MySQLCommander) Command(sql string) string {
	return "MySQL 执行命令：" + sql
}

// 具体产品：SQLite 数据库连接器
type SQLiteConnect struct {
}

func (s *SQLiteConnect) Connect() string {
	return "SQLite 连接成功"
}

// 具体产品：SQLite 数据库命令执行器
type SQLiteCommander struct {
}

func (s *SQLiteCommander) Command(sql string) string {
	return "SQLite 执行命令：" + sql
}

// 抽象工厂接口
type DBFactory interface {
	CreateConnecter() DBConnecter
	CreateCommander() DBCommander
}

// 具体工厂：MySQL 数据库工厂
type MySQLFactory struct {
}

func (m *MySQLFactory) CreateConnecter() DBConnecter {
	return &MySQLConnect{}
}

func (m *MySQLFactory) CreateCommander() DBCommander {
	return &MySQLCommander{}
}

// 具体工厂：SQLite 数据库工厂
type SQLiteFactory struct {
}

func (s *SQLiteFactory) CreateConnecter() DBConnecter {
	return &SQLiteConnect{}
}

func (s *SQLiteFactory) CreateCommander() DBCommander {
	return &SQLiteCommander{}
}

func UseDB(factory DBFactory) {
	connecter := factory.CreateConnecter()
	commander := factory.CreateCommander()
	fmt.Printf("connecter.Connect(): %v\n", connecter.Connect())
	fmt.Printf("commander.Command(\"SELECT * FROM table\"): %v\n", commander.Command("SELECT * FROM table"))
}

func main() {
	// 使用 MySQL 数据库工厂
	mySQLFactory := &MySQLFactory{}
	UseDB(mySQLFactory)
	// 使用 SQLite 数据库工厂
	SQLiteFactory := &SQLiteFactory{}
	UseDB(SQLiteFactory)
}
