package beacon

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ajduncan/vulcan/internal/vulcan"
	"github.com/ajduncan/vulcan/pkg/services"
)

func BeaconHandler(w http.ResponseWriter, r *http.Request) {
	// Send a response first and log the data
	w.WriteHeader(http.StatusNoContent)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error reading request body: %v", err)
	}

	fmt.Printf("POST request body: %s\n", b)
}

func RunBeaconService() {
	address := vulcan.Getenv("BEACON_HOST", "127.0.0.1") + ":" + vulcan.Getenv("BEACON_PORT", "8000")
	vs := services.NewVulcanService("beacon", address)
	vs.Router.HandleFunc("/api/v1/beacon", BeaconHandler)
	vs.RunVulcanServer()
}
