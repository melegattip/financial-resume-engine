package logger

import (
	"context"
	"log"

	"github.com/melegattip/financial-resume-engine/internal/core/logs"
)

func Error(ctx context.Context, err error, message string, tags logs.Tags) {
	log.Printf("ERROR: %s - %v - %+v", message, err, tags)
}
