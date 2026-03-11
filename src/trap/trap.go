package trap

import (
	"context"
	"os"
	"os/signal"
)

func Trap(run func(), signals ...os.Signal) {
	go func() {
		ctx, stop := signal.NotifyContext(context.Background(), signals...)
		defer stop()
		<-ctx.Done()
		run()
		os.Exit(1)
	}()
}
