package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	env := os.Getenv("env")
	if env == "" {
		env = "unknown"
	}
	var headers, headersFlat string
	// Loop over header names
	headers += "\n"
	headersFlat += "["
	for name, values := range r.Header {
		// Loop over all values for the name.
		for _, value := range values {
			headers += "\t" + name + ":" + value + "\n"
			headersFlat += name + ":" + value + ","
		}
	}
	headersFlat += "]"
	log.Println("Env: ", env, "RemoteAddr", r.RemoteAddr, "Host: ", r.Host, "Path: ", r.URL.Path, "Headers: ", headersFlat)
	fmt.Fprintf(w, "Env: %s \nHost: %s\nPath: %s\nHeaders:%s", env, r.Host, r.URL.Path, headers)

}

func main() {
	http.HandleFunc("/", handler)
	log.SetOutput(os.Stdout)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
