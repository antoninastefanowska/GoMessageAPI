package Cassandra

import (
	"github.com/gocql/gocql"
	"fmt"
)

var Session *gocql.Session

const CLUSTER_URL = "127.0.0.1"
const CLUSTER_NAME = "messageapi"

func init() {
	var err error

	cluster := gocql.NewCluster(CLUSTER_URL)
	cluster.Keyspace = CLUSTER_NAME
	Session, err = cluster.CreateSession()

	if err != nil {
		panic(err)
	}
	
	fmt.Println("Cassandra initialized successfully.")
}