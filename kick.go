// The Isomorphic Go Project
// Copyright (c) Wirecog, LLC. All rights reserved.
// Use of this source code is governed by a BSD-style
// license, which can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

var appPath string
var mainSourceFile string
var gopherjsAppPath string

func buildGopherJSProject() {

	if gopherjsAppPath != "" {

		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal("Encountered an error while attempting to get the cwd: ", err)
		} else {
			os.Chdir(gopherjsAppPath)
			gjsCommand := exec.Command("gopherjs", "build")
			gjsCommand.Stdout = os.Stdout
			gjsCommand.Stderr = os.Stderr
			gjsCommand.Start()
			os.Chdir(cwd)
		}
	}

}

func restart(cmd *exec.Cmd) *exec.Cmd {
	var newCommand *exec.Cmd
	stop(cmd)
	newCommand = start()
	return newCommand

}

func initializeWatcher(shouldRestart chan bool, dirList []string) {

	supportedExtensions := map[string]int{".go": 1, ".html": 1, ".tmpl": 1}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:

				if event.Op == fsnotify.Write || event.Op == fsnotify.Rename {
					if _, ok := supportedExtensions[filepath.Ext(event.Name)]; ok {
						shouldRestart <- true
					}
				}

			case err := <-watcher.Errors:
				if err != nil {
					log.Println("error:", err)
				}
			}
		}
	}()

	err = watcher.Add(appPath)
	if err != nil {
		log.Fatal(err)
	}
	// watch subdirectories also
	for _, element := range dirList {
		watcher.Add(element)
	}

	<-done

}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func main() {

	flag.StringVar(&appPath, "appPath", "", "The path to your Go project")
	flag.StringVar(&mainSourceFile, "mainSourceFile", "", "The Go source file with the main func()")
	flag.StringVar(&gopherjsAppPath, "gopherjsAppPath", "", "The path to your GopherJS project (optional)")
	flag.Parse()

	// Exit if no appPath is supplied
	if appPath == "" {
		fmt.Println("You must supply the appPath parameter")
		os.Exit(1)
	}

	if appPathExists, appPathErr := pathExists(appPath); appPathExists != true || appPathErr != nil {
		fmt.Println("The path you specified to your Go application project does not exist.")
		os.Exit(1)
	}

	if mainSourceFile == "" {
		fmt.Println("You must supply the mainSourceFile parameter")
		os.Exit(1)
	}

	if sourceFileExists, sourceFilePathErr := pathExists(appPath + "/" + mainSourceFile); sourceFileExists != true || sourceFilePathErr != nil {
		fmt.Println("The path to the main source file you provided does not exist.")
		os.Exit(1)
	}

	if gopherjsAppPath != "" {
		if gopherjsFileExists, gopherjsFileErr := pathExists(gopherjsAppPath); gopherjsFileExists != true || gopherjsFileErr != nil {
			fmt.Println("The path you specified to the GopherJS application project does not exist.")
			os.Exit(1)
		}
	}

	dirList := []string{}
	filepath.Walk(appPath, func(path string, f os.FileInfo, err error) error {

		if f.IsDir() == true {
			dirList = append(dirList, path)
		}
		return nil
	})

	shouldRestart := make(chan bool, 1)

	go initializeWatcher(shouldRestart, dirList)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	cmd := start()

	for {

		select {

		case <-interrupt:
			stop(cmd)
			os.Exit(0)

		case <-shouldRestart:
			fmt.Println("\nInstant KickStart Applied! (Recompiling and restarting project.)")
			cmd = restart(cmd)

		}

	}

}
