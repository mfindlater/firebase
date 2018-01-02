package firestore

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/mfindlater/firebase"
)

type CollectionReference struct {
	*js.Object
	ID     string
	Parent *DocumentReference
}

func (c *CollectionReference) Add(data interface{}) *firebase.Promise {
	return &firebase.Promise{Object: c.Call("add", data)}
}
