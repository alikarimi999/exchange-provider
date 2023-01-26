package database

import "fmt"

func (d *mongoDb) agent(fn string) string {
	return fmt.Sprintf("MongoDB.%s", fn)
}
