package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	kafka "github.com/opensourceways/kafka-lib/agent"
)

func handler(msg []byte, header map[string]string) error {
	fmt.Println("received msg", msg, header)
	return fmt.Errorf("test err")
	//return nil
}

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s groupid\n",
			os.Args[0])
		os.Exit(1)
	}

	conf := &kafka.Config{
		Address: "127.0.0.1:9092",
	}

	kafka.Init(conf, nil)

	groupid := os.Args[1]
	//conf := kafka.ConfigMap{
	//	"bootstrap.servers":  "127.0.0.1:9092",
	//	"enable.auto.commit": "false",
	//}
	//conf["group.id"] = groupid
	// auto.offset.reset参数只控制首次加入的consumer的行为
	// earliest: 当各分区下有已提交的offset时，从提交的offset开始消费；无提交的offset时，从头开始消费
	// latest: 当各分区下有已提交的offset时，从提交的offset开始消费；无提交的offset时，消费新产生的该分区下的数据
	// none: topic各分区都存在已提交的offset时，从offset后开始消费；只要有一个分区不存在已提交的offset，则抛出异常

	//conf["auto.offset.reset"] = "earliest"

	e := kafka.SubscribeTopics([]string{"topic1", "topic2"}, groupid, handler)
	//c, err := kafka.NewConsumer(&conf)

	if e != nil {
		fmt.Printf("Failed to create consumer: %s", e)
		os.Exit(1)
	}

	//topic := make([]string, 0)
	//topic = append(topic, "p3")
	//err = c.SubscribeTopics([]string{"topic1", "topic2"}, handler)
	// Set up a channel for handling Ctrl-C, etc
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Process messages
	run := true
	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		}
	}
}
