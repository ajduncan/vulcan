package scuttlebutt

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/ajduncan/vulcan/internal/vulcan"
	"github.com/ajduncan/vulcan/pkg/service"
)

var (
	cwd_arg = flag.String("cwd", "", "set cwd")
)

func init() {
	flag.Parse()
	if *cwd_arg != "" {
		if err := os.Chdir(*cwd_arg); err != nil {
			fmt.Println("Chdir error:", err)
		}
	}
}

func TestScuttlebuttService(t *testing.T) {
	address := vulcan.Getenv("SCUTTLEBUTT_HOST", "127.0.0.1") + ":" + vulcan.Getenv("SCUTTLEBUTT_PORT", "8002")
	vs := service.NewVulcanService("scuttlebutt", address)

	if vs == nil {
		t.Errorf("Scuttlebutt service initialization failed")
	}
	// todo.
}
