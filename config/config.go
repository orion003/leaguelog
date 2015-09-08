package config

import ()

type Config struct {
	Database Database `json:database`
}

type Database struct {
	Url string `json:url`
}
