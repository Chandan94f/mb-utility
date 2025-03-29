package view

import "github.com/mb-utility/utility/logger"

type View interface {
	addTwo(x, y int) (int, error)
}

type viewImpls struct {
	logger  *logger.Logger
	loggerX *logger.LoggerX
}

func NewCaller(logger *logger.Logger, loggerX *logger.LoggerX) View {
	return &viewImpls{logger: logger, loggerX: loggerX}
}
