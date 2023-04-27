package loggers

import (
	"log"

	"go.uber.org/zap"
)

type Logger struct {
	Zap *zap.Logger
}

const (
	DEVELOPMENT = "DEVELOPMENT"
	PRODUCTION  = "PRODUCTION"
)

func (l *Logger) Init(mode string) {
	var err error
	if mode == DEVELOPMENT {
		l.Zap, err = zap.NewDevelopment()
		if err != nil {
			log.Fatalf("can't initialize zap logger mode %v: %v", mode, err)
		}
		defer l.Zap.Sync()
	}
	if mode == PRODUCTION {
		l.Zap, err = zap.NewProduction()
		if err != nil {
			log.Fatalf("can't initialize zap logger mode %v: %v", mode, err)
		}
		defer l.Zap.Sync()
	}
}
