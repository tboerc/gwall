package main

import "errors"

var (
	errInvalidIndex = errors.New("error: index is not valid")
	errNoIP         = errors.New("error: no IP provided")
	errNoIndex      = errors.New("error: index provided")
	errConfigWrite  = errors.New("error: config file write error")
	errExecPath     = errors.New("error: could not get executable path")
)
