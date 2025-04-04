package main

import (
	"os"
	"os/signal"
	"syscall"
)

// API Tutorial based on
// https://www.youtube.com/watch?v=pwZuNmAzaH8&list=PL0xRBLFXXsP6nudFDqMXzrvQCZrxSOm-2

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	server := NewAPIServer(":3000")
	go server.Run()

	<-done
}
