package firestore

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/mfindlater/firebase"
)

//Firestore Cloud Firestore service interface.
type Client struct {
	*js.Object
}

func NewClient() *Client {
	return &Client{Object: js.Global.Get("firebase").Call("firestore")}
}

//Batch returns WriteBatch.
func (c *Client) Batch() *WriteBatch {
	return &WriteBatch{Object: c.Call("batch")}
}

//Collection returns a CollectionReference.
func (c *Client) Collection(collectionPath string) *CollectionReference {
	return &CollectionReference{Object: c.Call("collection", collectionPath)}
}

//Doc gets a DocumentReference instance that refers to the document at the specified path.
func (c *Client) Doc(documentPath string) *DocumentReference {
	return &DocumentReference{Object: c.Call("doc", documentPath)}
}

/*
EnablePersistence returns Promise.

Attempts to enable persistent storage, if possible.

Must be called before any other methods (other than Settings()).

If this fails, EnablePersistence() will reject the promise it returns. Note that even after this failure, the firestore instance will remain usable, however offline persistence will be disabled.

There are several reasons why this can fail, which can be identified by the code on the error.

failed-precondition: The app is already open in another browser tab.
unimplemented: The browser is incompatible with the offline persistence implementation.
*/
func (c *Client) EnablePersistence() *firebase.Promise {
	return &firebase.Promise{Object: c.Call("enablePersistence")}
}

// RunTransaction executes the given updateFuncand then attempts to commit the changes applied within the transaction. If any document read within the transaction has changed, Cloud Firestore retries the updateFunc. If it fails to commit after 5 attempts, the transaction fails.
func (c *Client) RunTransaction(updateFunc func(*Transaction)) *firebase.Promise {
	return &firebase.Promise{Object: c.Call("enablePersistence", updateFunc)}
}

//SetLogLevel sets the verbosity of Cloud Firestore logs (debug, error, or silent).
func (c *Client) SetLogLevel(logLevel string) {
	c.Call("doc", logLevel)
}

//Settings specifies custom settings to be used to configure the Firestore instance. Must be set before invoking any other methods.
func (c *Client) Settings(settings *Settings) {
	c.Call("settings", settings)
}

//Transaction a reference to a transaction.
type Transaction struct {
	*js.Object
}

//Settings Specifies custom configurations for your Cloud Firestore instance. You must set these before invoking any other methods.
type Settings struct {
	*js.Object
}

/*
WriteBatch A write batch, used to perform multiple writes as a single atomic unit.

A WriteBatch object can be acquired by calling the Firestore.Batch() method. It provides methods for adding writes to the write batch. None of the writes are committed (or visible locally) until WriteBatch.Commit() is called.

Unlike transactions, write batches are persisted offline and therefore are preferable when you don't need to condition your writes on read data.
*/
type WriteBatch struct {
	*js.Object
}
