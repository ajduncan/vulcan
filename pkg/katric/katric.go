package katric

import (
	// "fmt"
	// "io/ioutil"
	// "net/http"

	"github.com/ajduncan/vulcan/internal/vulcan"
	"github.com/ajduncan/vulcan/pkg/service"
)

/*
func LogHandler(w http.ResponseWriter, r *http.Request) {
	// Send a response first
	w.WriteHeader(http.StatusNoContent)

	// get the response body
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// sentry
		fmt.Fprintf(w, "Error reading request body: %v", err)
	}

	// parse the log, combine as needed.
}

func TrackerHandler(w http.ResponseWriter, r *http.Request) {
	// Send a response first
	w.WriteHeader(http.StatusNoContent)

	// get the response body
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// sentry
		fmt.Fprintf(w, "Error reading request body: %v", err)
	}

	// transform the payload

	// submit to scylladb

}
*/
func RunKatricService() {
	address := vulcan.Getenv("KATRIC_HOST", "127.0.0.1") + ":" + vulcan.Getenv("KATRIC_PORT", "8001")
	vs := service.NewVulcanService("katric", address)
	// vs.Router.HandleFunc("/api/v1/katric/tracker", TrackerHandler)
	// vs.Router.HandleFunc("/api/v1/katric/log", LogHandler)
	vs.RunVulcanServer()
}
