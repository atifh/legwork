package main

import (
        "fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"legwork/config"
//        "github.com/Pallinder/go-randomdata"

)

func main() {
	config.LoadConfig()
	fmt.Println(config.DatabaseConfig().DbName())

	db, err := sql.Open("postgres", config.DatabaseConfig().ConnectionString())
	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("Closing DB!")
}
