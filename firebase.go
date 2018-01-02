package firebase

import (
	"github.com/gopherjs/gopherjs/js"
)

type Client struct {
	*js.Object
}

func NewClient() *Client {
	return &Client{Object: js.Global.Get("firebase")}
}

func (c *Client) NewApp(options ...func(*App)) *App {

	a := &App{}

	a.options = make(map[string]string)

	for _, option := range options {
		option(a)
	}

	if a.Name != "" {
		a.Object = c.Call("initializeApp", a.options, a.Name)
		return a
	}

	a.Object = c.Call("initializeApp", a.options)

	return a
}
