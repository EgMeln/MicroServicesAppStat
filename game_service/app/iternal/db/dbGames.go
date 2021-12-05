package db

import (
	"github.com/EgMeln/game_system/game_service/app/iternal/gameSystem"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

func PostgreSQL(logFight string, winHero gameSystem.Hero) string {
	db, err := sqlx.Connect("postgres", "user=postgres password=54236305 dbname=postgres")
	if err != nil {
		log.Fatalln(err)
	}
	if err := db.Ping(); err != nil {
		log.Println("Database Error", err)
	}
	tx := db.MustBegin()
	tx.MustExec("INSERT INTO heroestable (id,name,unittype,health,power,defense,mana) VALUES ($1,$2,$3,$4,$5,$6,$7)", winHero.GetID(), winHero.GetName(), winHero.GetUnitType(), winHero.Health(), winHero.GetPower(), winHero.GetDefense(), winHero.GetMana())
	tx.Commit()
	return logFight
}
