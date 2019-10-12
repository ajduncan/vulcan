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
	t := tracker.NewTracker(pl /* payload */, pl.Get("id") /* id */, pl.Get("subject") /* subject */)
	t.Print()
}

func RunBeaconService() {
	address := vulcan.Getenv("BEACON_HOST", "127.0.0.1") + ":" + vulcan.Getenv("BEACON_PORT", "8000")
	vs := service.NewVulcanService("beacon", address)
	vs.Router.HandleFunc("/api/v1/beacon", BeaconHandler)
	vs.RunVulcanServer()
}
