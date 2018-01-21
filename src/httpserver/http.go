package main

import (
	"fmt"
	"flag"
	"net/http"
	"os"
	)

func serve() {
	var dir string
	var err error

	port := flag.String("port", "3000", "Listening Port")
	path := flag.String("path", "", "path to server over http")
	flag.Parse()

	if *path == "" {
		dir, err = os.Getwd()
		if err != nil {
			fmt.Println("Permission not available for accessing this directory ")
		}
	} else {
		dir = *path
	}

	http.ListenAndServe(":"+*port, http.FileServer(http.Dir(dir)))
}