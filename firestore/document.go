package firestore

import "github.com/gopherjs/gopherjs/js"

//DocumentReference refers to a document location in a Firestore database and can be used to write, read, or listen to the location. The document at the referenced location may or may not exist. A DocumentReference can also be used to create a CollectionReference to a subcollection.
type DocumentReference struct {
	*js.Object
}

type DocumentSnapshot struct {
	*js.Object
}
