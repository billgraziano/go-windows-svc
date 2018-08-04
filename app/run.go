package app

import (
	"github.com/pkg/errors"
	"golang.org/x/sys/windows/svc/debug"
)

// Run launches the service
func Run(wl debug.Log, svcName, sha1ver string) error {

	s, err := setup(wl, svcName, sha1ver)
	if err != nil {
		return errors.Wrap(err, "setup")
	}

	// Your service should be launched as a GO routine
	go yourApp(s)

	return nil
}
