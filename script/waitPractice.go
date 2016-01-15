package main
import (
	"log"
	"time"
)

func main() {
	channel2()
}

func normal() {
	log.Print("Started:")

	log.Print("sleep 1 started")
	time.Sleep(1 * time.Second)
	log.Print("sleep 1 finished")

	log.Print("sleep 2 started")
	time.Sleep(2 * time.Second)
	log.Print("sleep 2 finished")

	log.Print("sleep 3 started")
	time.Sleep(3 * time.Second)
	log.Print("sleep 3 finished")

	log.Print("Finished")
}

func parallel() {
	log.Print("Parallel Started")

	go func() {
		log.Print("sleep 3 started")
		time.Sleep(3 * time.Second)
		log.Print("sleep 3 finished")
	}()

	go func() {
		log.Print("sleep 2 started")
		time.Sleep(2 * time.Second)
		log.Print("sleep 2 finished")
	}()

	go func() {
		log.Print("sleep 1 started")
		time.Sleep(1 * time.Second)
		log.Print("sleep 1 finished")
	}()

	log.Print("Parallel is to Finish")
	time.Sleep(4 * time.Second)

	log.Print("Parallel Finished")
}

func channel() {
	log.Print("Channel Started")

	sleep1Finished := make(chan bool)
	sleep2Finished := make(chan bool)
	sleep3Finished := make(chan bool)

	go func() {
		log.Print("sleep 1 started")
		time.Sleep(1 * time.Second)
		log.Print("sleep 1 finished")
		sleep1Finished <- true
	}()

	go func() {
		log.Print("sleep 2 started")
		time.Sleep(2 * time.Second)
		log.Print("sleep 2 finished")
		sleep2Finished <- true
	}()

	go func() {
		log.Print("sleep 3 started")
		time.Sleep(3 * time.Second)
		log.Print("sleep 3 finished")
		sleep3Finished <- true
	}()

	log.Print("Channel is to Finish.")
	<- sleep1Finished
	log.Print("Channel is to Finish...")
	<- sleep2Finished
	log.Print("Channel is to Finish.....")
	<- sleep3Finished

	log.Print("Channel Finished")
}

func channel2() {
	log.Print("Channel2 Started")

	finished := make(chan bool)

	go func() {
		log.Print("sleep 1 started")
		time.Sleep(1 * time.Second)
		log.Print("sleep 1 finished")
		finished <- true
	}()

	go func() {
		log.Print("sleep 2 started")
		time.Sleep(2 * time.Second)
		log.Print("sleep 2 finished")
		finished <- true
	}()

	go func() {
		log.Print("sleep 3 started")
		time.Sleep(3 * time.Second)
		log.Print("sleep 3 finished")
		finished <- true
	}()

	log.Print("Channel2 is to Finish.")
	<- finished
	log.Print("Channel2 is to Finish...")
	<- finished
	log.Print("Channel2 is to Finish.....")
	<- finished

	log.Print("Channel2 Finished")
}