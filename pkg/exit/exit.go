package exit

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Init callback function que permite graceful shutdown do servidor
func Init(cb func()) {
	sigs := make(chan os.Signal, 1)
	terminate := make(chan bool)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		log.Println("Exit reason: ", sig)
		close(terminate)
	}()

	<-terminate
	cb()
}
