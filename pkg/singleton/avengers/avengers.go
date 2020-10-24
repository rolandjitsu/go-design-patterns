package avengers

import (
	"encoding/json"

	"github.com/rolandjitsu/go-design-patterns/pkg/singleton/db"
)

func GetAvengers(c *db.Conn) (avengers []string, err error) {
	var data []byte
	data, err = c.Query("SELECT * FROM avengers;")
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &avengers)
	return
}
