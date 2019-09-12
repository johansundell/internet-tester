package main

import (
	"time"

	"github.com/kardianos/service"
)

var logger service.Logger

type program struct {
	exit chan struct{}
}

func (p *program) Start(s service.Service) error {
	if service.Interactive() {
		logger.Info("Running in terminal.")
	} else {
		logger.Info("Running under service manager.")
	}
	p.exit = make(chan struct{})

	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}

func (p *program) run() error {
	logger.Infof("I'm running %v.", service.Platform())

	ticker := time.NewTicker(time.Duration(settings.SecToWait) * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				// TODO: Test internet and write err to log
				if err := checkInternet(settings.UrlToCheck); err != nil {
					logger.Error("No internet")
				} else {
					if settings.Debug {
						logger.Info("We have internet :)")
					}
				}

			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	for {
		select {
		case <-p.exit:
			close(quit)
			return nil
		}
	}
}

func (p *program) Stop(s service.Service) error {
	// Any work in Stop should be quick, usually a few seconds at most.
	logger.Info("I'm Stopping!")
	close(p.exit)
	return nil
}
