package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/takama/daemon"
)

const (

	// name of the service
	name        = "AltarixDemo"
	description = "Altarix demo the work of daemon"

	// port which daemon should be listen
	port = ":9977"
)

// dependencies that are NOT required by the service, but might be used
//var dependencies = []string{"dummy.service"}

var stdlog, errlog *log.Logger
var fErrors, fEvents *os.File

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
			log.Fatalf("invalid command")

			return usage, nil
		}
	}

	// Do something, call your goroutines, etc
	go LogicImpl()

	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	// loop work cycle with accept connections or interrupt
	// by system signal
	for {
		select {
		case killSignal := <-interrupt:

			if ISDebug {
				log.Println("Got signal: ", killSignal)
			} else {
				stdlog.Println("Got signal: ", killSignal)
			}

			Shutdown()

			if killSignal == os.Interrupt {
				return "Daemon was interrupted by system signal", nil
			}
			return "Daemon was killed", nil
		}
	}
}

func InitLogging() {
	if ISDebug {
		log.Println("[InitLogging] Started")
	}

	fErrors, err = os.OpenFile("error.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fErrors.Close()
		log.Fatal("[InitLogging] Failed opening error.log with error: ", err)
		os.Exit(1)
	}

	fEvents, err = os.OpenFile("events.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fEvents.Close()
		log.Fatal("[InitLogging] Failed opening error.log with error: ", err)
		os.Exit(1)
	}

	stdlog = log.New(fEvents, "[Info]", 0)
	errlog = log.New(fErrors, "[Error]", 0)

	log.SetOutput(fEvents)

	if ISDebug {
		log.Println("[InitLogging] Finished")
	}
}

func CloseLoggingFiles() {
	defer fErrors.Close()
	defer fEvents.Close()
}

func DeamonMain() {
	InitLogging()

	if ISDebug {
		log.Println("[DeamonMain] Started")
	}

	var srv daemon.Daemon
	srv, err = daemon.New(name, description /*, dependencies...*/)
	printError(file_line())

	service := &Service{srv}

	var status string = ""
	status, err = service.Manage()

	printError(file_line())

	log.Println(status)

	if ISDebug {
		log.Println("[DeamonMain] Finished")
	}

	CloseLoggingFiles()
}
