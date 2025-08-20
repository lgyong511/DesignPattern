package main

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

//单例模式

/* 确保一个类只有一个实例，并提供一个全局访问点。
适用场景：
    配置管理、日志记录、数据库连接池等需要全局唯一实例的场景。
*/

var (
	db   *sql.DB
	once sync.Once
)

// 全局访问点
func GetDB(_type, dsn string) *sql.DB {
	// 保证只初始化一次
	once.Do(func() {
		var err error
		db, err = sql.Open(_type, dsn)
		if err != nil {
			panic(err)
		}
	})
	return db
}

// lgyong
var (
	_db   *sql.DB
	_once sync.Once
)

// 全局访问点
func InitDB() *sql.DB {
	// 保证只初始化一次
	_once.Do(func() {
		var err error
		// Open只检查字符串是否有效
		_db, err = sql.Open("mysql", "root:root@tcp(192.168.31.200:3306)/go_test")
		if err != nil {
			panic(err)
		}
		_db.SetMaxIdleConns(10)  // 设置最大空闲连接数
		_db.SetMaxOpenConns(100) // 设置最大连接数
		//检查数据库连接
		if err = _db.Ping(); err != nil {
			panic(err)
		}
	})
	return _db
}

func main() {
	db1 := GetDB("mysql", "root:root@tcp(192.168.31.200:3306)/go_test")
	defer db1.Close()

	db2 := GetDB("mysql", "root:root@tcp(192.168.31.200:3306)/go_test")
	defer db2.Close()

	//两个实例指向同一个数据库连接
	fmt.Printf("db1: %v\n", db1)
	fmt.Printf("db2: %v\n", db2)

	err := db1.Ping()
	if err != nil {
		panic(err)
	}

	var id int
	var name string
	var age int
	db1.QueryRow("select * from user where id=?", 1).Scan(&id, &name, &age)
	fmt.Println(id, name, age)

	InitDB().QueryRow("select * from user where id=?", 2).Scan(&id, &name, &age)
	fmt.Println(id, name, age)
}
