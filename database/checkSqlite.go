package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func GetDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./sqlite.db")

	if err != nil {
		return nil, err
	}
	//db.SetMaxOpenConns(1) // Limit to one concurrent connection
	log.Println("数据库连接成功")
	return db, nil
}
func CloseDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		panic(err)
	}
	log.Println("数据库连接关闭")
}

func CheckSqlite() error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	defer CloseDB(db)
	//检查sqlite是否存在user表
	return checkTable(db)
}
func initSqlite(db *sql.DB) error {

	//读取init.database
	sqlFile, err := os.ReadFile("./init/init.sql")
	if err != nil {
		log.Fatal("初始化数据库失败，找不到init.sql文件", err)
		return err
	}
	sqlStmt := string(sqlFile)

	//create table
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return err
	}
	return nil
}
func checkTable(db *sql.DB) error {
	//https://beets.io/blog/sqlite-nightmare.html
	//检查sqlite是否存在users表
	sqlStmt := `SELECT count(*) FROM sqlite_master WHERE type='table' AND name='users';`
	var count int
	err := db.QueryRow(sqlStmt).Scan(&count)
	if err != nil {
		panic(err)
	}
	// count == 0 说明不存在users表
	// 如果存在users表，则认为已经初始化过了
	// 初始化sqlite
	if count == 0 {
		return initSqlite(db)
	}
	return nil
}
