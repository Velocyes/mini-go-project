package consts

import "errors"

var (
	ErrNilConfig = errors.New("config is nil")
	ErrInvalidUrlVariable = errors.New("invalid url variable")
	ErrInvalidProductID = errors.New("invalid product id")
)
