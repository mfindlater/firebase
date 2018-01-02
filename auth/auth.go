package auth

import "github.com/gopherjs/gopherjs/js"


type Client struct {
	*js.Object
}

func NewClient() *Client {
	return &Client{Object: js.Global.Get("firebase").Call("auth")}
}
