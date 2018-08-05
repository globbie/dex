package knowdy

import (
	"testing"
	"github.com/globbie/dex/storage"
	"github.com/sirupsen/logrus"
	"os"
	"github.com/globbie/dex/storage/conformance"
)

var logger = &logrus.Logger{
	Out:       os.Stderr,
	Formatter: &logrus.TextFormatter{DisableColors: true},
	Level:     logrus.DebugLevel,
}

func TestKnowdy(t *testing.T) {
	logger := &logrus.Logger{
		Out:       os.Stderr,
		Formatter: &logrus.TextFormatter{DisableColors: true},
		Level:     logrus.DebugLevel,
	}
	newStorage := func() storage.Storage {
		return New(logger)
	}
	conformance.RunTests(t, newStorage)
	//conformance.RunTransactionTests(t, newStorage)
}
