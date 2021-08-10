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
	sqlStr := "select name, age from user where id=?"
	var u1 User
	// 内部会自动关闭
	// 必须传入指针，否则不会修改其值
	rowObj := db.QueryRow(sqlStr, 1)
	rowObj.Scan(&u1.id, &u1.name, &u1.age)
	// 打印结构
	fmt.Printf("%v\n", u1)
}

// 查询多条语句
func queryMore(n int) {
	// 1.
	sqlStr := `select id, name, age from user where id > ?;`
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("exec %s query failed, err:%v\n", sqlStr, err)
		return
	}
	// 一定要关闭rows
	defer rows.Close()
	// 4.循环取值
	for rows.Next() {
		var u1 User
		err := rows.Scan(&u1.id, &u1.name, &u1.age)
		if err != nil {
			fmt.Printf("scan failed:%s\n", err)
			return
		}
		fmt.Printf("%#v\n", u1)
	}
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
	// 设置与数据库建立连接的最大数目，如果n大于0且小于最大闲置连接数，则最大闲置连接数
	// 会设置为最大连接数，如果n<=0,则最大连接数无限制
	db.SetMaxOpenConns(10)

	// 设置最大空闲连接数，如果n<=0,则无最大空闲连接数，如果n大于最大连接数，则最大空闲连接数为最大连接数
	// 所有的连接，不可能一直空闲，当空闲时间太长会断开，只留下一定数据的连接
	db.SetMaxIdleConns(5)
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
	//sqlStr := `select id, name, age from user where id=?`

	// 如果设置最大连接数为10，下面的循环有11次，需要11个连接，
	// 则最后一次连接会阻塞
	//for i := 1; i <= 11; i++ {
	//	// 最多查询一条记录，args表示sqlStr中的？所代表的参数
	//	// 到连接池中拿出一个连接用于查询
	//	_ = db.QueryRow(sqlStr, 1)
	//	fmt.Printf("已查询了%d次\n", i)
	//}

	queryMore(0)
}
