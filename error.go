package firebase

import (
	"github.com/gopherjs/gopherjs/js"
)

type Error struct {
	*js.Object
	Code    string `js:"code"`
	Message string `js:"message"`
	Name    string `js:"name"`
	Stack   string `js:"stack"`
}

func (e Error) Error() string {
	return "name: " + e.Name + " code:" + e.Code + " message: " + e.Message + " stack: " + e.Stack
}
