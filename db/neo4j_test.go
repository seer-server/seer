package db

import (
	"testing"
)

func TestQueryingNeo4j() {
	db, err := New()
	if err != nil {
		t.Error(err)

		return
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		t.Error(err)

		return
	}

	name := "screenname"
	stmt, err := tx.Prepare("CREATE (:Test {value: {0}, name: {1}})")
	if err != nil {
		t.Error(err)

		return
	}

	stmt.Exec("test value", name)

	stmt, err = tx.Prepare(`
		MATCH (t:Test)
		WHERE t.value = {0}
		RETURN t.name
	`)
	if err != nil {
		t.Error(err)

		return
	}

	rows, err := stmt.Query("some value")
	if err != nil {
		t.Error(err)

		return
	}

	var retName string
	for rows.Next() {
		err = rows.Scan(&retName)
		if err != nil {
			t.Error(err)

			return
		}
	}

	if retName != name {
		t.Errorf("Expected %q but found %q", name, retName)
	}
}
