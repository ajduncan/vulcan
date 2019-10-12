package scuttlebutt

import (
	"net/http"

	"github.com/ajduncan/vulcan/internal/vulcan"
	"github.com/ajduncan/vulcan/pkg/service"
)

func ReportHandler(w http.ResponseWriter, r *http.Request) {
	// todo
	w.WriteHeader(200)
}

func RunScuttlebuttService() {
	address := vulcan.Getenv("BEACON_HOST", "127.0.0.1") + ":" + vulcan.Getenv("BEACON_PORT", "8002")
	vs := service.NewVulcanService("scuttlebutt", address)
	vs.Router.HandleFunc("/report", ReportHandler)
	vs.RunVulcanServer()
}
