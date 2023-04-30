package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/nsqio/go-nsq"
)

func publish(wg *sync.WaitGroup, producer *nsq.Producer, topic string, channel string, delay time.Duration, amount int) {
	defer wg.Done()
	for i := 1; i <= amount; i++ {
		go func(topic string, i int) {
			if err := producer.Publish(topic, []byte(fmt.Sprintf("from producer to %v: %v.", topic, i))); err != nil {
				panic(err)
			} else {
				log.Printf("Delivered message to %v\n", topic)
			}
		}(topic, i)
		time.Sleep(delay)
	}
}

func main() {
	nsqConfig := nsq.NewConfig()

	producer, err := nsq.NewProducer("127.0.0.1:4150", nsqConfig)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup

	min, max := 1, 5

	for i := min; i <= max; i++ {
		wg.Add(1)
		go publish(&wg, producer, fmt.Sprintf("go-nsq-%v", i), fmt.Sprintf("go-nsq-%v", i), ((time.Second / 4) * time.Duration(i)), ((max+1)*50)-(min*10))
	}

	wg.Wait()

	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-exitChan

	producer.Stop()
	fmt.Println("")
	log.Println("Shutdown Signal Received!")
	log.Println("Bye Bye!")
}
