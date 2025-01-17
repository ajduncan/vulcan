package ellipsis

import (
	"github.com/ajduncan/vulcan/internal/vulcan"
	"github.com/ajduncan/vulcan/pkg/service"
)

func RunEllipsisService() {
	address := vulcan.Getenv("BEACON_HOST", "127.0.0.1") + ":" + vulcan.Getenv("BEACON_PORT", "8003")
	vs := service.NewVulcanService("ellipsis", address)
	vs.RunVulcanServer()
}
