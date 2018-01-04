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

func (p *Promise) Convert() error {
	c := make(chan *Error)

	go func() {
		p.Then(func() {
			c <- nil
		}).Catch(func(e *Error) {
			c <- e
		})
	}()

	return <-c
}

type result struct {
	value *js.Object
	err   *Error
}

func (p *Promise) ConvertWithResult() (*js.Object, error) {
	c := make(chan result)

	go func() {
		p.Then(func(o *js.Object) {
			c <- result{value: o, err: nil}
		}).Catch(func(e *Error) {
			c <- result{value: nil, err: e}
		})
	}()

	out := <-c
	return out.value, out.err
}
