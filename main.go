package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	JwtSecret string
)

func init() {
	// !!! NB bad practice to provide a hardcoded secret - would be a chance of this leaking to PROD !!!
	jwtSecret := flag.String("jwtSecret", "TESTING-ONLY-651c969ee2f44c7a917b9c5eb41fbbea", "The JWT Secret")
	flag.Parse()

	if *jwtSecret == "" {
		log.Fatal("JWT Secret not presented via command line args")
	}

	JwtSecret = *jwtSecret
}

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	store, err := NewPostgreStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":3000", store)
	go server.Run()

	<-done
}
