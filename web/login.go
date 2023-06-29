package web

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"thinkPrinter/database"
	"thinkPrinter/tools"
)

// UserCredentials 仅用于登录和注册
type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	SNO      string `json:"sno"`
	SName    string `json:"name"`
}

// User 用于登录后的用户信息
type User struct {
	UID      int     `json:"uid"`
	SNO      string  `json:"sno"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	SName    string  `json:"sname"`
	Balance  float64 `json:"balance"`
	VIP      int     `json:"vip"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	tools.OutputLog(r)
	// 检查请求方法. 如果不是POST, 返回405
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	loginData := UserCredentials{}
	// 读取请求体中的数据
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		panic(err)
	}
	db, err := database.GetDB()
	defer database.CloseDB(db)

	if err != nil {
		panic(err)
	}
	// 从数据库中读取用户信息
	user := User{}

	if loginData.Username == "" || loginData.Password == "" {
		// 前端返回了空的用户名或密码，大概是想搞事情？
		_, err := fmt.Fprintf(w, "？？？？？？？")
		if err != nil {
			panic(err)
		}
		return
	} else if loginData.Username != "" {
		log.Printf("根据用户名登录: %s\n", loginData.Username)
		// 根据用户名登录
		sqlStmt := `SELECT uid, sno, username, password, sname, balance, vip FROM users WHERE username=? ;`
		rows, err := db.Query(sqlStmt, loginData.Username)
		if err != nil {
			panic(err)
		}
		for rows.Next() {
			err := rows.Scan(&user.UID, &user.SNO, &user.Username, &user.Password, &user.SName, &user.Balance, &user.VIP)
			if err != nil {
				panic(err)
			}
		}
		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {
				panic(err)
			}
		}(rows)
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Printf("根据学号登录: %s\n", loginData.SNO)
		// 根据学号登录
		sqlStmt := `SELECT uid, sno, username, password, sname, balance, vip FROM users WHERE sno=? ;`
		rows, err := db.Query(sqlStmt, loginData.SNO)
		if err != nil {
			panic(err)
		}
		for rows.Next() {
			err := rows.Scan(&user.UID, &user.SNO, &user.Username, &user.Password, &user.SName, &user.Balance, &user.VIP)
			if err != nil {
				panic(err)
			}
		}
		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {
				panic(err)
			}
		}(rows)
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}

	// 将用户输入的密码加密
	loginData.Password = tools.Encrypt(loginData.Password)

	// 如果用户名和密码正确, 返回登录成功
	if user.Username == loginData.Username && user.Password == loginData.Password {
		_, err := fmt.Fprintf(w, "登录成功")
		if err != nil {
			panic(err)
		}
	} else {
		_, err := fmt.Fprintf(w, "用户名不存在或密码错误")
		if err != nil {
			panic(err)
		}
	}
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	tools.OutputLog(r)
	// 检查请求方法. 如果不是POST, 返回405
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	userCredentials := UserCredentials{}
	// 读取请求体中的数据
	err := json.NewDecoder(r.Body).Decode(&userCredentials)
	if err != nil {
		panic(err)
	}

	db, err := database.GetDB()
	defer database.CloseDB(db)
	if err != nil {
		panic(err)
	}

	// 检查用户是否已经存在
	// 判断Username或SNO哪一个不为空，然后执行不同的sql语句，如果都有值，使用sno， 因为sno是唯一的, 如果都为空，返回错误
	if userCredentials.Username == "" && userCredentials.SNO == "" {
		_, err := fmt.Fprintf(w, "用户名和学号不能为空")
		if err != nil {
			panic(err)
		}
		return
	} else if userCredentials.Username == "" {
		sqlStmt := `SELECT count(*) FROM users WHERE sno=?;`
		rows, err := db.Query(sqlStmt, userCredentials.SNO)
		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {
				panic(err)
			}
		}(rows)

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
				_, err := fmt.Fprintf(w, "学号已存在")
				if err != nil {
					panic(err)
				}
				return
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

	} else {
		sqlStmt := `SELECT count(*) FROM users WHERE username=?;`
		rows, err := db.Query(sqlStmt, userCredentials.Username)
		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {
				panic(err)
			}
		}(rows)
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
				_, err := fmt.Fprintf(w, "用户名已存在")
				if err != nil {
					panic(err)
				}
				// 如果用户名已存在，返回错误
				return
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
	}

	// 将用户密码加密
	userCredentials.Password = tools.Encrypt(userCredentials.Password)
	marshal, err := json.Marshal(userCredentials)
	if err != nil {
		panic(err)
	}
	log.Println(string(marshal))
	//	注册用户
	sqlStmt := `INSERT INTO users (username, password, sname, sno) VALUES (?, ?, ?, ?);`
	_, err = db.Exec(sqlStmt, userCredentials.Username, userCredentials.Password, userCredentials.SName, userCredentials.SNO)
	if err != nil {
		log.Println(err)
		_, errr := fmt.Fprintf(w, "注册失败, 请重试:%s", err)
		if errr != nil {
			log.Fatal(err)
		}
		panic(err)
	}

	_, err = fmt.Fprintf(w, "注册成功")
	if err != nil {
		panic(err)
	}

}
