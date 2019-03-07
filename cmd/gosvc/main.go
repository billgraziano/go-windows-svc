// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package main

import (
	"github.com/billgraziano/go-windows-svc/app"
	"github.com/pkg/errors"
)

// This is the name you will use for the NET START command
const svcName = "gosvc"

// This is the name that will appear in the Services control panel
const svcNameLong = "GO Service"

// This is assigned the full SHA1 hash from GIT
var sha1ver string

func svcLauncher() error {

	err := app.Run(elog, svcName, sha1ver)
	if err != nil {
		return errors.Wrap(err, "app.run")
	}

	return nil
}
