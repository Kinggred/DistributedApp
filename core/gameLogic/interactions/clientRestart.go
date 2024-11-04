package interactions

import (
	"context"
	glo "tic-tac-toe/core/global"
	"tic-tac-toe/core/networking"
	"time"
)

// Figuring out how to allow this took me 3 days
// I hate this approach, but it should do
func RunClientInGoroutine(ctx context.Context) {
	glo.KillClient.Kill()

	newCtx, newKillFunction := context.WithCancel(ctx)
	glo.KillClient.SetKillFunction(newKillFunction)

	for {
		time.Sleep(100 * time.Microsecond)
		if !glo.KillClient.GetStatus() {
			break
		}
	}
	go func() {
		networking.RunClient(newCtx)
	}()
}
