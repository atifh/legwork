package main

import (
        "fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"legwork/config"
       "github.com/Pallinder/go-randomdata"

)

// DB is a global variable to hold db connection
var DB *sql.DB

// ConnectDB opens a connection to the database
func connectDB() {
	db, err := sql.Open("postgres", config.DatabaseConfig().ConnectionString())
	if err != nil {
		panic(err.Error())
	}

	DB = db
}

// creates dummy users and dumps into User table
func createDummyUsers(count int) {
	sqlStatement := `INSERT INTO users (name, email, age, bio, location)  VALUES ($1, $2, $3, $4, $5) RETURNING id`
	// id := 0
	var id string

	for i := 0; i < count; i++ {
		err := DB.QueryRow(sqlStatement,
			randomdata.SillyName(), randomdata.Email(),
			randomdata.Number(18, 70), randomdata.Paragraph(), randomdata.City()).Scan(&id)

		if err != nil {
			panic(err)
		}
		fmt.Println("New User ID is:", id)
	}
}

func main() {
	config.LoadConfig()
	connectDB()

	driver, err := postgres.WithInstance(DB, &postgres.Config{})
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

	// FIXME: Should run this only when there is no user in the DB
	fmt.Println("Creating dummy users")
	createDummyUsers(5)

	defer DB.Close()

	fmt.Println("Closing DB!")
}
