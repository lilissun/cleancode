package routing

import (
	"fmt"
	"log"
	"net/http"
)

// HTTPHandle for router
func (router *Router) HTTPHandle(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	req := r.FormValue("req")
	response, err := router.Route(name, []byte(req))
	if err != nil {
		log.Printf("router route name=[%s] and request=[%s] error=[%s]", name, req, err)
		http.Error(w, fmt.Sprintf("service error %s", err), 500)
		return
	}
	w.Write([]byte(response))
}
