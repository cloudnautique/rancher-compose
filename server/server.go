package server

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

func StartServer() {

	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/v1/projects", ProjectHandler).Methods("POST")
	http.Handle("/", gorillaRoute)

	log.Infof("Server listening on port 8081")
	http.ListenAndServe(":8081", nil)
}
