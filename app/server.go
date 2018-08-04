package app

import (
	"golang.org/x/sys/windows/svc/debug"
)

type server struct {
	winlog debug.Log
	// a local logger
	// a database connection
	// your app configuration
}
