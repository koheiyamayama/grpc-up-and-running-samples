package repository

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	ErrNotFound ErrProductsRepo = errors.New("not found")
	ErrInternal ErrProductsRepo = errors.New("internal error")
)

type ErrProductsRepo interface {
	Error() string
}

type ProductsInfraError struct {
	Msg string
	Err ErrProductsRepo
}

func (e *ProductsInfraError) Error() string {
	return fmt.Errorf(e.Msg+": %w", e.Err).Error()
}

func (e *ProductsInfraError) Is(target error) bool {
	if t, ok := target.(ErrProductsRepo); ok {
		// TODO: learning goのreflectパッケージの章を読む。
		return reflect.DeepEqual(e.Err, t)
	}
	return false
}
