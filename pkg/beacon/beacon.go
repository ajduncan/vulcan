package beacon

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/ajduncan/vulcan/internal/vulcan"
)

func BeaconHandler(w http.ResponseWriter, r *http.Request) {
	// Send a response first and log the data
	w.WriteHeader(204)

	vars := mux.Vars(r)
	// q := r.URL.Query()
	fmt.Fprintf(w, "Image guid: %v\n", vars["guid"])
}

func Serve() {
	r := mux.NewRouter()
	r.HandleFunc("/", vulcan.HomeHandler)
	r.HandleFunc("/health", vulcan.HealthCheckHandler)

	// Should the image be unique to the site?
	r.HandleFunc("/beacon/{guid}/0.gif", BeaconHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         vulcan.Getenv("BEACON_HOST", "127.0.0.1") + ":" + vulcan.Getenv("BEACON_PORT", "8000"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Serving beacon for analytics.\n")
	log.Fatal(srv.ListenAndServe())
}
