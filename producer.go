package main

import (
	"math/rand"
	"os"

	kafka "github.com/opensourceways/kafka-lib/agent"
)

func main() {

	//configFile := os.Args[1]
	conf := &kafka.Config{
		Address: "127.0.0.1:9092",
	}

	kafka.Init(conf, nil)

	topic := os.Args[1]

	users := [...]string{"eabara", "jsmith", "sgarcia", "jbernard", "htanaka", "awalther"}
	items := [...]string{"book", "alarm clock", "t-shirts", "gift card", "batteries"}

	for n := 0; n < 10; n++ {
		key := users[rand.Intn(len(users))]
		data := items[rand.Intn(len(items))]
		header := make(map[string]string)
		header[key] = data
		kafka.Publish(topic, header, []byte{})
	}
}
