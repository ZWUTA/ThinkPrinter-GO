package test

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

//	func TestSqliteDemo(t *testing.T) {
//		//err := database.CheckSqlite()
//		//if err != nil {
//		//	return
//		//}
//	}
type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	SNO      string `json:"sno"`
	SName    string `json:"name"`
}

func TestSqliteDemo2(t *testing.T) {
	db, err := sql.Open("sqlite3", "../sqlite.db")

	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(1) // Limit to one concurrent connection
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)
	log.Println("数据库连接成功")

	userCredentials := UserCredentials{
		Username: "admin",
		Password: "admin",
		SNO:      "2018210000",
	}

	sqlStmt := `SELECT count(*) FROM users WHERE sno=?;`
	rows, err := db.Query(sqlStmt, userCredentials.SNO)

	if err != nil {
		panic(err)
	}
	if rows.Next() {
		var count int
		err := rows.Scan(&count)
		if err != nil {
			panic(err)
		}
		if count != 0 {
			log.Println("学号已存在")
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	err = rows.Close()
	if err != nil {
		panic(err)
	}

	sqlStmt = `INSERT INTO users (username, password, sname, sno) VALUES (?, ?, ?, ?);`
	_, err = db.Exec(sqlStmt, userCredentials.Username, userCredentials.Password, userCredentials.SName, userCredentials.SNO)
	log.Println("注册成功")
	log.Println(err)
	if err != nil {
		log.Println(err)
		log.Println("用户名已存在")
		panic(err)
	}
}

func TestLock(t *testing.T) {
	//TestSqliteDemo2(t)
	for i := 0; i < 10; i++ {
		TestSqliteDemo2(t)
	}
}
