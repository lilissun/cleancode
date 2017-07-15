package main

import (
	"flag"
	"log"
	"net/http"

	"bitbucket.org/lilissun/cleancode/module/modulefix/routing"
)

var (
	flagAddress = flag.String("addr", "0.0.0.0:8080", "listen address")
	router      = routing.NewRouter()
)

func register(name string, handle func([]byte) ([]byte, error)) {
	if err := router.Register(name, handle); err != nil {
		log.Panicf("register name=[%s] error=[%s]", name, err)
	}
}

func main() {
	http.HandleFunc("/", router.HTTPHandle)
	if err := http.ListenAndServe(*flagAddress, nil); err != nil {
		log.Printf("listen and serve error=[%s]", err)
	}
}
