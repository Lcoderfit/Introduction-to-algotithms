package main

/*
go get -u -v github.com/go-sql-driver/mysql

1.sql包提供了保证SQL或类SQL数据库的泛用接口
因此无论使用什么数据库的驱动，调用的接口都是一样的
（就是说写一遍代码，即使改了驱动，代码也不需要改变，只需要改变驱动的配置）

2.为什么需要数据库连接池（事先建立特定数量的连接，按需使用，事后无需再建立，用完放回）
因为这样可以避免每次都去与数据库建立连接，减少资源消耗

3.原生支持连接池，并发安全:database/sql(没有具体的实现，只列出了第三方驱动需要实现的接口)

4.下载驱动：go get -u github.com/go-sql-driver/mysql
*/

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 是一个连接池对象，并发安全
var db *sql.DB

type User struct {
	id   int
	name string
	age  int
}

// 查询
func query() {
	//sqlStr := "select name, age from user where id=?"
	//var u User
}

// 插入
func insert() {

}

func initDB() (err error) {
	// 连接数据库
	dsn := "root:lcoder124541@tcp(127.0.0.1:3306)/test"
	// 不会去dsn中的校验账号跟密码是否正确，只会判断dsn格式是否正确
	// 判断数据库是否连接需要用Ping指令
	// github.com/go-sql-driver/mysql会先在init函数中注册一个mysql的驱动
	// 然后通过Open("mysql", dsn)就可以找到该驱动并把dsn传给它
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("open %s failed, err:%v\n", dsn, err)
		return
	}

	// 查看是否ping通,尝试连接数据库
	err = db.Ping()
	if err != nil {
		fmt.Printf("open %s failed, err:%v\n", dsn, err)
		return
	}
	return
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("数据库连接失败:%v\n", err)
		return
	}
	fmt.Println("数据库连接成功")

	// 1.查询单条记录
	sqlStr := `select id, name, age from user where id=?`
	var u1 User

	// 最多查询一条记录，args表示sqlStr中的？所代表的参数
	rowObj := db.QueryRow(sqlStr, 1)
	// 内部会自动关闭
	rowObj.Scan(&u1.id, &u1.name, &u1.age)
	// 打印结构
	fmt.Printf("%v\n", u1)

}
