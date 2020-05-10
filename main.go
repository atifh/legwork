package main

import (
        "fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"legwork/config"
//        "github.com/Pallinder/go-randomdata"

)

// func createDummyUsers() {
// 	sqlStatement := `INSERT INTO users (name, email, age, bio, location)  VALUES ($1, $2, $3, $4) RETURNING id`
// 	id := 0
// 	err = db.QueryRow(sqlStatement, 30, "jon@calhoun.io", "Jonathan", "Calhoun").Scan(&id)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("New record ID is:", id)
// }

func main() {
	config.LoadConfig()

	db, err := sql.Open("postgres", config.DatabaseConfig().ConnectionString())
	if err != nil {
		panic(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	migration, err := migrate.NewWithDatabaseInstance(
		"file://migrations", config.DatabaseConfig().DbName(), driver)

	if err != nil {
		panic(err)
	}

	fmt.Println("Applying database migrations!")
	err = migration.Up()
	if err != nil && err != migrate.ErrNoChange {
		panic(err)
	}
	fmt.Println("Ran all migrations")

	defer db.Close()
	fmt.Println("Closing DB!")
}
