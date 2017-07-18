package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	flagAddress = flag.String("addr", "0.0.0.0:3000", "listen address")
	router      = make(map[string]func([]byte) ([]byte, error))
)

func route(name string, req []byte) ([]byte, error) {
	if proc, exist := router[name]; exist {
		return proc(req)
	}
	return nil, fmt.Errorf("service name=[%s] is unregistered", name)
}

func handle(w http.ResponseWriter, r *http.Request) {
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

func main() {
	flag.Parse()
	http.HandleFunc("/api", handle)
	err := http.ListenAndServe(*flagAddress, nil)
	if err != nil {
		log.Printf("http server listen and serve error=[%s]", err)
	}
}
