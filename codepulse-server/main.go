package main

import (
	"tevyt.com/codepulse-server/config"
	"tevyt.com/codepulse-server/server"
)

func main() {
	conf, err := config.LoadConfiguration()
	if err != nil {
		panic(err)
	}

	server.StartServer(conf.Server)

}
