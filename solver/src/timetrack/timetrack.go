package timetrack

import (
	"fmt"
	"time"

	"github.com/Mandala/go-log"
)

// Keep track of how much time a specific solution took
func TimeTrack(logger *log.Logger, start time.Time, name string) {
	elapsed := time.Since(start)
	logger.Info(fmt.Sprintf("%s took %s", name, elapsed))
}
