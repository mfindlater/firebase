package firebase

import (
	"errors"

	"github.com/gopherjs/gopherjs/js"
)

type App struct {
	*js.Object
	Name    string            `js:"name"`
	Options map[string]string `js:"options"`
}

func NewApp(apiKey string, opts ...func(*App)) (*App, error) {

	f, err := getFirebase()
	if err != nil {
		return nil, err
	}

	a := &App{Options: make(map[string]string)}

	for _, option := range opts {
		option(a)
	}

	a.Options["apiKey"] = apiKey

	if a.Name != "" {
		a.Object = f.Call("initializeApp", a.Options, a.Name)
		return a, nil
	}

	a.Object = f.Call("initializeApp", a.Options)
	a.Name = a.Get("name").String()

	return a, nil
}

func (a *App) Delete() *Promise {
	return &Promise{Object: a.Object.Call("delete")}
}

type Options struct {
}

func (o Options) AuthDomain(authDomain string) func(*App) {
	return func(a *App) {
		a.Options["authDomain"] = authDomain
	}
}

func (o Options) DatabaseURL(databaseURL string) func(*App) {
	return func(a *App) {
		a.Options["databaseURL"] = databaseURL
	}
}

func (o Options) StorageBucket(storageBucket string) func(*App) {
	return func(a *App) {
		a.Options["storageBucket"] = storageBucket
	}
}

func (o Options) MessagingSenderID(messagingSenderID string) func(*App) {
	return func(a *App) {
		a.Options["messagingSender"] = messagingSenderID
	}
}

func (o Options) Name(name string) func(*App) {
	return func(a *App) {
		a.Name = name
	}
}

func getFirebase() (*js.Object, error) {
	f := js.Global.Get("firebase")
	if f == js.Undefined {
		return nil, errors.New("could not get firebase instance")
	}

	return f, nil
}
