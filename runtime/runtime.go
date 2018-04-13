package runtime

import (
	"github.com/chespinoza/test-zenroom-go/pkg/globals"
	"go.uber.org/zap"
)

func loop(stream Stream) {
	for {
		select {
		case input := <-stream.New():
			globals.Logger.Info("Incoming data", zap.String("data", input))
			// Encode ..
		}
	}
}
