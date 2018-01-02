package firestore

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/mfindlater/firebase"
)

type Client struct {
	*js.Object
}

func NewClient() *Client {
	return &Client{Object: js.Global.Get("firebase").Call("firestore")}
}

func (c *Client) Batch() *WriteBatch {
	return &WriteBatch{Object: c.Call("batch")}
}

func (c *Client) Collection(collectionPath string) *CollectionReference {
	return &CollectionReference{Object: c.Call("collection", collectionPath)}
}

func (c *Client) Doc(documentPath string) *DocumentReference {
	return &DocumentReference{Object: c.Call("doc", documentPath)}
}

func (c *Client) EnablePersistence() *firebase.Promise {
	return &firebase.Promise{Object: c.Call("enablePersistence")}
}

func (c *Client) RunTransaction(updateFunc func(*Transaction)) *firebase.Promise {
	return &firebase.Promise{Object: c.Call("enablePersistence", updateFunc)}
}

func (c *Client) SetLogLevel(logLevel string) {
	c.Call("doc", logLevel)
}

func (c *Client) Settings(settings *Settings) {
	c.Call("settings", settings)
}

type Transaction struct {
	*js.Object
}

type Settings struct {
	*js.Object
}

type WriteBatch struct {
	*js.Object
}
