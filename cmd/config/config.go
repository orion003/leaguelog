package config

import ()

type Config struct {
	Database Database `json:"database"`
	Routing  Routing  `json:"routing"`
	Auth     Auth     `json:"auth"`
}

type Database struct {
	Url string `json:"url"`
}

type Routing struct {
	Root string `json:"root"`
	Port string `json:"port"`
}

type Auth struct {
	Key string `json:"key"`
}
