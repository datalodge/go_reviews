package cassandra

import (
	"fmt"
	"time"

	"github.com/gocql/gocql"
)

// Session for cassandra connection
var Session *gocql.Session

func init() {
	connect()
}

func connect() {
	var err error
	cluster := gocql.NewCluster("ec2-54-193-62-206.us-west-1.compute.amazonaws.com")
	cluster.Keyspace = "data_lodge"
	cluster.Consistency = gocql.One
	Session, err = cluster.CreateSession()
	if err != nil {
		// Must be because the db is not ready yet
		time.Sleep(10000 * time.Millisecond)
		connect()
	}

	fmt.Println("cassandra init done!")
}
