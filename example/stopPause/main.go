// Copyright 2015 Daniel Theophanes.
// Use of this source code is governed by a zlib-style
// license that can be found in the LICENSE file.

// simple does nothing except block while running the service.
package main

import (
	"log"
	"os"
	"time"

	"github.com/lucasdecamargo/kardianos"
)

var logger kardianos.Logger

type program struct{}

func (p *program) Start(s kardianos.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() {
	// Do work here
}
func (p *program) Stop(s kardianos.Service) error {
	// Stop should not block. Return with a few seconds.
	<-time.After(time.Second * 13)
	return nil
}

func main() {
	svcConfig := &kardianos.Config{
		Name:        "GoServiceExampleStopPause",
		DisplayName: "Go Service Example: Stop Pause",
		Description: "This is an example Go service that pauses on stop.",
	}

	prg := &program{}
	s, err := kardianos.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) > 1 {
		err = kardianos.Control(s, os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
