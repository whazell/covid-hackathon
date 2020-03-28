package covid

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// Use global right now for ease of use
var db *sql.DB

func ConnectDatabase(c DatabaseConfig) error {
	var err error
	cstr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Name)
	db, err = sql.Open("mysql", cstr)
	if err != nil {
		return err
	}
	return nil
}

// Return the database connection incase it needs to be used outside the package
func DatabaseConnection() *sql.DB {
	return db
}
