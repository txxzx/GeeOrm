package test

import (
	"github.com/txxzx/GeeOrm/clause"
	"reflect"
	"testing"
)

/**
    @date: 2022/11/5
**/

type Type int

const (
	INSERT Type = iota
	VALUES
	SELECT
	LIMIT
	WHERE
	ORDERBY
)

func testSelect(t *testing.T) {
	var clause clause.Clause
	clause.Set(3, 3)
	clause.Set(2, "User", []string{"*"})
	clause.Set(4, "Name = ?", "Tom")
	clause.Set(5, "Age ASC")
	sql, vars := clause.Build(2, 4, 5, 3)
	t.Log(sql, vars)
	if sql != "SELECT * FROM User WHERE Name = ? ORDER BY Age ASC LIMIT ?" {
		t.Fatal("failed to build SQL")
	}

	if !reflect.DeepEqual(vars, []interface{}{"Tom", 3}) {
		t.Fatal("failed to build SQLVars")
	}
}

func TestClause_Build(t *testing.T) {
	t.Run("select", func(t *testing.T) {
		testSelect(t)
	})
}
