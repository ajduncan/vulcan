package ellipsis

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
		Addr:         vulcan.Getenv("ELLIPSIS_HOST", "127.0.0.1") + ":" + vulcan.Getenv("ELLIPSIS_PORT", "8003"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("... up.\n")
	log.Fatal(srv.ListenAndServe())
}
