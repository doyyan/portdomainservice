package main

import (
	"os"
	"os/signal"
	"syscall"

	ctxt "github.com/doyyan/portdomainservice/configuration/context"
	"github.com/doyyan/portdomainservice/configuration/logging"
)

func main() {
	// channel to stop after Interrupt or Kill signals
	errChan := make(chan error)
	ctx, cancel := ctxt.NewCancelContext()

	// create a new logger for cross application logging
	logger := logging.NewLogger()

	go func() {
		// channel to listen on Interrupt or Kill signal from OS
		sig := <-NotifySignals()
		cancel()
		logger.Warn(sig)
		// send to errChannel as we have recieved a Kill/Interrupt signal
		errChan <- nil
	}()
	// Recieve and end the main() and thus stop the server!
	<-errChan
}

// NotifySignals returns a channel that will be
// triggered when SIGINT or SIGTERM is received
func NotifySignals() <-chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	return c
}
