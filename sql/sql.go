package sql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectSql() (db *sql.DB, err error) {
	fmt.Println("hello")

	DB, err := sql.Open("mysql", "csmysql:Chenshuo711@(sh-cynosdbmysql-grp-i3oj2lcm.sql.tencentcdb.com:28580)/comment?charset=utf8")
	// defer DB.Close()
	err = DB.Ping()
	return DB, err

	// var answer int
	// rows, err := DB.Query("select count(*) from mhyComment")
	// if err != nil {
	// 	fmt.Println("error!!!")
	// 	fmt.Println(err)
	// }
	// rows.Next()
	// fmt.Println("success!!")
	// rows.Scan(&answer)
	// fmt.Println(answer)

}

func QueryRow(db *sql.DB) (rows *sql.Rows, err error) {
	str := "select count(*) from mhyComment"
	rows, err = db.Query(str)
	return rows, err
}
