package routing

import (
	"log"
	"net/http"
	"time"
)

// HTTPHandle for router
func (router *Router) HTTPHandle(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	req := r.FormValue("req")
	start := time.Now()
	resp, err := route(name, []byte(req))
	if err != nil {
		log.Print(newErrorLog(name, req, err, time.Since(start)))
		http.Error(w, err.Error(), 500)
		return
	}
	log.Print(newInfoLog(name, req, string(resp), time.Since(start)))
	w.Write(resp)
}
