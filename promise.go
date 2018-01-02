package firebase

import "github.com/gopherjs/gopherjs/js"

type Promise struct {
	*js.Object
}

func (p *Promise) All(values ...interface{}) *Promise {
	return &Promise{p.Call("all", values)}
}

func (p *Promise) Reject(err interface{}) *Promise {
	return &Promise{p.Call("reject", err)}
}

func (p *Promise) Resolve(value interface{}) *Promise {
	return &Promise{p.Call("resolve", value)}
}

func (p *Promise) Catch(onReject ...interface{}) *Promise {
	return &Promise{p.Call("catch", onReject)}
}

func (p *Promise) Then(args ...interface{}) *Promise {
	if len(args) == 1 {
		onResolve := args[0]
		return &Promise{p.Call("then", onResolve)}
	}

	if len(args) == 2 {
		onResolve := args[0]
		onReject := args[1]
		return &Promise{p.Call("then", onResolve, onReject)}
	}

	return &Promise{p.Call("then")}
}
