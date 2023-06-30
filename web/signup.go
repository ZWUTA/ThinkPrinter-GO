package web

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	. "thinkPrinter/database"
	"thinkPrinter/entity"
)

func SignUp(c *gin.Context) {
	// 数据库读取user
	var user entity.User
	// JSON解析到user
	var userDTO entity.User

	// 读取请求体中的数据
	err := json.NewDecoder(c.Request.Body).Decode(&userDTO)
	if err != nil {
		log.Panicln(err)
	}

	result := DB.Where("username = ?", userDTO.Username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 用户不存在，可以注册
		} else {
			// 其他错误
			log.Panicln(result.Error)
		}
	} else {
		// 用户已存在
		c.String(http.StatusConflict, "用户已存在")
		if err != nil {
			log.Panicln(err)
		}
		return
	}

	//userCredentials := UserCredentials{}
	//// 读取请求体中的数据
	//err := json.NewDecoder(r.Body).Decode(&userCredentials)
	//if err != nil {
	//	panic(err)
	//}
	//
	//db, err := database.InitDB()
	//defer database.CloseDB(db)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// 检查用户是否已经存在
	//// 判断Username或SNO哪一个不为空，然后执行不同的sql语句，如果都有值，使用sno， 因为sno是唯一的, 如果都为空，返回错误
	//if userCredentials.Username == "" && userCredentials.SNO == "" {
	//	_, err := fmt.Fprintf(w, "用户名和学号不能为空")
	//	if err != nil {
	//		panic(err)
	//	}
	//	return
	//} else if userCredentials.Username == "" {
	//	sqlStmt := `SELECT count(*) FROM users WHERE sno=?;`
	//	rows, err := db.Query(sqlStmt, userCredentials.SNO)
	//	defer func(rows *sql.Rows) {
	//		err := rows.Close()
	//		if err != nil {
	//			panic(err)
	//		}
	//	}(rows)
	//
	//	if err != nil {
	//		panic(err)
	//	}
	//	if rows.Next() {
	//		var count int
	//		err := rows.Scan(&count)
	//		if err != nil {
	//			panic(err)
	//		}
	//		if count != 0 {
	//			_, err := fmt.Fprintf(w, "学号已存在")
	//			if err != nil {
	//				panic(err)
	//			}
	//			return
	//		}
	//	}
	//	err = rows.Err()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	err = rows.Close()
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//} else {
	//	sqlStmt := `SELECT count(*) FROM users WHERE username=?;`
	//	rows, err := db.Query(sqlStmt, userCredentials.Username)
	//	defer func(rows *sql.Rows) {
	//		err := rows.Close()
	//		if err != nil {
	//			panic(err)
	//		}
	//	}(rows)
	//	if err != nil {
	//		panic(err)
	//	}
	//	if rows.Next() {
	//		var count int
	//		err := rows.Scan(&count)
	//		if err != nil {
	//			panic(err)
	//		}
	//		if count != 0 {
	//			_, err := fmt.Fprintf(w, "用户名已存在")
	//			if err != nil {
	//				panic(err)
	//			}
	//			// 如果用户名已存在，返回错误
	//			return
	//		}
	//	}
	//	err = rows.Err()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	err = rows.Close()
	//	if err != nil {
	//		panic(err)
	//	}
	//}
	//
	//// 将用户密码加密
	//userCredentials.Password = tools.Encrypt(userCredentials.Password)
	//marshal, err := json.Marshal(userCredentials)
	//if err != nil {
	//	panic(err)
	//}
	//log.Println(string(marshal))
	////	注册用户
	//sqlStmt := `INSERT INTO users (username, password, sname, sno) VALUES (?, ?, ?, ?);`
	//_, err = db.Exec(sqlStmt, userCredentials.Username, userCredentials.Password, userCredentials.SName, userCredentials.SNO)
	//if err != nil {
	//	log.Println(err)
	//	_, errr := fmt.Fprintf(w, "注册失败, 请重试:%s", err)
	//	if errr != nil {
	//		log.Fatal(err)
	//	}
	//	panic(err)
	//}
	//
	//_, err = fmt.Fprintf(w, "注册成功")
	//if err != nil {
	//	panic(err)
	//}

}
