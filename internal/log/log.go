package log

import (
	log "go.uber.org/zap"
)

// New Create a logger using the environment variable to define the log level
// it returns an error when the environment variable is an invalid log level
func New() (*log.SugaredLogger, error) {
	logger, err := log.Config{
		Level:       log.NewAtomicLevelAt(log.DebugLevel),
		Development: false,
		Sampling: &log.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    log.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build()
	if err != nil {
		return nil, err
	}

	return logger.Sugar(), nil
}
