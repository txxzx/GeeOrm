package test

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/txxzx/GeeOrm"
	"testing"
)

/**
    @date: 2022/11/3
**/

//type User struct {
//	Name string `geeorm:"PRIMARY KEY"`
//	Age  int
//}

func TestSession_CreateTable(t *testing.T) {
	engine, _ := GeeOrm.NewEngine("sqlite3", "gee.db")
	defer engine.Close()
	eng := engine.NewSession()
	s := eng.Model(&User{})
	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HasTable() {
		t.Fatal("Failed to create table User")
	}
}
