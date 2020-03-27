package internal

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type connectionOptions struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	Schema   string
	SslMode  string
}

func NewDB() (*gorm.DB, error) {
	return newConnection(
		connectionOptions{
			Host:     "127.0.0.1",
			Port:     "5431",
			User:     "fifa_gen_dev",
			Password: "fifa_gen_dev",
			DbName:   "fifa",
			SslMode:  "disable",
			Schema:   "generator",
		},
	)
}

func newConnection(c connectionOptions) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host="+c.Host+" port="+c.Port+" user="+c.User+" password="+c.Password+" dbname="+c.DbName+" sslmode="+c.SslMode)

	if err == nil {
		db.DB().SetMaxOpenConns(100)
		db.SingularTable(true)

		schema := strings.TrimSpace(c.Schema)
		if len(schema) > 0 {
			db.Exec(fmt.Sprintf("SET SEARCH_PATH to %s", schema))
		}

		db.LogMode(false) // Set to TRUE if wanna debug
	}

	return db, err
}
