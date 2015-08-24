package main

import (
	"fmt"
	"net/http"
	
	"github.com/gorilla/mux"
)

const root string = "../web/angular/"

func main() {
    r := mux.NewRouter()
    
	r.HandleFunc("/", indexHandler)
	r.PathPrefix("/app/").Handler(http.StripPrefix("/app/", http.FileServer(http.Dir(root + "app/"))))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir(root + "assets/")))) 

	fmt.Print("Listening on port 8000")
	http.ListenAndServe(":8000", r)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, root + "index.html")
}