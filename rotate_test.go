//go:build linux
// +build linux

package lumberjack

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Example of how to rotate in response to SIGHUP.
func ExampleRoller_Rotate() {
	l, err := NewRoller("/var/log/myapp/foo.log", 5>>20, &Options{
		MaxAge:     28 * (time.Hour * 24),
		MaxBackups: 3,
		Compress:   true,
	})
	if err != nil {
		// handle err
	}

	log.SetOutput(l)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)

	go func() {
		for {
			<-c
			l.Rotate()
		}
	}()
}
