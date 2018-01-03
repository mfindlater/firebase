package firebase

import (
	"errors"

	"github.com/gopherjs/gopherjs/js"
)

const SDKVersion = "4.8.1"

type Options struct {
	*js.Object
	APIKey            string `js:"apiKey"`
	AuthDomain        string `js:"authDomain"`
	DatabaseURL       string `js:"databaseURL"`
	StorageBucket     string `js:"storageBucket"`
	MessagingSenderID string `js:"messagingSenderId"`
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

	obj := js.Global.Get("Object").New()

	if options.APIKey != "" {
		obj.Set("apiKey", options.APIKey)
	}

	if options.AuthDomain != "" {
		obj.Set("authDomain", options.AuthDomain)
	}

	if options.DatabaseURL != "" {
		obj.Set("databaseURL", options.DatabaseURL)
	}

	if options.StorageBucket != "" {
		obj.Set("storageBucket", options.StorageBucket)
	}

	if options.MessagingSenderID != "" {
		obj.Set("messagingSenderId", options.MessagingSenderID)
	}

	a := &App{Options: &Options{Object: obj}}

	for _, option := range opts {
		option(a)
	}

	if a.Name != "" {
		a.Object = f.Call("initializeApp", a.Options, a.Name)
		return a, nil
	}

	a.Object = f.Call("initializeApp", a.Options)
	a.Name = a.Get("name").String()

	return a, nil
}

func (a *App) Delete() error {

	p := &Promise{Object: a.Object.Call("delete")}

	r := func() error {
		hasError := false
		errorToReturn := &Error{}
		p.Then(func() {
			hasError = false
		}).Catch(func(e *Error) {
			errorToReturn = e
		})

		if hasError {
			return errorToReturn
		}

		return nil
	}()

	return r
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
