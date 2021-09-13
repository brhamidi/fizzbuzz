package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(writer http.ResponseWriter, resp *http.Request) {
	fmt.Printf("[%s] requested.\n", time.Now().String())
	fmt.Fprintln(writer, "Hello World")
}

func main() {
	fmt.Println("app starting..")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
