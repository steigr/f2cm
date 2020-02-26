package cmd

import (
	"log"
	"os"
	"os/signal"
)

func checkWorkloadMapping() {
	if len(configMapNames) != len(directoryNames) {
		log.Panicln("configmaps does not match directories")
	}
}

func waitForTermination() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	done := make(chan bool, 1)

	go func() {
		_ = <-signals
		done <- true
	}()

	<-done
}
