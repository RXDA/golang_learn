package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func consume(){
	topic := "topic-ol-iot-env-data"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "172.17.3.24:9092,172.17.3.22:9092,172.17.3.23:9092", topic, partition)
	if err != nil{
		panic(err)
	}
	//conn.SetReadDeadline(time.Now().Add(10*time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		_, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b))
	}

	batch.Close()
	conn.Close()
}
