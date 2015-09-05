package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"recleague/controller"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

func NewRouter(c *controller.Controller) *mux.Router {
	r := mux.NewRouter()
	routes := createRoutes(c)

	for _, route := range routes {
		r.HandleFunc(route.Pattern, route.Handler).
			Name(route.Name).
			Methods(route.Method)
	}

	return r
}

func createRoutes(c *controller.Controller) []Route {
    routes := []Route{
    	Route{
    		"Index",
    		"GET",
    		"/",
    		indexHandler,
    	},
    	Route{
    		"AddEmail",
    		"POST",
    		"/users",
    		c.AddEmail,
    	},
    }
    
    return routes
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, root+"index.html")
}
