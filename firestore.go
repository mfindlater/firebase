package firebase

import (
	"github.com/gopherjs/gopherjs/js"
)

type Firestore struct {
	*js.Object
}

func (f *Firestore) Batch() *WriteBatch {
	return &WriteBatch{Object: f.Call("batch")}
}

func (f *Firestore) Collection(collectionPath string) *CollectionReference {
	return &CollectionReference{Object: f.Call("collection", collectionPath)}
}

func (f *Firestore) Doc(documentPath string) *DocumentReference {
	return &DocumentReference{Object: f.Call("doc", documentPath)}
}

func (f *Firestore) EnablePersistence() *Promise {
	return &Promise{Object: f.Call("enablePersistence")}
}

func (f *Firestore) RunTransaction(updateFunc func(*Transaction)) *Promise {
	return &Promise{Object: f.Call("enablePersistence", updateFunc)}
}

func (f *Firestore) SetLogLevel(logLevel string) {
	f.Call("doc", logLevel)
}

func (f *Firestore) Settings(settings *Settings) {
	f.Call("settings", settings)
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

type DocumentReference struct {
	*js.Object
}

type DocumentSnapshot struct {
	*js.Object
}

type CollectionReference struct {
	*js.Object
	ID     string
	Parent *DocumentReference
}

func (c *CollectionReference) Add(data interface{}) *Promise {
	return &Promise{Object: c.Call("add", data)}
}

type Query struct {
	*js.Object
}

type QuerySnapshot struct {
	*js.Object
}

func (q *Query) EndAt(args ...interface{}) *Query {
	return &Query{Object: q.Call("endAt", args)}
}

func (q *Query) EndBefore(args ...interface{}) *Query {
	return &Query{Object: q.Call("endBefore", args)}
}

func (q *Query) Get(args ...interface{}) *QuerySnapshot {
	return &QuerySnapshot{Object: q.Call("get", args)}
}

func (q *Query) Limit(limit int) *Query {
	return &Query{Object: q.Call("limit", limit)}
}

func (q *Query) OrderBy(fieldPath string, directionStr *string) *Query {
	if directionStr != nil {
		return &Query{Object: q.Call("orderBy", fieldPath, *directionStr)}
	}
	return &Query{Object: q.Call("orderBy", fieldPath)}
}

func (q *Query) StartAfter(args ...interface{}) *Query {
	return &Query{Object: q.Call("startAfter", args)}
}

func (q *Query) StartAt(args ...interface{}) *Query {
	return &Query{Object: q.Call("startAt", args)}
}

func (q *Query) Where(fieldPath string, opStr string, value interface{}) *Query {
	return &Query{Object: q.Call("where", fieldPath, opStr, value)}
}
