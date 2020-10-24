package db

import (
	"encoding/json"
)

type Conn struct{}

func (c *Conn) Query(q string) (data []byte, err error) {
	avengers := []string{"Iron Man", "Thor", "Wasp", "Ant-Man", "Hulk"}
	data, err = json.Marshal(&avengers)
	return
}
