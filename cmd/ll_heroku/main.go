package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"leaguelog/auth/jwt"
	"leaguelog/cmd/config"
	"leaguelog/controller"
	"leaguelog/database/postgres"
	"leaguelog/logging"
)

var conf config.Config

func main() {
	err := initializeConfig()
	if err != nil {
		fmt.Printf("Unable to initialize config: %v\n", err)
		os.Exit(1)
	}

	log := logging.NewLog15()
	c := controller.NewController(log)

	initializeRepo(c)

	fmt.Printf("Initializing root route: %s \n", conf.Routing.Root)
	r := controller.NewRouter(c, conf.Routing.Root)

	fmt.Printf("Listening on port %s\n", conf.Routing.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", conf.Routing.Port), r)
}

func initializeRepo(c *controller.Controller) {
	manager, err := postgres.NewPgManager(conf.Database.Url)
	if err != nil {
		fmt.Printf("Unable to intialize repo: %v\n", err)
		os.Exit(1)
	}

	repo := postgres.NewPgRepository(manager)

	c.SetRepository(repo)
}

func initializeTokenService(c *controller.Controller) {
	hmac := []byte("4bab192f03bd92e7c95ee7a1d754ab0ae453d4bc1df803bf2fdde7e4ff58b895c0af2f0cf0f6b3522f0b53f11548a1240ef07c3bf4e098e4b940f26f0cb7925a")
	ts := jwt.InitializeJwt(hmac)

	c.SetTokenService(ts)
}

func initializeConfig() error {
	port := os.Getenv("PORT")
	if port == "" {
		return errors.New("Unable to determine the port.")
	}

	root := os.Getenv("ROOT_PATH")
	if root == "" {
		return errors.New("Unable to determine the root.")
	}

	db := os.Getenv("DATABASE_URL")
	if db == "" {
		return errors.New("Unable to determine the database.")
	}

	conf = config.Config{
		Database: config.Database{Url: db},
		Routing:  config.Routing{Root: root, Port: port},
	}

	return nil
}
