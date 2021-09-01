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

	设置连接池最大空闲连接数，由于数据库建立连接很消耗资源，而连接池的原理即是事先建立好一定数量的连接放着
	如果需要就直接从池子里面拿，并将其标记为忙，用完后放回，标记为空闲。这样就不用每次都建立数据库连接，节省资源
	建立连接，需要消耗的资源：内存，握手通信啥的。。。

	1.一开始调用Open函数的时候并不会创建连接，只有真正用到的时候才会建立连接，
	2.连接最长存活时间通过SetConnMaxLifetime设置, 当超过这个时间连接不会被复用，（默认值为0，即永不过期）
	3.建立一个连接后，用完不会立即释放，而是放入连接池内，标记为空闲状态，等需要的时候直接拿出来，
	4.最大空闲连接数由SetMaxIdleConns设置，当所需要的连接数超过连接池中已有连接，则会建立新的连接，
	用完后会先检查该连接在使用过程中是否已经无效，无效则不放入连接池，直接释放

	5.如果该连接用完后仍有效，则检查是否有等待连接的阻塞请求，有则直接给最早的请求，没有则判断连接池是否达到上限，未达到上限则放回连接池；
	达到上限则直接释放连接。
	6.总连接数=正在使用的连接数+空闲连接数； 总连接数通过SetMaxOpenConns设置，默认等于0（即没有限制），如果总连接数超过这个值，则会
	将连接请求放入队列中

	1.设置最大空闲连接数, 如果不调用该函数进行设置,则默认为2, 如果设置的数量 n <= 0,
	则db.maxIdle为-1,即没有可以复用的连接(连接池的最大空闲连接数为0条)

	2.当设置的最大空闲连接数超过最大连接数，则会将最大空闲连接数重置为与最大连接数相等
	3.如果当前连接池中的空闲连接数超过设置的最大空闲连接数，则会强制关闭

	sqlDB.SetMaxIdleConns(10)
	// 设置最大连接数, 当设置的最大空闲连接数超过最大连接数，则会将最大空闲连接数重置为与最大连接数相等, 默认为0，即没有限制
	sqlDB.SetMaxOpenConns(100)
	// 1.设置连接可复用的最大时间, 指的是连接池中的空闲连接如果超过这个时间就会断开连接, 默认为0, 即永远复用
	// 源码中maxLifetimeClosed是由于连接池中空闲连接超过了MaxLifetime而被强制关闭的连接数
	// maxIdleClosed是由于空闲连接数超过了连接池的最大空闲连接数而被强制关闭的连接数
	//
	// 2.核心逻辑: 假设输入参数为n, 先计算出当前时间a减去n之后的时间b,然后遍历连接池(db.freeConn)中的所有空闲连接,
	// 如果连接的创建时间在b之前,也就是说该连接从创建到现在已近经过了大于n的时间,那么会先将该连接记录在一个closing切片中,
	// 然后将记录连接池中连接的切片的最后面一个元素,移动到切片的最前面,最后去除末尾的元素,(相当于去掉首部元素然后把末尾元素移动到首部),
	// 如此循环直到连接池遍历完为止
	sqlDB.SetConnMaxLifetime(10 * time.Second)
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
