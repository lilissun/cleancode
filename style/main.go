package main

import (
	"flag"
	"fmt"
	"strings"
)

var (
	router   = make(map[string]func(string) string)
	flagName = flag.String("name", "Guest", "the user's name")
)

func sayhi(name string) string {
	return fmt.Sprintf("Nice to meet you, %s!", name)
}

func praise(name string) string {
	return fmt.Sprintf("You are one cute honey, %s!", name)
}

func question(name string) string {
	return fmt.Sprintf("Sorry, I did not know you, %s.", name)
}

func init() {
	router["tom"] = sayhi
	router["amy"] = praise
}

func main() {
	flag.Parse()

	if handle, exist := router[strings.ToLower(*flagName)]; exist {
		fmt.Println(handle(*flagName))
		return
	}
	fmt.Println(question(*flagName))
	return
}
