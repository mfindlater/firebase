package firestore

import "github.com/gopherjs/gopherjs/js"

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
