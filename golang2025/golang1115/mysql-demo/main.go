// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"time"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func initDB() (*sql.DB, error) {
// 	dsn := "root:root@tcp(localhost:3309)/cyy_acs_fe?" +
// 		"charset=utf8mb4&" + // 使用utf8mb4支持所有unicode字符
// 		"parseTime=true&" + // 解析时间字段为time.Time
// 		"multiStatements=true&" + // 解析时间字段为time.Time
// 		"loc=Asia%2FShanghai&" +
// 		"timeout=10s&" +
// 		"readTimeout=30s&" +
// 		"writeTimeout=30s"

// 	db, err := sql.Open("mysql", dsn)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// 连接池配置
// 	db.SetMaxOpenConns(100)                 // 最大打开连接数
// 	db.SetMaxIdleConns(25)                  // 最大空闲连接数
// 	db.SetConnMaxLifetime(30 * time.Minute) // 连接最大生命周期
// 	db.SetConnMaxIdleTime(10 * time.Minute) // 连接最大空闲时间

// 	// 验证连接
// 	if err := db.Ping(); err != nil {
// 		return nil, err
// 	}

// 	return db, nil
// }

// func main() {
// 	db, err := sql.Open("mysql", "root:root@/database")
// 	if err != nil {
// 		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
// 	}
// 	defer db.Close()

// 	// Prepare statement for inserting data
// 	stmtIns, err := db.Prepare("INSERT INTO squareNum VALUES( ?, ? )") // ? = placeholder
// 	if err != nil {
// 		panic(err.Error()) // proper error handling instead of panic in your app
// 	}
// 	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

// 	// Prepare statement for reading data
// 	stmtOut, err := db.Prepare("SELECT squareNumber FROM squarenum WHERE number = ?")
// 	if err != nil {
// 		panic(err.Error()) // proper error handling instead of panic in your app
// 	}
// 	defer stmtOut.Close()

// 	// Insert square numbers for 0-24 in the database
// 	for i := 0; i < 25; i++ {
// 		_, err = stmtIns.Exec(i, (i * i)) // Insert tuples (i, i^2)
// 		if err != nil {
// 			panic(err.Error()) // proper error handling instead of panic in your app
// 		}
// 	}

// 	var squareNum int // we "scan" the result in here

// 	// Query the square-number of 13
// 	err = stmtOut.QueryRow(13).Scan(&squareNum) // WHERE number = 13
// 	if err != nil {
// 		panic(err.Error()) // proper error handling instead of panic in your app
// 	}
// 	fmt.Printf("The square number of 13 is: %d", squareNum)

// 	// Query another number.. 1 maybe?
// 	err = stmtOut.QueryRow(1).Scan(&squareNum) // WHERE number = 1
// 	if err != nil {
// 		panic(err.Error()) // proper error handling instead of panic in your app
// 	}
// 	fmt.Printf("The square number of 1 is: %d", squareNum)
// }

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

// 数据库配置
const (
	userName = "root"
	password = "root"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "test"
)

// Db数据库连接池
var DB *sql.DB

type User struct {
	id    int64
	name  string
	age   int8
	sex   int8
	phone string
}

// 注意方法名大写，就是public
func InitDB() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	//打开数据库,前者是驱动名，所以要导入：_ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")
}

// 查询操作
func Query() {
	var user User
	rows, e := DB.Query("select * from user where id in (1,2,3)")
	if e == nil {
		errors.New("query incur error")
	}
	for rows.Next() {
		e := rows.Scan(user.sex, user.phone, user.name, user.id, user.age)
		if e != nil {
			fmt.Println(json.Marshal(user))
		}
	}
	rows.Close()
	DB.QueryRow("select * from user where id=1").Scan(user.age, user.id, user.name, user.phone, user.sex)

	stmt, e := DB.Prepare("select * from user where id=?")
	query, e := stmt.Query(1)
	query.Scan()
}

func DeleteUser(user User) bool {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare("DELETE FROM user WHERE id = ?")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	//设置参数以及执行sql语句
	res, err := stmt.Exec(user.id)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	//提交事务
	tx.Commit()
	//获得上一个insert的id
	fmt.Println(res.LastInsertId())
	return true
}

func InsertUser(user User) bool {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return false
	}
	//准备sql语句
	stmt, err := tx.Prepare("INSERT INTO user (`name`, `phone`) VALUES (?, ?)")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	//将参数传递到sql语句中并且执行
	res, err := stmt.Exec(user.name, user.phone)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	//将事务提交
	tx.Commit()
	//获得上一个插入自增的id
	fmt.Println(res.LastInsertId())
	return true
}

func main() {
	InitDB()
	Query()
	defer DB.Close()
}
