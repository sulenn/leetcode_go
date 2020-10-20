package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var logger = log.New()

func init() {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stderr)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)

	//log.se
}

func main() {
	//logger.Out = os.Stdout
	//
	//_, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//if err == nil {
	//	logger.Out = os.Stdout
	//} else {
	//	logger.Info("Failed to log to file, using default stderr")
	//}

	logger.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	logger.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	logger.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")

	// A common pattern is to re-use fields between logging statements by re-using
	// the logrus.Entry returned from WithFields()
	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
}
