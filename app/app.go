package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kisinga/ATS/app/storage"
)

type App struct {
	d        *storage.Database
	handlers map[string]http.HandlerFunc
}

func NewApp(d *storage.Database, prod bool) App {
	app := App{
		d:        d,
		handlers: make(map[string]http.HandlerFunc),
	}
	searchHandler := app.Api
	if !prod {
		searchHandler = disableCors(searchHandler)
	}
	app.handlers["/api"] = searchHandler
	app.handlers["/home"] = http.FileServer(http.Dir("../frontend/dist")).ServeHTTP
	return app
}

func (a *App) Serve(port string) error {
	for path, handler := range a.handlers {
		http.Handle(path, handler)
	}
	log.Println("Web server is available on port 8080")
	return http.ListenAndServe(port, nil)
}

func (a *App) Api(w http.ResponseWriter, r *http.Request) {

}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}

// Needed in order to disable CORS for local development
func disableCors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		h(w, r)
	}
}
