package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	// mysql 드라이버를 사용하여, 데이터베이스에 접속
	// 사용자: testuser, 비밀번호: testpass, URL
	db, err := gorm.Open("mysql", "testuser:testpass@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var version string
	db.Exec("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)

	dbName := "test2"	// 생성할 데이터베이스 이름

	// Raw SQL로 데이터베이스 생성
	db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	db.Exec("commit;")
}