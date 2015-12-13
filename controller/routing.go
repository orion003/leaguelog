package controller

import (
	"net/http"
	"time"

	"leaguelog/Godeps/_workspace/src/github.com/gorilla/mux"

	"leaguelog/logging"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

var root string

func NewRouter(c *Controller, rt string) *mux.Router {
	root = rt

	r := mux.NewRouter()
	r.NotFoundHandler = c

	routes := createRoutes(c)

	for _, route := range routes {
		var handler http.HandlerFunc
		handler = logger(route.Handler, route.Name, c.log)

		r.HandleFunc(route.Pattern, handler).
			Name(route.Name).
			Methods(route.Method)
	}

	r.PathPrefix("/app/").Handler(http.StripPrefix("/app/", http.FileServer(http.Dir(root+"app/"))))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir(root+"assets/"))))

	return r
}

func createRoutes(c *Controller) []Route {
	routes := []Route{
		Route{
			"Leagues",
			"GET",
			"/api/leagues",
			c.GetLeagues,
		},
		Route{
			"League",
			"GET",
			"/api/league/{leagueID:[0-9]+}",
			c.GetLeague,
		},
		Route{
			"League.Standings",
			"GET",
			"/api/league/{leagueID:[0-9]+}/standings",
			c.GetLeagueStandings,
		},
		Route{
			"League.Schedule",
			"GET",
			"/api/league/{leagueID:[0-9]+}/schedule",
			c.GetLeagueSchedule,
		},
		Route{
			"User.Register",
			"POST",
			"/api/user/register",
			c.UserRegister,
		},
		Route{
			"User.Login",
			"POST",
			"/api/user/login",
			c.UserLogin,
		},
	}

	return routes
}

func logger(inner http.Handler, name string, log logging.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Info("Request: "+name, "method", r.Method, "uri", r.RequestURI, "time", time.Since(start))
	}
}
