package main

import (
	"io/fs"
	"log"
	"os"
	"os/signal"
	"syscall"
	"task-scheduler/internal/daemon"
)

func main() {
	logFile, err := os.OpenFile("/tmp/taskd.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, fs.ModeSocket)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	server, err := daemon.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	err = server.Start()
	if err != nil {
		log.Fatal(err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	errChan := make(chan error, 1)
	go func() {
		errChan <- server.Start()
	}()

	select {
	case err := <-errChan:
		log.Printf("Server stopped with error: %v", err)
	case sig := <-sigChan:
		log.Printf("Received signal: %v", sig)
		if err := server.Shutdown(); err != nil {
			log.Printf("Error during shutdown: %v", err)
		}
	}
}
