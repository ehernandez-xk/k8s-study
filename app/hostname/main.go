package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	port = "8080"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		fmt.Fprintln(w, "Unable to get the hostname")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	hostname := fmt.Sprintf("Hostname: %v", name)
	fmt.Println(hostname)
	fmt.Fprintln(w, hostname)
	w.WriteHeader(http.StatusOK)
}

func main() {
	fmt.Println("Starting server")
	http.HandleFunc("/", rootHandler)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
