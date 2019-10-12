package katric

import (
	"github.com/ajduncan/vulcan/internal/vulcan"
	"github.com/ajduncan/vulcan/pkg/service"
)

func RunKatricService() {
	address := vulcan.Getenv("BEACON_HOST", "127.0.0.1") + ":" + vulcan.Getenv("BEACON_PORT", "8001")
	vs := service.NewVulcanService("katric", address)
	vs.RunVulcanServer()
}
