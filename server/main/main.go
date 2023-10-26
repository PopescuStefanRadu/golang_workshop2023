package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, stopNotifying := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stopNotifying()

	server := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			if _, err := w.Write([]byte("Hello World")); err != nil {
				fmt.Printf("Encountered an error when when writing response: %s\n", err)
			}
		}),
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		fmt.Println("Shutting down gracefully")
		deadline, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if err := server.Shutdown(deadline); err != nil {
			fmt.Printf("Error: could not shutdown gracefully: %s\n", err.Error())
		} else {
			fmt.Println("Shutdown completed")
		}
	}()

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println(fmt.Errorf("unexpected server error: %w", err))
	}
	wg.Wait()
}
