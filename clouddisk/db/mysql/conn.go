package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/clouddisk?charset=utf8")
	db.SetMaxOpenConns(1000)
	err := db.Ping()
	if err != nil {
		fmt.Println("failed conn to mysql: ", err.Error())
		os.Exit(1)
	}
}

// DBConn返回数据连接对象
func DBConn() *sql.DB {
	return db
}
