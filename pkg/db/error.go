package db

import (
	"github.com/gocraft/dbr"
)

// package errors
var (
	ErrNotFound           = dbr.ErrNotFound
	ErrNotSupported       = dbr.ErrNotSupported
	ErrTableNotSpecified  = dbr.ErrTableNotSpecified
	ErrColumnNotSpecified = dbr.ErrColumnNotSpecified
	ErrInvalidPointer     = dbr.ErrInvalidPointer
	ErrPlaceholderCount   = dbr.ErrPlaceholderCount
	ErrInvalidSliceLength = dbr.ErrInvalidSliceLength
	ErrCantConvertToTime  = dbr.ErrCantConvertToTime
	ErrInvalidTimestring  = dbr.ErrInvalidTimestring
)
