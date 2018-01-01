package firestore

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/mfindlater/firebase"
)

//CollectionReference an be used for adding documents, getting document references, and querying for documents (using the methods inherited from Query
type CollectionReference struct {
	*js.Object
	ID     string
	Parent *DocumentReference
}

//Add returns Promise containing non-nil firebase.firestore.DocumentReference
func (c *CollectionReference) Add(data interface{}) *firebase.Promise {
	return &firebase.Promise{Object: c.Call("add", data)}
}
