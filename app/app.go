package app

import (
	"context"
	"go-mongo/app/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NewApp(ctx context.Context, cfg config.Config) error {
	r := mux.NewRouter()
	app, err := NewClient(ctx, cfg)
	if err != nil {
		return err
	}

	path := "/player"
	r.HandleFunc(path, app.All).Methods("GET")
	r.HandleFunc(path+"/{id}", app.Load).Methods("GET")
	r.HandleFunc(path, app.Insert).Methods("POST")
	r.HandleFunc(path+"/{id}", app.Update).Methods("PUT")
	r.HandleFunc(path+"/{id}", app.Delete).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
	return nil
}
