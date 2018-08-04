package handlers

import (
	"github.com/kensay98/vindecoder/app/decoder"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/kensay98/vindecoder/logger"
	"github.com/kensay98/vindecoder/storage"
)

var log = logger.GetLogger()
var db = storage.GetStorage()
var currentDecoder decoder.Decoder

type App struct {
	router *mux.Router
}

func (app *App) Run(addr string) {
	http.Handle("/", app.router)
	log.Fatal(http.ListenAndServe(addr, app.router))
}

func init() {
	currentDecoder = decoder.UsableDecoder
}

func GetApp() (app *App) {
	r := mux.NewRouter()
	r.HandleFunc("/vin/{vin}", DecodeVin).Methods("GET")

	r.Use(loggingMiddleware)
	r.Use(contentTypeMiddleware)

	return &App{
		router: r,
	}
}



