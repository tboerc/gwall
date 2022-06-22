package messages

import (
	"errors"
)

var (
	ErrServerRun    = errors.New("error: could not start the server")
	ErrDatabaseConn = errors.New("error: could not connect to database")

	ErrInvalidHash         = errors.New("erro: the encoded hash is not in the correct format")
	ErrIncompatibleVersion = errors.New("error: incompatible version of argon2")
)
