package firebase

import "github.com/gopherjs/gopherjs/js"

type Promise struct {
	*js.Object
}

func (p *Promise) Then(args ...func()) *Promise {
	if len(args) == 1 {
		onResolve := args[0]
		p.Object.Call("then", onResolve)
		return p
	}

	if len(args) == 2 {
		onResolve := args[0]
		onReject := args[1]
		p.Object.Call("then", onResolve, onReject)
		return p
	}

	p.Object.Call("then")
	return p
}

func (p *Promise) Catch(catch func(*js.Error)) {
	p.Object.Call("catch", catch)
}
