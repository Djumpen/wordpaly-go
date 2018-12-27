package main

import (
	"fmt"

	"github.com/djumpen/wordplay-go/config"
	"github.com/djumpen/wordplay-go/mysqldb"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	cfg := config.ReadConfig()
	db := mysqldb.New(cfg.DB)

	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}
	n, err := migrate.Exec(db.DB, "mysql", migrations, migrate.Up)
	if err != nil {
		panic(err)
	}
	fmt.Println("-----------------------")
	fmt.Printf("Applied %d migrations!\n", n)
	fmt.Println("-----------------------")
}
