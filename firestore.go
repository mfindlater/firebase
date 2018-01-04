package firebase

import (
	"github.com/gopherjs/gopherjs/js"
)

type Firestore struct {
	*js.Object
	App *App `js:"app"`
}

func (f *Firestore) Batch() *WriteBatch {
	return &WriteBatch{Object: f.Call("batch")}
}

func (f *Firestore) Collection(collectionPath string) *CollectionReference {
	c := &CollectionReference{}
	c.FirestoreQuery = &FirestoreQuery{Object: f.Call("collection", collectionPath)}
	return c
}

func (f *Firestore) Doc(documentPath string) *DocumentReference {
	return &DocumentReference{Object: f.Call("doc", documentPath)}
}

func (f *Firestore) EnablePersistence() error {
	return (&Promise{Object: f.Call("enablePersistence")}).Convert()
}

func (f *Firestore) RunTransaction(updateFunc func(*Transaction)) error {
	return (&Promise{Object: f.Call("runTransaction", updateFunc)}).Convert()
}

func (f *Firestore) SetLogLevel(logLevel string) {
	f.Call("setLogLevel", logLevel)
}

func (f *Firestore) Settings(settings map[string]string) {
	f.Call("settings", settings)
}

type FirestoreQuery struct {
	*js.Object
	Firestore *Firestore `js:"firestore"`
}

func (q *FirestoreQuery) EndAt(args ...interface{}) *FirestoreQuery {
	return &FirestoreQuery{Object: q.Call("endAt", args)}
}

func (q *FirestoreQuery) EndBefore(args ...interface{}) *FirestoreQuery {
	return &FirestoreQuery{Object: q.Call("endBefore", args)}
}

func (q *FirestoreQuery) Get(args ...interface{}) *FirestoreQuerySnapshot {
	return &FirestoreQuerySnapshot{Object: q.Call("get", args)}
}

func (q *FirestoreQuery) Limit(limit int) *FirestoreQuery {
	return &FirestoreQuery{Object: q.Call("limit", limit)}
}

func (q *FirestoreQuery) OrderBy(fieldPath string, directionStr *string) *FirestoreQuery {
	if directionStr != nil {
		return &FirestoreQuery{Object: q.Call("orderBy", fieldPath, *directionStr)}
	}
	return &FirestoreQuery{Object: q.Call("orderBy", fieldPath)}
}

func (q *FirestoreQuery) StartAfter(args ...interface{}) *FirestoreQuery {
	return &FirestoreQuery{Object: q.Call("startAfter", args)}
}

func (q *FirestoreQuery) StartAt(args ...interface{}) *FirestoreQuery {
	return &FirestoreQuery{Object: q.Call("startAt", args)}
}

func (q *FirestoreQuery) Where(fieldPath string, opStr string, value interface{}) *FirestoreQuery {
	return &FirestoreQuery{Object: q.Call("where", fieldPath, opStr, value)}
}

type FirestoreQueryMetadata struct {
	*js.Object
	FromCache        bool `js:"fromCache"`
	HasPendingWrites bool `js:"hasPendingWrites"`
}

type FirestoreQuerySnapshot struct {
	*js.Object
	DocChanges []*DocumentChange       `js:"docChanges"`
	Docs       []*DocumentSnapshot     `js:"docs"`
	Empty      bool                    `js:"empty"`
	Metadata   *FirestoreQueryMetadata `js:"metadata"`
	Query      *FirestoreQuery         `js:"query"`
	Size       int                     `js:"size"`
}

func (q *FirestoreQuerySnapshot) ForEach(callback func(*DocumentSnapshot), thisArg *js.Object) {
	q.Call("forEach", callback, thisArg)
}

type CollectionReference struct {
	*FirestoreQuery
	ID     string             `js:"id"`
	Parent *DocumentReference `js:"parent"`
}

func (c *CollectionReference) Add(data interface{}) (*DocumentReference, error) {
	o, err := (&Promise{Object: c.Call("add", data)}).ConvertWithResult()
	return &DocumentReference{Object: o}, err
}

type Transaction struct {
	*js.Object
}

func (t *Transaction) Delete() *Transaction {
	return &Transaction{Object: t.Call("delete")}
}

func (t *Transaction) Get(documentRef *DocumentReference) (*DocumentSnapshot, error) {
	o, err := (&Promise{Object: t.Call("get", documentRef)}).ConvertWithResult()
	return &DocumentSnapshot{Object: o}, err
}

func (t *Transaction) Set(documentRef *DocumentReference, data interface{}, options interface{}) *Transaction {
	return &Transaction{Object: t.Call("set", documentRef, data, options)}
}

func (t *Transaction) Update(documentRef *DocumentReference, args ...interface{}) *Transaction {
	return &Transaction{Object: t.Call("update", documentRef, args)}
}

type WriteBatch struct {
	*js.Object
}

func (w *WriteBatch) Commit() error {
	return (&Promise{Object: w.Call("commit")}).Convert()
}

func (w *WriteBatch) Delete() *WriteBatch {
	return &WriteBatch{Object: w.Call("delete")}
}

func (w *WriteBatch) Set(documentRef *DocumentReference) *WriteBatch {
	return &WriteBatch{Object: w.Call("set", documentRef)}
}

func (w *WriteBatch) Update(documentRef *DocumentReference, args ...interface{}) *WriteBatch {
	return &WriteBatch{Object: w.Call("update", documentRef, args)}
}

type DocumentReference struct {
	*js.Object
	Firestore *Firestore           `js:"firestore"`
	ID        string               `js:"id"`
	Parent    *CollectionReference `js:"parent"`
}

func (d *DocumentReference) Collection(collectionPath string) *CollectionReference {
	c := &CollectionReference{}
	c.FirestoreQuery = &FirestoreQuery{Object: d.Call("collection", collectionPath)}
	return c
}

func (d *DocumentReference) Delete() *Promise {
	return &Promise{Object: d.Call("delete")}
}

func (d *DocumentReference) Get() *Promise {
	return &Promise{Object: d.Call("get")}
}

func (d *DocumentReference) OnSnapshot(args ...interface{}) {
	d.Call("onSnapshot", args)
}

func (d *DocumentReference) Set(data interface{}, options interface{}) *Promise {
	return &Promise{Object: d.Call("set", data, options)}
}

func (d *DocumentReference) Update(args ...interface{}) *Promise {
	return &Promise{Object: d.Call("update")}
}

type FirestoreSnapshotMetadata struct {
	*js.Object
	FromCache        bool `js:"fromCache"`
	HasPendingWrites bool `js:"hasPendingWrites"`
}

type DocumentSnapshot struct {
	*js.Object
	Exists   string                    `js:"exists"`
	ID       string                    `js:"id"`
	Metadata FirestoreSnapshotMetadata `js:"metadata"`
	Ref      *DocumentReference        `js:"ref"`
}

func (d *DocumentSnapshot) Data() *js.Object {
	return d.Call("data")
}

func (d *DocumentSnapshot) Get(fieldPath string) *js.Object {
	return d.Call("get", fieldPath)
}

type GeoPoint struct {
	*js.Object
	Latitude  string `js:"latitude"`
	Longitude string `js:"longitude"`
}

func NewGeoPoint(latitude, longitude int) *GeoPoint {
	return &GeoPoint{Object: js.Global.Get("GeoPoint").New(latitude, longitude)}
}

type DocumentChange struct {
	*js.Object
}
