package katric

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/ajduncan/vulcan/internal/vulcan"
)

func Serve() {
	r := mux.NewRouter()
	r.HandleFunc("/", vulcan.HomeHandler)
	r.HandleFunc("/health", vulcan.HealthCheckHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         vulcan.Getenv("KATRIC_HOST", "127.0.0.1") + ":" + vulcan.Getenv("KATRIC_PORT", "8001"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Katric interface up.\n")
	log.Fatal(srv.ListenAndServe())
}
