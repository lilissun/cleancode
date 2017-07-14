package routing

import (
	"fmt"
	"net/http"
)

// HTTPHandle for router
func (router *Router) HTTPHandle(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "service name is missing", 500)
		return
	}
	body := r.FormValue("body")
	response, err := router.Process(name, body)
	if err != nil {
		http.Error(w, fmt.Sprintf("service error %s", err), 500)
		return
	}
	w.Write([]byte(response))
}
