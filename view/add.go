package view

import "errors"

func (c *viewImpls) AddTwo(x, y int) (int, error) {

	if x == 0 {
		c.loggerX.Warn("x is zero, so logging")
		return 0, errors.New("x is zero, so returning erro")
	}

	return x + y, nil
}
