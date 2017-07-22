// The Isomorphic Go Project
// Copyright (c) Wirecog, LLC. All rights reserved.
// Use of this source code is governed by a BSD-style
// license, which can be found in the LICENSE file.

// +build windows

package main

import (
	"os"
	"os/exec"
	"strconv"
)

func start() *exec.Cmd {
	buildGopherJSProject()
	cmd := exec.Command("go", "run", appPath+"/"+mainSourceFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	return cmd
}

func stop(cmd *exec.Cmd) {
	stop := exec.Command("TASKKILL", "/T", "/F", "/PID", strconv.Itoa(cmd.Process.Pid))
	stop.Run()
}
