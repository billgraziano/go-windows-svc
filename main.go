// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package main

import (
	"time"

	"github.com/pkg/errors"
)

// This is the name you will use for the NET START command
const svcName = "myservice"

// This is the name that will appear in the Services control panel
const svcNameLong = "My Service"

// Do your setup here
// Returning an error will prevent the service from starting
func setup() error {
	elog.Info(1, "In setup")

	// Returning an error prevents the service from starting
	return nil
}

// The wrapper of your app
func yourApp() {
	elog.Info(1, "In yourApp")

	// This is just some sample code to do something
	time.Sleep(1 * time.Second)
	elog.Info(1, "Still running")

	time.Sleep(2 * time.Second)
	elog.Info(1, "And running")

	time.Sleep(3 * time.Second)
	elog.Info(1, "Last message")
	elog.Info(1, "But the service will keep running")

	// Notice that if this exits, the service continues to run
}

func svcLauncher() error {
	err := setup()
	if err != nil {
		return errors.Wrap(err, "setup")
	}

	// Your service should be launched as a GO routine
	// It can launch other go routines, web servers, etc.
	go yourApp()

	return nil
}
