package service

import "github.com/pedrobarbosak/go-errors"

func Error() error {
	return errors.New("some service failed because value as:", 100)
}
