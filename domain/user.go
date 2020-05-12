package domain

// import "database/sql"

type User struct {
	ID           string      `db:"id"`
	Name         string      `db:"name"`
	// Age          int8        `db:"age"`
	Email        string      `db:email`
	Bio          string      `db:"bio"`
	Location     string      `db:"location`
	// IsActive     bool        `db:is_active`

}
