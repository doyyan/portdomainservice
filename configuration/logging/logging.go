package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

// NewLogger creates a new logger using the specified options.
func NewLogger() *logrus.Logger {
	log := logrus.New()
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&logrus.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(logrus.WarnLevel)

	return log
}
