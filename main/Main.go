package main

import (
	"bufio"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"strconv"
	"sync"
	"time"
)

var broker = "172.16.210.157" // ~~ localhost
var topic = "myTopic"
var reader = bufio.NewReader(os.Stdin)
var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)
	go producer()
	go consumer()
	wg.Wait()
}
func producer() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
	})
	defer func() {
		p.Close()
		wg.Done()
	}()
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 5; i++ {
		_ = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic},
			Value:          []byte(strconv.Itoa(i)),
			Key:            nil,
			Timestamp:      time.Time{},
			TimestampType:  0,
			Opaque:         nil,
			Headers:        nil,
		}, nil)
	}
	for {
		mess, _ := reader.ReadString('\n')
		if mess == "stop" {
			fmt.Println("Stop signal!")
			os.Exit(1)
		}
		_ = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic},
			Value:          []byte(mess),
			Key:            []byte("key(" + mess + ")"),
			Timestamp:      time.Time{},
			TimestampType:  0,
			Opaque:         nil,
			Headers:        nil,
		}, nil)
	}
}

func consumer() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          "grId",
		"auto.offset.reset": "latest",
	})
	defer func() {
		_ = c.Close()
		wg.Done()
	}()
	if err != nil {
		fmt.Println(err)
	}
	_ = c.SubscribeTopics([]string{topic}, nil)
	for {
		mess, _ := c.ReadMessage(-1)
		fmt.Println(string(mess.Value))
	}
}
