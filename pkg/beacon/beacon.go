package beacon

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	// "github.com/go-redis/redis"

	"github.com/ajduncan/vulcan/internal/vulcan"
	"github.com/ajduncan/vulcan/pkg/service"
	"github.com/ajduncan/vulcan/pkg/tracker"
)

// EventEnqueue takes a tracker and adds it to the redis
// queue to be processed.
/*
func EventEnqueue(key string, t *tracker.Tracker) {
	address := vulcan.Getenv("REDIS_HOST", "127.0.0.1") + ":" + vulcan.Getenv("REDIS_PORT", "6379")
	password := vulcan.Getenv("REDIS_PASSWORD", "")
	database := vulcan.Getenv("REDIS_DATABASE", 0)

	rclient := redis.NewClient(&redis.Options{
			Addr:     address,
			Password: password,
			DB:       database,
	})

	err := rclient.Set(key, t.Serialize(), 0).Err()
	if err != nil {
		// sentry ...
		panic(err)
	}
}
*/
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
	id := pl.Get("id")

	if id != "" {
		// event tracker for a site
		t := tracker.NewTracker(pl /* payload */, pl.Get("id") /* id */, pl.Get("subject") /* subject */)
		t.Print()
		//EventEnqueue(id, t)
	}
}

func RunBeaconService() {
	address := vulcan.Getenv("BEACON_HOST", "127.0.0.1") + ":" + vulcan.Getenv("BEACON_PORT", "8000")
	vs := service.NewVulcanService("beacon", address)
	vs.Router.HandleFunc("/api/v1/beacon", BeaconHandler)
	vs.RunVulcanServer()
}
