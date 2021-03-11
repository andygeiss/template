package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

var (
	name    string = "no-name"
	build   string = "no-build"
	version string = "no-version"
)

func main() {
	log.Printf("%s %s (%s)\n", name, version, build)

	sut := func(w http.ResponseWriter, r *http.Request) {

	}

	http.HandleFunc("/", sut)
	http.ListenAndServe(":3000", nil)
}
