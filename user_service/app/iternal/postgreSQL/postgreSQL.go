package postgreSQL

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"strconv"
)

type UserDB struct {
	Id       int    `db:"id"`
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

func GetUserByID(id string) string {
	intId, _ := strconv.Atoi(id)
	db, err := sqlx.Connect("postgres", "user=postgres password=54236305 dbname=postgres")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}
	if err := db.Ping(); err != nil {
		log.Println("Database Error", err)
	}
	user := UserDB{}
	err = db.Get(&user, "SELECT * FROM users WHERE id = $1", intId)
	return user.Username + " " + user.Email
}

func GetAllUser() string {
	db, err := sqlx.Connect("postgres", "user=postgres password=54236305 dbname=postgres")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}
	if err := db.Ping(); err != nil {
		log.Println("Database Error", err)
	}
	user := UserDB{}
	rows, err := db.Queryx("SELECT * FROM users")
	var str string
	for rows.Next() {
		err := rows.StructScan(&user)
		if err != nil {
			log.Fatalln(err)
		}
		str += user.Username + " " + user.Email
	}
	return str
}

func AddUser(name, email, pass string) string {
	db, err := sqlx.Connect("postgres", "user=postgres password=54236305 dbname=postgres")
	if err != nil {
		log.Fatalln(err)
	}
	if err := db.Ping(); err != nil {
		log.Println("Database Error", err)
	}
	tx := db.MustBegin()
	tx.MustExec("INSERT INTO users (username,email,password) VALUES ($1,$2,$3)", name, email, pass)
	tx.Commit()
	return "Success"
}
func DeleteUser(nameDel string) string {
	db, err := sqlx.Connect("postgres", "user=postgres password=54236305 dbname=postgres")
	if err != nil {
		log.Fatalln(err)
	}
	if err := db.Ping(); err != nil {
		log.Println("Database Error", err)
	}
	tx := db.MustBegin()
	tx.MustExec("DELETE FROM users WHERE username = $1", nameDel)
	tx.Commit()
	return "Success"
}
