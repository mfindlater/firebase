package firebase

import (
	"errors"

	"github.com/gopherjs/gopherjs/js"
)

const SDKVersion = "4.8.1"

type Options struct {
	APIKey            string
	AuthDomain        string
	DatabaseURL       string
	ProjectID         string
	StorageBucket     string
	MessagingSenderID string
}

type App struct {
	*js.Object
	Name    string   `js:"name"`
	Options *Options `js:"options"`
}

func NewApp(options *Options, opts ...func(*App)) (*App, error) {

	f, err := getFirebase()
	if err != nil {
		return nil, err
	}

	o := map[string]string{
		"apiKey":            options.APIKey,
		"authDomain":        options.AuthDomain,
		"databaseURL":       options.DatabaseURL,
		"projectId":         options.ProjectID,
		"storageBucket":     options.StorageBucket,
		"messagingSenderId": options.MessagingSenderID,
	}

	a := &App{Options: options}

	for _, option := range opts {
		option(a)
	}

	a.Object = f.Call("initializeApp", o)
	//a.Name = a.Get("name").String()

	return a, nil
}

func (a *App) Delete() error {
	return (&Promise{Object: a.Object.Call("delete")}).Convert()
}

func (a *App) Firestore() *Firestore {
	return &Firestore{Object: a.Call("firestore")}
}

func (a *App) Auth() *Auth {
	return &Auth{Object: a.Call("auth")}
}

func Name(name string) func(*App) {
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
