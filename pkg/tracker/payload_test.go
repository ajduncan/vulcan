package tracker

import (
	"flag"
	"fmt"
	"os"
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

var pl = map[string][]string{
	"id":       []string{"testid"},
	"subject":  []string{"test subject"},
	"state":    []string{"load"},
	"location": []string{"www.test.biz"},
}

func TestNewPayload(t *testing.T) {
	p := NewPayload(pl)
	if p == nil {
		t.Errorf("Error creating payload.")
	}
}

func TestAdd(t *testing.T) {
	p := NewPayload(pl)
	expected := "testvalue"
	p.Add("state", expected)
	v := p.Get("state")

	if v != expected {
		t.Errorf("Get(state) returned unexpected value: got %v want %v", v, expected)
	}
}

func TestAddURLDict(t *testing.T) {
	urldict := map[string][]string{
		"id": []string{"testvalueid"},
	}

	p := NewPayload(urldict)
	expected := "testvalueid"
	v := p.Get("id")

	if v != expected {
		t.Errorf("AddURLDict, Get(id) returned unexpected value: got %v want %v", v, expected)
	}

}

func TestAddDict(t *testing.T) {
	urldict := map[string][]string{
		"testkeyid": []string{"testvalueid"},
	}

	dict := map[string]string{
		"id":      "testvalueid",
		"testfoo": "testbar",
		"abc":     "def",
	}
	p := NewPayload(urldict)
	p.AddDict(dict)

	expected := "testvalueid"
	v := p.Get("id")

	if v != expected {
		t.Errorf("AddURLDict, Get(id) returned unexpected value: got %v want %v", v, expected)
	}

	// test invalid keys
	expected = ""
	v = p.Get("testfoo")

	if v != expected {
		t.Errorf("AddURLDict, Get(testfoo) returned unexpected value: got %v want %v", v, expected)
	}

	// test non-existent keys
	expected = ""
	v = p.Get("nonexistent")

	if v != expected {
		t.Errorf("AddURLDict, Get(nonexistent) returned unexpected value: got %v want %v", v, expected)
	}

}

func TestGet(t *testing.T) {
	p := NewPayload(pl)
	expected := "testvalue"
	p.Add("id", expected)
	v := p.Get("id")

	if v != expected {
		t.Errorf("Get(id) returned unexpected value: got %v want %v", v, expected)
	}

	// test invalid keys
	expected = ""
	p.Add("invalid", "neverexists")
	v = p.Get("invalid")

	if v != expected {
		t.Errorf("Get(invalid) returned unexpected value: got %v want %v", v, expected)
	}

}
