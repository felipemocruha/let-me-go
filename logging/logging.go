package logging

import (
	"go.uber.org/zap"
)

func NewLogger() (*zap.SugaredLogger, error) { 
	prod, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return prod.Sugar(), nil
}
