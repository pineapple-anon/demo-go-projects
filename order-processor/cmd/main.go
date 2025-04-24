package main

import (
    "context"
    "encoding/json"
    "log"
    "eg.com/order-processor/internal/kafka"
    "eg.com/order-processor/internal/processor"
	"eg.com/order-processor/internal/models"
)

func main() {
    consumer := kafka.NewConsumer("incoming-orders", "localhost:9092")
    proc := processor.New()

    for {
        msg, err := consumer.ReadMessage(context.Background())
        if err != nil {
            log.Fatal("Error reading:", err)
        }

        var order models.Order
        if err := json.Unmarshal(msg.Value, &order); err != nil {
            log.Println("Invalid order:", err)
            continue
        }

        if proc.Process(order) {
            log.Println("Processed order:", order.ID)
            // Send to processed-orders topic (Kafka producer code here)
        } else {
            log.Println("Duplicate order:", order.ID)
        }
    }
}
