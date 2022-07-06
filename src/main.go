package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// App initialization and run to be managed by this method
// failure of Run needs to be handled.
func Run(log *log.Logger) {

	log.Println("Starting to run the GoHelloWorls application...")
	ticker := time.NewTicker(2 * time.Second)
	for {
		log.Println("Executing the run process...")
		<-ticker.C
	}
}

// main
//more changes
func main() {
	log := *log.New(os.Stdout, "go-hello-world: ", log.LstdFlags)
	fmt.Println("Initializing GoHelloWorld!!")
	Run(&log)
}
