package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	. "github.com/charly3pins/fifa-gen-api/internal"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

const (
	migrationsFolder = "cmd/migration/sqls/"
)

func createFile(fname string) {
	if _, err := os.Create(fname); err != nil {
		log.Fatal(err)
	}
}

func createCmd(timestamp int64, name string) {
	base := fmt.Sprintf("%v%v_%v.", migrationsFolder, timestamp, name)
	err := os.MkdirAll(migrationsFolder, os.ModePerm)
	if err != nil {
		log.Fatal("error creating migrations folder ", err)
	}
	createFile(base + "up.sql")
	createFile(base + "down.sql")
}

func migrationStep() int {
	limit := -1
	if flag.Arg(1) != "" {
		n, err := strconv.ParseUint(flag.Arg(1), 10, 64)
		if err != nil {
			log.Fatal("error invalid limit argument N", err)
		}
		limit = int(n)
	}

	return limit
}

func main() {
	flag.Parse()

	db, err := NewDB()
	if err != nil {
		log.Fatal(err)
	}
	db.LogMode(true)
	defer db.Close()

	dbname := "fifa" // Change DB name if need it
	postgres.DefaultMigrationsTable = fmt.Sprintf("%s_%s", dbname, postgres.DefaultMigrationsTable)

	err = db.Exec("SET search_path TO public").Error
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db.DB(), &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationsFolder,
		dbname,
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer m.Close()

	startTime := time.Now()
	switch flag.Arg(0) {
	case "up":
		limit := migrationStep()
		if limit >= 0 {
			err = m.Steps(limit)
		} else {
			err = m.Up()
		}
	case "down":
		limit := migrationStep()
		if limit >= 0 {
			err = m.Steps(-limit)
		} else {
			err = m.Down()
		}
	case "create":
		args := flag.Args()[1:]
		createFlagSet := flag.NewFlagSet("create", flag.ExitOnError)
		createFlagSet.Parse(args)

		if createFlagSet.NArg() == 0 {
			log.Fatal("Specify a name for the migration")
		}

		createCmd(startTime.Unix(), createFlagSet.Arg(0))
	default:
		log.Fatalf("Command '%s' not found", flag.Arg(0))
	}

	if err != nil {
		if err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Println(err)
	}

	log.Printf("Finished after: %s", time.Now().Sub(startTime).String())
}
