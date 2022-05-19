package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var (
	staticPages = populateStaticPage()
)

func main() {
	go dbConnection()
	go api()

	sigs := make(chan os.Signal, 1)
	done := make(chan bool)

	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		sig := <-sigs
		_ = sig
		done <- true
	}()

	<-done
	fmt.Println("program closed")
}
