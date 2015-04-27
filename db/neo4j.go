package db

import (
	"database/sql"
	"os"

	"github.com/tree-server/trees/log"
	_ "gopkg.in/cq.v1"
)

type Neo4jClient struct {
	*sql.DB
	log *log.Logger
}

func New() (*Neo4jClient, error) {
	db, err := sql.Open("neo4j-cypher", "http://localhost:7474")
	if err != nil {
		return nil, err
	}

	return &Neo4jClient{
		DB:  db,
		log: log.Make("database", ":stdout:"),
	}, nil
}

func (n *Neo4jClient) Close() {
	n.DB.Close()
}
