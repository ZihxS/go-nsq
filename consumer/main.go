package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nsqio/go-nsq"
)

var (
	consumer1 *nsq.Consumer
	consumer2 *nsq.Consumer
	consumer3 *nsq.Consumer
	consumer4 *nsq.Consumer
	consumer5 *nsq.Consumer
)

type msgHandler struct{}

func (h *msgHandler) HandleMessage(msg *nsq.Message) error {
	if len(msg.Body) == 0 {
		return nil
	}

	log.Printf("Message from %s (ID: %s): %s\n", msg.NSQDAddress, msg.ID, string(msg.Body))

	return nil
}

func run() {
	nsqConfig := nsq.NewConfig()

	// consumer 1
	consumer1, err := nsq.NewConsumer("go-nsq-1", "go-nsq-1", nsqConfig)
	if err != nil {
		panic(err)
	}

	consumer1.AddHandler(&msgHandler{})

	if err = consumer1.ConnectToNSQLookupd("localhost:4161"); err != nil {
		panic(err)
	}

	// consumer 2
	consumer2, err := nsq.NewConsumer("go-nsq-2", "go-nsq-2", nsqConfig)
	if err != nil {
		panic(err)
	}

	consumer2.AddHandler(&msgHandler{})

	if err = consumer2.ConnectToNSQLookupd("localhost:4161"); err != nil {
		panic(err)
	}

	// consumer 3
	consumer3, err := nsq.NewConsumer("go-nsq-3", "go-nsq-3", nsqConfig)
	if err != nil {
		panic(err)
	}

	consumer3.AddHandler(&msgHandler{})

	if err = consumer3.ConnectToNSQLookupd("localhost:4161"); err != nil {
		panic(err)
	}

	// consumer 4
	consumer4, err := nsq.NewConsumer("go-nsq-4", "go-nsq-4", nsqConfig)
	if err != nil {
		panic(err)
	}

	consumer4.AddHandler(&msgHandler{})

	if err = consumer4.ConnectToNSQLookupd("localhost:4161"); err != nil {
		panic(err)
	}

	// consumer 5
	consumer5, err := nsq.NewConsumer("go-nsq-5", "go-nsq-5", nsqConfig)
	if err != nil {
		panic(err)
	}

	consumer5.AddHandler(&msgHandler{})

	if err = consumer5.ConnectToNSQLookupd("localhost:4161"); err != nil {
		panic(err)
	}
}

func main() {
	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		run()
	}()

	<-exitChan

	consumer1.Stop()
	consumer2.Stop()
	consumer3.Stop()
	consumer4.Stop()
	consumer5.Stop()

	fmt.Println("")
	log.Println("Shutdown Signal Received!")
	log.Println("Bye Bye!")
}
