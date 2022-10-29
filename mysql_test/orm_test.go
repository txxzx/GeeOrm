package mysql_test

import (
	"database/sql"
	"log"
	"testing"
)

/**
    @date: 2022/10/29
**/

func Test(t *testing.T) {
	db, _ := sql.Open("splite3", "gee.db")
	defer func() { _ = db.Close() }()
	_, _ = db.Exec("DROP TABLE IF EXISTS User;")
	_, _ = db.Exec("CREATE TABLE User(Name text);")
	result, err := db.Exec("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam")
	if err != nil {
		affected, _ := result.RowsAffected()
		log.Println(affected)
		row := db.QueryRow("SELECT Name FROM User LIMIT 1")
		//rows,err :=db.Query("SELECT Name FROM User LIMIT 1")
		//if err!=nil{
		//	log.Println(err)
		//}
		//rows.Next()
		var name string
		if err := row.Scan(&name); err != nil {
			log.Println(name)
		}
	}
}
