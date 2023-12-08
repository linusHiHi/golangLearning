package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //要导入相应驱动包，否则会报错
	"github.com/jmoiron/sqlx"
)

// 定义一个全局对象db
var db *sqlx.DB

func initDB() {
	var err error

	var dsn = "root:qwq233383@tcp(127.0.0.1:3306)/students?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}

	fmt.Println("connecting to MySQL...")
	return
}

func insertDB(name string, age int, sex string) {
	sqlStr := "insert into student(name, age, sex) values (?,?,?)"
	ret, err := db.Exec(sqlStr, name, age, sex)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// 删除数据
func deleteDB() {
	sqlStr := "delete from student where id > ?"
	ret, err := db.Exec(sqlStr, 0)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

type User struct {
	Id   int
	Name string
	Age  int
	Sex  string
}

func queryMultiRowDemo() {
	sqlStr := "select id, name, age from student where id > ?"
	var users []User
	err := db.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", users)
}

func main() {

	//初始化连接
	initDB()
	deleteDB()
	insertDB("小明", 18, "男")
	insertDB("小王", 18, "男")
	insertDB("小李", 18, "男")
	insertDB("小红", 18, "女")
	insertDB("小刚", 18, "男")
	insertDB("小壮", 18, "男")
	insertDB("小勇", 18, "男")
	insertDB("小白", 18, "男")
	insertDB("小天", 18, "男")
	insertDB("小汪", 18, "男")
	queryMultiRowDemo()

}
