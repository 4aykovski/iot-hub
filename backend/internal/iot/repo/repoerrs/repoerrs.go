package repoerrs

import "errors"

var (
	ErrNoDevice = errors.New("device not found")
	ErrNoData   = errors.New("data not found")
)
