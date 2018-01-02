package firebase

import "github.com/gopherjs/gopherjs/js"

func APIKey(apiKey string) func(*App) {
	return func(a *App) {
		a.options["apiKey"] = apiKey
	}
}

func AuthDomain(authDomain string) func(*App) {
	return func(a *App) {
		a.options["authDomain"] = authDomain
	}
}

func DatabaseURL(databaseURL string) func(*App) {
	return func(a *App) {
		a.options["databaseURL"] = databaseURL
	}
}

func ProjectID(projectID string) func(*App) {
	return func(a *App) {
		a.options["projectId"] = projectID
	}
}

func StorageBucket(storageBucket string) func(*App) {
	return func(a *App) {
		a.options["storageBucket"] = storageBucket
	}
}

func MessagingSenderID(messagingSenderID string) func(*App) {
	return func(a *App) {
		a.options["messagingSender"] = messagingSenderID
	}
}

func Name(name string) func(*App) {
	return func(a *App) {
		a.Name = name
	}
}

type App struct {
	*js.Object
	options map[string]string
	Name    string
}

func (a *App) Delete() *Promise {
	return &Promise{Object: a.Object.Call("delete")}
}
