package view

import "github.com/Chandan94f/mb-utility/utility/logger"

type View interface {
	AddTwo(x, y int) (int, error)
}

type viewImpls struct {
	logger  *logger.Logger
	loggerX *logger.LoggerX
}

func NewCaller(logger *logger.Logger, loggerX *logger.LoggerX) View {
	return &viewImpls{logger: logger, loggerX: loggerX}
}
