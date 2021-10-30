package main

import (
	"fmt"
	"github.com/takama/daemon"
	"os"
	"os/signal"
	"syscall"
)

const (
	// name of the service
	name        = "RabbitMQConsumerServiceDemo"
	description = "Galek: Consumer for RabbitMQ demo"
)

// Service has embedded daemon
type Service struct {
	daemon.Daemon
}

// Manage by daemon commands or run the daemon
func (service *Service) Manage() (string, error) {

	usage := "Usage: myservice install | remove | start | stop | status"

	// if received any kind of command, do it
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			return service.Install()
		case "remove":
			return service.Remove()
		case "start":
			return service.Start()
		case "stop":
			return service.Stop()
		case "status":
			return service.Status()
		default:
			return usage, fmt.Errorf("invalid command")
		}
	}

	// Do something, call your goroutines, etc
	err := Consumer("main consumer")
	if err != nil {
		return "", err
	}
	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	// loop work cycle with accept connections or interrupt
	// by system signal
	/*for {
		select {
		case conn := <-listen:
			go handleClient(conn)

		case killSignal := <-interrupt:
			log.Println("Got signal: ", killSignal)
			Shutdown()

			if killSignal == os.Interrupt {
				return "Daemon was interrupted by system signal", nil
			}
			return "Daemon was killed", nil
		}
	}*/

	return "", nil
}
