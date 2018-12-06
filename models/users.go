package models

import (
	_ "database/sql"
	"github.com/kurohi/shareskill/db"
	"fmt"
	"database/sql"
)

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Dept  string `json:"dept"`
	Photo string `json:"photo"`
}

type Users struct {
	Users []User `json:"user"`
}

var con *sql.DB

func GetUsers() Users {
	con := db.CreateCon()

	sqlStatement := "SELECT id,name,dept,photo FROM users ORDER BY id;"

	rows, err := con.Query(sqlStatement)
	fmt.Println(rows)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	results := Users{}

	for rows.Next() {
		user := User{}
		err2 := rows.Scan(&user.Id, &user.Name, &user.Dept, &user.Photo)
		if err2 != nil {
			fmt.Println(err2)
		}
		results.Users = append(results.Users, user)
	}
	return results
}
