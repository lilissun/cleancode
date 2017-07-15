package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	flagAddress = flag.String("addr", "0.0.0.0:80", "listen address")
)

func route(name string, req []byte) ([]byte, error) {
	if name == RGB2HSLName {
		return rgb2hsl(req)
	}
	if name == HSL2RGBName {
		return hsl2rgb(req)
	}
	return nil, fmt.Errorf("service name=[%s] is unregistered", name)
}

func handle(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	req := r.FormValue("req")
	resp, err := route(name, []byte(req))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(resp)
}

func main() {
	flag.Parse()
	http.HandleFunc("/", handle)
	http.ListenAndServe(*flagAddress, nil)
}
