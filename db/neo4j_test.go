package db

import (
	"testing"
)

func TestQueryingNeo4j(t *testing.T) {
	db, err := New()
	if err != nil {
		t.Error(err)

		return
	}
	defer db.Close()

	name := "screenname"
	stmt, err := db.Prepare("CREATE (:Test {value: {0}, name: {1}})")
	if err != nil {
		t.Error(err)

		return
	}

	stmt.Exec("test value", name)

	stmt, err = db.Prepare(`
		MATCH (t:Test)
		WHERE t.value = {0}
		RETURN t.name
	`)
	if err != nil {
		t.Error(err)

		return
	}

	rows, err := stmt.Query("test value")
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

	stmt, err = db.Prepare(`
		MATCH (t:Test)
		WHERE t.value = {0}
		DELETE t
	`)
	if err != nil {
		t.Error(err)

		return
	}

	stmt.Exec("test value")

	if retName != name {
		t.Errorf("Expected %q but found %q", name, retName)
	}
}
