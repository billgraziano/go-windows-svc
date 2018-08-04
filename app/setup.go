package app

import (
	"fmt"

	"golang.org/x/sys/windows/svc/debug"
)

// if setup returns an error, the service doesn't start
func setup(wl *debug.Log, svcName, sha1ver string) (server, error) {
	var s server

	// did we get a full SHA1?
	if len(sha1ver) == 40 {
		sha1ver = sha1ver[0:7]
	}

	if sha1ver == "" {
		sha1ver = "dev"
	}

	s.winlog = *wl

	// Note: any logging here goes to Windows App Log
	// I suggest you setup local logging
	s.winlog.Info(1, fmt.Sprintf("%s: setup (%s)", svcName, sha1ver))

	// read configuration
	// configure more logging

	return s, nil
}
