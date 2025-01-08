package api

import "go_pass/config"

type Api struct {
	Config config.Config
	Url    string
}

func NewApi(conf config.Config, url string) *Api {
	return &Api{
		Config: conf,
		Url:    url,
	}
}
