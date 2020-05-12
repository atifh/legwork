package main

import (
        "fmt"
	"encoding/json"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"legwork/config"
	"legwork/domain"
	"github.com/fatih/structs"
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

// updates tsvector values (name, email, bio, location) for the given user
func updateTSVector(id string) {
	sqlStatement := `UPDATE users SET tsv =	setweight(to_tsvector(name), 'A')
 || setweight(to_tsvector(email), 'B') || setweight(to_tsvector(location), 'C')
 || setweight(to_tsvector(bio), 'D') WHERE id=$1`

	_, err := DB.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}
	fmt.Println("tvs vector updated for User ID:", id)
}

// creates dummy users and dumps into User table
func createDummyUsers(count int) {
	sqlStatement := `INSERT INTO users (name, email, age, bio, location)  VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var id string

	for i := 0; i < count; i++ {
		err := DB.QueryRow(sqlStatement,
			randomdata.SillyName(), randomdata.Email(),
			randomdata.Number(18, 70), randomdata.Paragraph(), randomdata.City()).Scan(&id)

		if err != nil {
			panic(err)
		}
		fmt.Println("New User ID is:", id)
		updateTSVector(id)
	}
}

// migrates PostgreSQL DB by referring ./migrations files
// for new changes
func migrateDB() {
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
}

// finds users matching the given string from Users table
func searchUsers(searchString string) (searchResults []map[string]interface{}) {
	sqlStatement := `SELECT id, name, email, location, bio FROM users, plainto_tsquery($1) q WHERE tsv @@ q;`
	var id, name, email, location, bio string
	// searchResults := []domain.User{}

	rows, err := DB.Query(sqlStatement, searchString)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name, &email, &location, &bio)
		if err != nil {
			panic(err)
		}
		user := domain.User {
			ID: id,
			Name: name,
			Email: email,
			Bio: bio,
			Location: location}

		searchResults = append(searchResults, structs.Map(user))

	}
	return
}

func main() {
	config.LoadConfig()
	connectDB()
	migrateDB()
	searchString := "dog blood"

	searchResults := searchUsers(searchString)
	resultJson, _ := json.MarshalIndent(searchResults, "", "    ")
	fmt.Printf("\nFound %d Search Results for %s\n\n", len(searchResults), searchString)
	fmt.Println(string(resultJson))


	// FIXME: Should run this only when there is no user in the DB
	// count := 5
	// fmt.Printf("Creating %d dummy users\n", count)
	// createDummyUsers(count)

	defer DB.Close()

	fmt.Println("Closing DB!")
}
