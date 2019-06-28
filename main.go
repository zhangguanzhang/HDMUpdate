package main

import (

	log "github.com/sirupsen/logrus"
	"hdm/cmd"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	cmd.Execute()

}


