package util

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func RunServicesSafely(services ...func(context.Context)) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
    defer stop()

    // Use WaitGroup to track goroutine completion
    var wg sync.WaitGroup
    
	for _, service := range services {
		wg.Add(1)
		go func(serve func(context.Context)) {
			defer wg.Done()
			serve(ctx)
		}(service)
	}

    // Wait for shutdown signal
    <-ctx.Done()
    log.Println("Graceful shutdown triggered.")
    
    // Wait for both goroutines to finish with timeout
    done := make(chan struct{})
    go func() {
        wg.Wait()
        close(done)
    }()
    
    select {
    case <-done:
        log.Println("All services stopped gracefully")
    case <-time.After(5 * time.Second):
        log.Println("Shutdown timeout - forcing exit")
    }
}