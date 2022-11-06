package test

import (
	"github.com/txxzx/GeeOrm/dialect"
	"github.com/txxzx/GeeOrm/schemal"
	"testing"
)

/**
    @date: 2022/11/3
**/

// schema_test.go
type User struct {
	Name string `GeeOrm:"PRIMARY KEY"`
	Age  int
}

var TestDial, _ = dialect.GetDialect("sqlite3")

// 测试schem的方法
func TestParse(t *testing.T) {
	schema := schemal.Parse(&User{}, TestDial)
	if schema.Name != "User" || len(schema.Fields) != 2 {
		t.Fatal("failed to parse User struct")
	}
	if schema.GetField("Name").Tag != "PRIMARY KEY" {
		t.Fatal("failed to parse primary key")
	}
}
