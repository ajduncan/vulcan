package ellipsis

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

func TestEllipsisService(t *testing.T) {
	address := vulcan.Getenv("ELLIPSIS_HOST", "127.0.0.1") + ":" + vulcan.Getenv("ELLIPSIS_PORT", "8000")
	vs := service.NewVulcanService("ellipsis", address)

	if vs == nil {
		t.Errorf("Ellipsis service initialization failed")
	}
	// todo.
}
