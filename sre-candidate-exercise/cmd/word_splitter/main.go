package main

import (
	"context"
	"github.com/covrco/sre-candidate-exercise/cmd/word_splitter/server"
	"github.com/covrco/sre-candidate-exercise/protocol/sentencepb"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("The server is starting...")

	// Some interrupt and error handler stuff to make sure that the process terminates properly
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	group, ctx := errgroup.WithContext(ctx)

	// Set up the TCP listener on port 9001
	listener, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("failed to set up listener on port 9001: %s", err)
	}

	grpcServer := grpc.NewServer()

	// Attach the implementation of the server from the protocol
	sentencepb.RegisterSentenceServer(grpcServer, &server.SentenceServerImpl{})

	// Start serving on the listener, but in a thread that is listening for returned errors
	group.Go(func() error {
		return grpcServer.Serve(listener)
	})
	log.Println("The server is started")

	// Block until an error occurs on a thread or the process is signaled to exit
	select {
	case <-interrupt:
		break
	case <-ctx.Done():
		break
	}

	log.Println("Received an exit signal. Starting the shutdown process.")

	cancel()

	// Stop the GRPC server if it still exists
	if grpcServer != nil {
		log.Println("Stopping the GRPC listener.")
		grpcServer.GracefulStop()
	}

	err = group.Wait()

	if err != nil {
		// Show the error if that is what started the shutdown
		log.Printf("The service is shutting down due to an error: %v", err)
		os.Exit(2)
	}
}
