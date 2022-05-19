package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Users struct {
	UsersId       int
	UsersMail     string
	UsersPassword string
	UsersPin      int
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

var db *sql.DB

func dbConnection() {
	fmt.Println("db connection")
	var err error
	db, err = sql.Open("sqlite3", "../database.db")
	checkErr(err)

	getUsers()
	//addUsers(db, "svskl@gmail.com", "Html123", 1881)

}

var userList []Users

func getUsers() {
	userList = nil
	rows, err := db.Query("Select * from users")

	if err == nil {
		for rows.Next() {
			var usersId int
			var usersMail string
			var usersPassword string
			var usersPin int

			err = rows.Scan(&usersId, &usersMail, &usersPassword, &usersPin)

			if err == nil {
				userList = append(userList, Users{UsersId: usersId, UsersMail: usersMail, UsersPassword: usersPassword, UsersPin: usersPin})

			} else {
				fmt.Println(err)
			}
		}
	} else {
		fmt.Println(err)
		return
	}

	rows.Close()
	fmt.Println(userList)

}

func addUsers(mail, pass, pin string) {
	if UserExists(mail) {
		log.Println("Mail adresi zaten kayıtlı")

	} else {
		password := make([]string, len(pass))
		for i := 0; i < len(pass); i++ {
			password[i] = string(pass[i])
		}
		temp := 0
		for _, value := range password {
			_, err := strconv.Atoi(value)
			if err == nil {
				temp++
			}
		}
		if temp != 0 {
			fmt.Println("şifre sayı içermektedir")
			_, e := strconv.Atoi(pin)
			if e == nil && string(pin[0]) != "0" && len(pin) == 4 {
				stmt, err := db.Prepare("INSERT INTO users(userMail,userPassword,userPin) VALUES(?,?,?)")
				checkErr(err)

				res, err := stmt.Exec(mail, pass, pin)
				checkErr(err)

				id, err := res.LastInsertId()
				checkErr(err)

				fmt.Println("son eklenen id: ", id)

				getUsers()
			} else {
				fmt.Println("pin kodu 4 haneli olmalıdır")
			}

		} else {
			fmt.Println("şifre en az 1 sayı içermelidir.")
		}
	}

}

func dataGet(mail, pass, p string) {

	pin, _ := strconv.Atoi(p)
	if UserExists(mail) {
		rows, err := db.Query("SELECT userId FROM users WHERE userMail IN (?)", mail)
		if err != nil {
			panic(err)
		}

		var list []Users
		for rows.Next() {
			//var usersMail string
			//var usersPassword string
			//var usersPin int
			var usersId int

			//err = rows.Scan(&usersMail, &usersPassword, &usersPin)
			err = rows.Scan(&usersId)

			if err == nil {

				//list = append(list, Users{UsersMail: usersMail, UsersPassword: usersPassword, UsersPin: usersPin})
				list = append(list, Users{UsersId: usersId})

			} else {
				fmt.Println(err)
			}
		}
		userid := list[0].UsersId
		fmt.Println("userid: ", userid)

		rows, err = db.Query("SELECT userMail, userPassword, userPin FROM users WHERE userId=?", userid)

		if err != nil {
			panic(err)
		}

		for rows.Next() {
			var usersMail string
			var usersPassword string
			var usersPin int

			err = rows.Scan(&usersMail, &usersPassword, &usersPin)

			if err == nil {

				list = append(list, Users{UsersMail: usersMail, UsersPassword: usersPassword, UsersPin: usersPin})

			} else {
				fmt.Println(err)
			}
		}

		fmt.Println(list)
		if mail == list[1].UsersMail && pass == list[1].UsersPassword && pin == list[1].UsersPin {
			log.Println("Giriş yaptınız")

		} else {
			log.Println("Hatalı giriş yaptınız")

		}

		rows.Close()
	} else {
		log.Println("Hatalı giriş yaptınız")
	}

}

func UserExists(mail string) bool {
	sqlStmt := `SELECT userMail FROM users WHERE userMail = ?`
	//_,err := db.Query("SELECT userMail FROM users WHERE userMail = ?", mail)
	err := db.QueryRow(sqlStmt, mail).Scan(&mail)
	if err != nil {
		if err != sql.ErrNoRows {
			// a real error happened! you should change your function return
			// to "(bool, error)" and return "false, err" here
			log.Print(err)
		}

		return false
	}

	return true

}
