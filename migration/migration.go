package main

import (
	"database/sql"
	"fmt"

	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/mysql"
	_ "github.com/mattes/migrate/source/file"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/manga")
	if err != nil {
		fmt.Println(err)
		return
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://sql",
		"mysql",
		driver,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = m.Up()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("success")
}
