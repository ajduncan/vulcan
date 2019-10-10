package katric

import (
	"github.com/ajduncan/vulcan/internal/vulcan"
	"github.com/ajduncan/vulcan/pkg/services"
)


func RunKatricService() {
	address := vulcan.Getenv("BEACON_HOST", "127.0.0.1") + ":" + vulcan.Getenv("BEACON_PORT", "8001")
	vs := services.NewVulcanService("katric", address)
	vs.RunVulcanServer()
}
