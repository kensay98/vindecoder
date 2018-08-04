package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func DecodeVin(w http.ResponseWriter, r *http.Request) {
	vin := mux.Vars(r)["vin"]

	// Search vin in our db
	vinFromDB := db.Get(vin) // TODO: Add error handling
	if vinFromDB != "" {
		log.Debug("Reading from cache")
		w.Write([]byte(vinFromDB))
		return
	}

	// Send request if vin couldn't find in our db
	decodedVin, err := currentDecoder.Decode(vin)

	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	data, _ := json.Marshal(decodedVin)
	db.Set(vin, string(data))

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
