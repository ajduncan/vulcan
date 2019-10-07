package scuttlebutt

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/ajduncan/vulcan/internal/vulcan"
)

func ReportHandler(w http.ResponseWriter, r *http.Request) {
	// todo
	w.WriteHeader(200)
}

func Serve() {
	r := mux.NewRouter()
	r.HandleFunc("/", vulcan.HomeHandler)
	r.HandleFunc("/report", ReportHandler)
	r.HandleFunc("/health", vulcan.HealthCheckHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         vulcan.Getenv("SCUTTLEBUTT_HOST", "127.0.0.1") + ":" + vulcan.Getenv("SCUTTLEBUTT_PORT", "8002"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Scuttlebutt reporting.\n")
	log.Fatal(srv.ListenAndServe())
}
