package database

import "fmt"

func (d *MongoDb) agent(fn string) string {
	return fmt.Sprintf("MongoDB.%s", fn)
}
