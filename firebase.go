package firebase

import (
	"github.com/gopherjs/gopherjs/js"
)

//Config firebase config.
type Config struct {
	APIKey            string `json:"apiKey"`
	AuthDomain        string `json:"authDomain"`
	DatabaseURL       string `json:"databaseURL"`
	ProjectID         string `json:"projectId"`
	StorageBucket     string `json:"storageBucket"`
	MessagingSenderID string `json:"messagingSenderId"`
}

//Firebase global firebase object.
type Client struct {
	*js.Object
}

//NewFirebase gets firebase global object.
func NewClient() *Client {
	return &Client{Object: js.Global.Get("firebase")}
}

//InitializeApp creates and initializes a Firebase app instance.
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
