package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	// "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type mysqlConfig struct {
	host     string
	user     string
	password string
	port     string
	db       string
}

// type config struct{
// 	host     "localhost"
// 	port     = "2021"
// 	user     = "root"
// 	password = "password"
// 	dbname   = "student"
// 	ssl      = "require"
// }

func Connect() *sql.DB {
	conf := mysqlConfig{host: "localhost",
		user:     "root",
		password: "password",
		port:     "2021",
		db:       "college",
	}

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", conf.user, conf.password, conf.host, conf.port, conf.db)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		fmt.Println(err)
	}
	return db
}
