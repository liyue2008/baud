package blob

import (
	"encoding/json"
	"log"
	"net/http"
)

type Container struct{}

func (s *Container) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case strings.HasPrefix(r.URL.Path, "fetch"):
		s.handleFetch(w, r)
	case strings.HasPrefix(r.URL.Path, "create"):
		s.handleInsert(w, r)
	case strings.HasPrefix(r.URL.Path, "delete"):
		s.handleDelete(w, r)
	case strings.HasPrefix(r.URL.Path, "status"):
		s.handleStatus(w, r)
	case strings.HasPrefix(r.URL.Path, "/debug/pprof"):
		s.handlePprof(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
