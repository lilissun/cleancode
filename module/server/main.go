package main

import (
	"flag"
	"log"
	"net/http"

	"bitbucket.org/lilissun/cleancode/module/routing"
)

var (
	flagAddress = flag.String("addr", "0.0.0.0:8080", "listen address")
	router      = routing.NewRouter()
)

func register(router *routing.Router, name string, handle func([]byte) []byte) {
	if err := router.Register(name, handle); err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", router.HTTPHandle)
	if err := http.ListenAndServe(*flagAddress, nil); err != nil {
		log.Printf("listen and serve error=[%s]", err)
	}
}
