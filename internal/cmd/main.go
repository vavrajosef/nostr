package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-nostr/nostr"
	"golang.org/x/sync/errgroup"
)

func main() {
	// NOTE: Create a new context for managing the lifetime of the main function
	ctx := context.Background()
	cl := nostr.NewClient()

	// NOTE: Initialize a struct containing all the service servers
	sc := struct {
		cl *http.Server
	}{
		cl: &http.Server{
			Addr:         "0.0.0.0:4200",
			Handler:      cl,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}

	// NOTE: Create an error group to manage the concurrent execution of service servers
	g, _ := errgroup.WithContext(ctx)

	// NOTE: Add service server Serve functions to the error group
	g.Go(sc.cl.ListenAndServe)

	// NOTE: Wait for all service servers to complete or return an error
	if err := g.Wait(); err != nil {
		panic(err)
	}
}
