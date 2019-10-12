package beacon

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/ajduncan/vulcan/internal/vulcan"
	"github.com/ajduncan/vulcan/pkg/service"
	"github.com/ajduncan/vulcan/pkg/tracker"
)

func BeaconHandler(w http.ResponseWriter, r *http.Request) {
	// Send a response first
	w.WriteHeader(http.StatusNoContent)

	// get the response body
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// sentry
		fmt.Fprintf(w, "Error reading request body: %v", err)
	}

	m, _ := url.ParseQuery(string(b))

	pl := tracker.NewPayload(m /* urldict */)

	// event tracker for a site
	t := tracker.NewTracker(pl /* payload */, "test" /* id */, "testsite" /* subject */)

	fmt.Printf("POST request body: %s\n", b)
	fmt.Printf("Tracker: %v\n", t)
	fmt.Printf("Payload: %v\n", t.Payload)

	fmt.Printf("Payload state: %v\n", t.Payload.Get("state"))
	fmt.Printf("Tracker version: %v\n", t.Version)
	fmt.Printf("Tracker timestamp: %v\n", t.Timestamp)
	fmt.Printf("Tracker id: %v\n", t.Id)
	fmt.Printf("Tracker subject: %v\n", t.Subject)
	fmt.Printf("Payload state: %v\n", t.Payload.Get("state"))

}

func RunBeaconService() {
	address := vulcan.Getenv("BEACON_HOST", "127.0.0.1") + ":" + vulcan.Getenv("BEACON_PORT", "8000")
	vs := service.NewVulcanService("beacon", address)
	vs.Router.HandleFunc("/api/v1/beacon", BeaconHandler)
	vs.RunVulcanServer()
}
