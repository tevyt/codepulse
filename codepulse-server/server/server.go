package server

import (
	"fmt"
	"net/http"
	"strconv"

	"tevyt.com/codepulse-server/config"
)

func StartServer(configuration config.ServerConfiguration) error {

	fs := http.FileServer(http.Dir("./static"))

	fmt.Printf("Listening on port %d", configuration.Port)
	return http.ListenAndServe(":"+strconv.Itoa(configuration.Port), fs)
}
