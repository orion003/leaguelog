package config

import ()

type Config struct {
	Database Database `json:database`
	Routing Routing `json:routing`
}

type Database struct {
	Url string `json:url`
}

type Routing struct {
    Root string `json:root`   
}
