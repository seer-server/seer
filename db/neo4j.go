package db

import (
	"database/sql"
	"log"
	"os"

	wslog "github.com/tree-server/web-server/log"
	_ "gopkg.in/cq.v1"
)

type Neo4jCleint struct {
	*sql.DB
	log *log.Logger
}

func New() (*Neo4jClient, err) {
	db, err := sql.Open("neo4j-cypher", "http://localhost:7474")
	if err != nil {
		return nil, err
	}

	return &Neo4jClient{
		db:  db,
		log: wslog.MakeConsole("database", os.Stderr),
	}, nil
}

func (n *Neo4jClient) Close() {
	n.DB.Close()
}
