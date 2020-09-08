package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	// case 1: dump to stdout with default level and format.
	log.Println("Hello world!")

	// case 2: dump to stdout with different levels.
	log.SetLevel(log.TraceLevel)
	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	// log.Fatal("Bye.") // Calls os.Exit(1) after logging
	// log.Panic("I'm bailing.") // Calls panic() after logging

	// case 3: dump to stdout with json format.
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"foo": "foo",
			"bar": "bar",
		},
	).Info("Something happened")

	// case 4: dump to file with json format
	file, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	var fileLogger = log.New()
	fileLogger.Out = file
	fileLogger.SetFormatter(&log.JSONFormatter{})

	fileLogger.WithFields(
		log.Fields{
			"foo": "foo",
			"bar": "bar",
		},
	).Info("Something happened again")
}
