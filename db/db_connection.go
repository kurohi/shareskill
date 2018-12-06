package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/*Creating the mysql connection*/
func CreateCon() *sql.DB{
	db, err := sql.Open("mysql", "shareskill:skillshare@tcp(127.0.0.1:3306)/shareskill")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}

	//defer db.Close()
	//making sure the db is connected
	err = db.Ping()
	fmt.Println(err)
	if err!=nil {
		fmt.Println("db is not connected")
		fmt.Println(err.Error())
	}
	return db
}
