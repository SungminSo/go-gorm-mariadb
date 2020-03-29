package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// sql.DB 객체 생성
	db, err := sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 하나의 Row를 갖는 SQL 쿼리
	// 하나의 Row -> QueryRow()
	// 하나 이상의 Row -> Query()
	// 하나의 Row에서 실제 데이타를 읽어 로컬 변수에 할당 -> Scan()
	// 복수개의 Row에서 다음 Row로 이동 -> Next()
	var name string
	err = db.QueryRow("SELECT name FROM test1 WHERE id = 1").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}