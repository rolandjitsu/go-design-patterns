package singleton

import (
	"errors"

	"github.com/rolandjitsu/go-design-patterns/pkg/singleton/avengers"
	"github.com/rolandjitsu/go-design-patterns/pkg/singleton/db"
)

var c *db.Conn

func init() {
	c = &db.Conn{}
}

func RunEndgame() (err error) {
	var members []string
	members, err = avengers.GetAvengers(c)
	if err != nil {
		return err
	}

	if len(members) != 5 {
		err = errors.New("Not enough members")
	}

	return
}
