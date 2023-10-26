package contexts_test

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestSyscall(t *testing.T) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM, os.Interrupt)

	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			fmt.Println(time.Now(), "Just doing normal work")
		case <-c:
			fmt.Println("Shutting down gracefully")
			return
		}
	}
}

func TestContext(t *testing.T) {
	ctx := context.Background()
	newContext, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()

	for i := 0; i < 2; i++ {
		go func(ctx context.Context, i int) {
			ticker := time.NewTicker(500 * time.Millisecond)

			for {
				select {
				case <-ctx.Done():
					fmt.Printf("Stopping worker %v\n", i)
					return
				case tickTime := <-ticker.C:
					fmt.Printf("Worker %v is ticking at time:%v\n", i, tickTime)
				}
			}
		}(newContext, i)
	}

	time.Sleep(3 * time.Second)
}
