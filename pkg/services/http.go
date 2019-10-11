// services includes generic handlers, net/http and mux code for instances of
// servers with API endpoints further defined within their respective packages.
package services

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type VulcanService struct {
	Instance string
	Address  string
	Router   *mux.Router
}

// NotFoundHandler provides a default not found handler for the instance.
func (vs *VulcanService) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, `404 page not found`)
}

func (vs *VulcanService) FavicoHandler(w http.ResponseWriter, r *http.Request) {
	// blank favico default handler.
	w.Header().Set("Content-Type", "image/x-icon")
	w.Header().Set("Cache-Control", "public, max-age=7776000")
	io.WriteString(w, "data:image/x-icon;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQEAYAAABPYyMiAAAABmJLR0T///////8JWPfcAAAACXBIWXMAAABIAAAASABGyWs+AAAAF0lEQVRIx2NgGAWjYBSMglEwCkbBSAcACBAAAeaR9cIAAAAASUVORK5CYII=\n")
}

// Home handler provides a default index handler for the instance.
func (vs *VulcanService) HomeHandler(w http.ResponseWriter, r *http.Request) {
	templateLocation := "web/" + vs.Instance + "/page/index.html"
	baseTemplateLocation := "web/" + vs.Instance + "/templates/base.html"
	tmpl, err := template.ParseFiles(templateLocation, baseTemplateLocation)
	if err != nil {
		vs.NotFoundHandler(w, r)
	} else {
		w.WriteHeader(http.StatusOK)
		tmpl.ExecuteTemplate(w, "base", struct{ Data string }{Data: "data"})
	}
}

// HealthCheckHandler provides a default health check response (in JSON) for the
// instance.
func (vs *VulcanService) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, `{"alive": true}`)
}

// Generic handler for /page/<page>.html requests, which reads from the
// root/web/<instance>/templates/<page>.html template.
func (vs *VulcanService) PageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	templateLocation := "web/" + vs.Instance + "/page/" + vars["page"] + ".html"
	baseTemplateLocation := "web/" + vs.Instance + "/templates/base.html"
	tmpl, err := template.ParseFiles(templateLocation, baseTemplateLocation)
	if err != nil {
		vs.NotFoundHandler(w, r)
	} else {
		w.WriteHeader(http.StatusOK)
		tmpl.ExecuteTemplate(w, "base", struct{ Data string }{Data: "data"})
	}
}

// Create a vulcan service with appropriate handlers.
// instance is a key that will be used in loading templates, static files, etc.
// address is the host and port to listen on
func NewVulcanService(instance string, address string) *VulcanService {
	vs := new(VulcanService)
	vs.Instance = instance
	vs.Address = address

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(vs.NotFoundHandler)
	r.HandleFunc("/favicon.ico", vs.FavicoHandler)
	r.HandleFunc("/", vs.HomeHandler)
	r.HandleFunc("/health", vs.HealthCheckHandler)
	r.HandleFunc("/page/{page}.html", vs.PageHandler)

	vs.Router = r
	return vs
}

// Creates a new net/http service with a VulcanService configuration,
// then run the http.Server
func (vs *VulcanService) RunVulcanServer() {
	server := &http.Server{
		Handler:      vs.Router,
		Addr:         vs.Address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Printf("Serving %v for analytics on: %v.\n", vs.Instance, vs.Address)
	log.Fatal(server.ListenAndServe())
}
