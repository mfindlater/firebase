package firebase

import (
	"github.com/gopherjs/gopherjs/js"
)

type Error struct {
	*js.Object
}

func (e Error) Error() string {
	return e.String()
}
