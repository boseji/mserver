// Copyright 2018 @boseji <salearj@hotmail.com> All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package mserver is a wrapper around `net/http` HTTP Server for Graceful shutdown of web server upon
//
// - SIGINT SIGKILL signals sent to the application
//
// - Internal Errors of the web server
//
package mserver

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Mserver or Managed Server is a derivation of the standard HTTP server
// Provided in Golang standard package `net/http`. This type provides the
// facility to stop the HTTP server gracefully and automatically on 3 Events:
// 1. Ctrl + C is pressed or SIGINT is send to program
// 2. SIGKILL is send to program
// 3. Error occurs in Normal Server operation
type Mserver struct {
	Server          *http.Server   // Instance of the Server
	stop            chan os.Signal // Signal Receiver for SIGINT and SIGKILL
	started         bool           // Indicates if the Server was started or not (default = false)
	ShutdownTimeout time.Duration  // Timeout before a force shutdown is called
	Error           error          // Last error in operations of the Server
}

// Server not started Error code
var ErrServerNotStarted = errors.New("Server Not Started Yet")

// Internal function to Stop the Server Gracefully using the `context` package
// This function does not execute if the server is not started
// This function automatically closes the `Mserver.stop` channel to prevent
// multiple calls to the Shutdown method.
// Also this function sets the `Mserver.started` to false to prevent entry
// into this function again.
func (p *Mserver) stopServerInternal() error {

	if !p.started {
		return ErrServerNotStarted
	}

	// Create the Required Shutdown timeout Context
	ctx, cancel := context.WithTimeout(context.Background(), p.ShutdownTimeout)
	defer cancel()

	// Close the Channel in case we missed
	defer close(p.stop)

	// Request the Server Shutdown
	p.started = false
	if p.Error = p.Server.Shutdown(ctx); p.Error != nil {
		return p.Error
	}

	return p.Error
}

// Internal function to run as a Goroutine initiating the Server start using
// `ListenAndServe`
// This function is designed such that it can't be called twice for the same
// `Mserver`.
// Additionally upon occurance of any error during the operation of the server
// this function automatically calls the `stopServerInternal` function
// to shutdown the server.
func (p *Mserver) startGoServerInternal() {

	// Do not allow 2 calls to this function
	if p.started {
		log.Println(" Server Already Stared")
		return
	}
	log.Println(" Server Started ...")
	p.started = true
	p.Error = p.Server.ListenAndServe()
	log.Println(" Server Stopping ...")
	if p.Error != nil {
		p.stopServerInternal()
	}
}

// StartDefaultServer creates a `http.DefaultServeMux` adds it to `http.Server`
// along with the provided `addr` which is the server address.
// The `timeout` is used as wait time before web server is terminated or
// stopped. It is bassically to force close the server in case it does not
// respond to a shutdown request. This feature uses the `context` package.
func (p *Mserver) StartDefaultServer(addr string, timeout time.Duration) {

	p.StartServer(addr, http.DefaultServeMux, timeout)

}

// StartServer creates a server(`http.Server`) using the provided `http.ServeMux` setting
// it up with the provided `addr` as the server address.
// The `mux` can be any type implementing `http.ServeMux` and also the `http.DefaultServeMux`.
// The `timeout` is used as wait time before web server is terminated or
// stopped. It is bassically to force close the server in case it does not
// respond to a shutdown request. This feature uses the `context` package.
func (p *Mserver) StartServer(addr string, mux *http.ServeMux, timeout time.Duration) {

	// Parameter Errors
	if len(addr) == 0 || timeout == (0*time.Second) || mux == nil {
		return
	}

	// Create the Interrupt Source
	p.stop = make(chan os.Signal)
	signal.Notify(p.stop, os.Kill, os.Interrupt)

	// Assign the Wait Timeout during Shutdown
	p.ShutdownTimeout = timeout

	// Create a Callable Server for Later
	p.Server = &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	p.Error = nil

	// Message
	log.Printf(" Starting Default Server at %s", addr)

	// Run the Server
	go p.startGoServerInternal()
}

// GracefulStop provides a way to stop the server properly with help of
// `context` package.
// The `waitForInterrupt` parameter is used to Listen for SIGINT or SIGKILL
// sent to the program.
// Additionally if an error was reported previous to the function call, it
// initiates the shutdown of the server.
// This is done using the member `Mserver.Error` which stores the Last error.
func (p *Mserver) GracefulStop(waitForInterrupt bool) error {

	// Wait for the Termination interrupt only if there are no previous errors
	if waitForInterrupt && p.Error == nil {
		log.Println(" Waiting for signals ...")
		<-p.stop
	}

	// Prevent Re-entry in case shutdown is called from an Error state in
	//  Goroutine
	if !p.started {
		return ErrServerNotStarted
	}

	log.Printf(" Shutting down Server at %s ...", p.Server.Addr)

	return p.stopServerInternal()
}

// ForceStop provides a forced way to terminate further processing by sending a false
// kill signal to a graseful shutdown running function
func (p *Mserver) ForceStop() {

	if !p.started {
		return
	}

	// Stop the Channel from receiving any further event
	signal.Stop(p.stop)
	// Send a Simulated kill event to Server
	p.stop <- os.Kill
}

// NewMserver creates a default Instance of the `Mserver` type and then calls
// the `StartDefaultServer` to begin default server operation.
func NewMserver(addr string, timeout time.Duration) *Mserver {
	m := &Mserver{}
	m.started = false
	m.Error = nil
	m.StartDefaultServer(addr, timeout)
	return m
}
