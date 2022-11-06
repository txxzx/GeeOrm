package test

/**
    @date: 2022/11/6
**/

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/txxzx/GeeOrm"
	"github.com/txxzx/GeeOrm/session"
	"testing"
)

var (
	user1 = &User{"Tom", 18}
	user2 = &User{"Sam", 25}
	user3 = &User{"Jack", 25}
)

func testRecordInit(t *testing.T) *session.Session {
	t.Helper()
	engine, _ := GeeOrm.NewEngine("sqlite3", "gee.db")
	defer engine.Close()
	eng := engine.NewSession()
	s := eng.Model(&User{})
	err1 := s.DropTable()
	err2 := s.CreateTable()
	_, err3 := s.Insert(user1, user2)
	if err1 != nil || err2 != nil || err3 != nil {
		t.Fatal("failed init test records")
	}
	return s
}

func TestSession_Insert(t *testing.T) {
	s := testRecordInit(t)
	affected, err := s.Insert(user3)
	if err != nil || affected != 1 {
		t.Fatal("failed to create record")
	}
}
func TestSession_Find(t *testing.T) {
	s := testRecordInit(t)
	var users []User
	if err := s.Find(&users); err != nil || len(users) != 2 {
		t.Fatal("failed to query all")
	}
}
