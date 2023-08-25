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
	var headers string
	// Loop over header names
	headers += "\n"
	for name, values := range r.Header {
		// Loop over all values for the name.
		for _, value := range values {
			headers += "\t" + name + ":" + value + "\n"
		}
	}
	fmt.Fprintf(w, "Env: %s \nHost: %s\nPath: %s\nHeaders:%s", env, r.Host, r.URL.Path, headers)

}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
