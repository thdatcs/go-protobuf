package helpers

import (
	"go.uber.org/zap"
)

// InitZap initializes zap
func InitZap() error {
	instance, err := zap.NewProduction()
	if err != nil {
		return err
	}
	zap.ReplaceGlobals(instance)
	return nil
}
