package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func Run(log *log.Logger) {

	log.Println("Starting to run the GoHelloWorls application...")
	ticker := time.NewTicker(2 * time.Second)
	for {
		log.Println("Executing the run process...")
		<-ticker.C
	}
}

func main() {
	log := *log.New(os.Stdout, "go-hello-world: ", log.LstdFlags)
	fmt.Println("Initializing GoHelloWorld!!")
	Run(&log)
}
