package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	flagAddress = flag.String("addr", "0.0.0.0:8080", "listen address")
	router      = make(map[string]func([]byte) ([]byte, error))
)

func init() {
	router["rgb2hsl"] = rgb2hsl
	router["hsl2rgb"] = hsl2rgb
}

func route(name string, req []byte) ([]byte, error) {
	if proc, exist := router[name]; exist {
		return proc(req)
	}
	return nil, fmt.Errorf("service name=[%s] is unregistered", name)
}

func handle(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	req := r.FormValue("req")
	resp, err := route(name, []byte(req))
	if err != nil {
		log.Printf("route name=[%s] and request=[%s] error=[%s]", name, req, err)
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(resp)
}

func main() {
	flag.Parse()
	http.HandleFunc("/", handle)
	err := http.ListenAndServe(*flagAddress, nil)
	if err != nil {
		log.Printf("http server listen and serve error=[%s]", err)
	}
}
