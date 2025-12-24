package main

//import (
//	"context"
//	"fmt"
//	"log"
//	"strconv"
//	"time"
//
//	"github.com/apache/rocketmq-clients/golang"
//)
//
//const (
//	Topic     = "TestTopic"
//	GroupName = "test"
//	Endpoint  = "127.0.0.1:9876"
//)
//
//func main() {
//	golang.ResetLogger()
//	// new producer instance
//	producer, err := golang.NewProducer(&golang.Config{
//		Endpoint: Endpoint,
//	},
//		golang.WithTopics(Topic),
//	)
//	if err != nil {
//		log.Fatal(err)
//	}
//	// start producer
//	err = producer.Start()
//	if err != nil {
//		log.Fatal(err)
//	}
//	// gracefule stop producer
//	defer producer.GracefulStop()
//
//	for i := 0; i < 10; i++ {
//		// new a message
//		msg := &golang.Message{
//			Topic: Topic,
//			Body:  []byte("this is a message : " + strconv.Itoa(i)),
//		}
//		// set keys and tag
//		msg.SetKeys("a", "b")
//		msg.SetTag("ab")
//		// send message in sync
//		resp, err := producer.Send(context.TODO(), msg)
//		if err != nil {
//			log.Fatal(err)
//		}
//		for i := 0; i < len(resp); i++ {
//			fmt.Printf("%#v\n", resp[i])
//		}
//		// wait a moment
//		time.Sleep(time.Second * 1)
//	}
//}
