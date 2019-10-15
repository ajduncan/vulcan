package katric

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

func TestKatricService(t *testing.T) {
	address := vulcan.Getenv("KATRIC_HOST", "127.0.0.1") + ":" + vulcan.Getenv("KATRIC_PORT", "8001")
	vs := service.NewVulcanService("katric", address)

	if vs == nil {
		t.Errorf("Katric service initialization failed")
	}
	// todo.
}
