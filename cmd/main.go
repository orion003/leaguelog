package main

import (
	"fmt"
	"net/http"
)

func indexHandler(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "../web/angular/index.html")
}

func main() {
	http.HandleFunc("/", indexHandler) // unknown paths should display this root

	fmt.Print("Listening on port 8000")
	http.ListenAndServe(":8000", nil)
}
