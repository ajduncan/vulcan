package beacon

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
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

func TestBeaconHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/api/v1/beacon", strings.NewReader("foo&bar=baz"))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(BeaconHandler)
	handler.ServeHTTP(w, req)

	// Expect StatusNoContent
	if status := w.Code; status != http.StatusNoContent {
		t.Errorf("BeaconHandler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}
}
