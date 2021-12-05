package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type InfoAbout struct {
	Id         int    `db:"name"`
	InfoString string `db:"outstr"`
}

func Sql(i int) string {
	db, err := sqlx.Connect("postgres", "user=postgres password=54236305 dbname=postgres")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}
	if err := db.Ping(); err != nil {
		log.Println("Database Error", err)
	}

	strings := InfoAbout{}
	err = db.Get(&strings, "SELECT * FROM strings WHERE name = $1", i)
	return strings.InfoString

}
