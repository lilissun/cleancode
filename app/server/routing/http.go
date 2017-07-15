package routing

import (
	"fmt"
	"log"
	"net/http"
)

// HTTPHandle for router
func (router *Router) HTTPHandle(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	request := r.FormValue("req")
	response, err := router.Route(name, []byte(request))
	if err != nil {
		log.Printf("router route name=[%s] and request=[%s] error=[%s]", name, request, err)
		http.Error(w, fmt.Sprintf("service error: %s", err), 500)
		return
	}
	log.Printf("router route name=[%s] and request=[%s] with response=[%s]", name, request, response)
	w.Write([]byte(response))
}
