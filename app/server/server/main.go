package main

import (
	"flag"
	"log"
	"net/http"

	"bitbucket.org/lilissun/cleancode/app/server/routing"
)

var (
	flagAddress = flag.String("addr", "0.0.0.0:3000", "listen address")
	router      = routing.NewRouter()
)

func register(name string, handle func([]byte) ([]byte, error)) {
	if err := router.Register(name, handle); err != nil {
		log.Panicf("register name=[%s] error=[%s]", name, err)
	}
}

func main() {
	http.HandleFunc("/api", router.HTTPHandle)
	if err := http.ListenAndServe(*flagAddress, nil); err != nil {
		log.Printf("listen and serve error=[%s]", err)
	}
}
