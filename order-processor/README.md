Project 1: Event-Driven Order Processing System
Goal: Simulate a real-world financial clearing application that handles high-throughput event ingestion and processing.

Key Features:
- Written in Go using idiomatic practices.
- Kafka integration: one topic for incoming orders, one for status updates.
- Use Go channels and goroutines to handle concurrent processing.
- In-memory deduplication and rate limiting.
- Unit tests (e.g., with testing package), integration test stubs with Kafka.

What You’ll Learn:
- Concurrency design in Go (using sync, context, channels).
- Kafka client in Go (confluent-kafka-go or segmentio/kafka-go).
- Structuring Go code for testing and separation of concerns.
- Logging, graceful shutdown, and recovery handling.

Optional Stretch:
- Docker-compose to spin up Kafka locally.
- JSON schema validation for incoming orders.

https://github.com/edenhill/kcat
https://github.com/provectus/kafka-ui

++++

Perfect—you're almost there! Here's how to gracefully shut down your services and interact with Kafka using `kafkacat` or a Kafka UI.

---

## 🛑 Stop Everything

### 🧹 To Stop Kafka and Zookeeper:
Run this from your project root (where `docker-compose.yml` is located):

```bash
docker-compose down
```

This stops and **removes** the containers.

> To **just stop** without removing them, use:
```bash
docker-compose stop
```
---

## 👁️ View and Send Kafka Messages

### Option 1: Use `kafkacat` (CLI tool)
Install `kafkacat` (now called `kcat`):

#### On Mac:
```bash
brew install kcat
```

#### On Linux:
```bash
sudo apt install kafkacat
```

### ✅ Consume messages (e.g., from `incoming-orders`):

```bash
kcat -b localhost:9092 -t incoming-orders -C
```

### ✅ Produce a test message:

```bash
echo '{"id":"123", "symbol":"AAPL", "quantity":100, "timestamp":1680000000}' | \
kcat -b localhost:9092 -t incoming-orders -P
```

You should see your Go app pick this up if it’s running.

---

### Option 2: Use Kafka UI (Web Interface)

1. Add this to your `docker-compose.yml`:

```yaml
  kafka-ui:
    image: provectuslabs/kafka-ui
    ports:
      - "8080:8080"
    environment:
      - KAFKA_CLUSTERS_0_NAME=local
      - KAFKA_CLUSTERS_0_BOOTSTRAPSERVERS=kafka:9092
    depends_on:
      - kafka
```

2. Restart Docker Compose:

```bash
docker-compose down
docker-compose up -d
```

3. Open your browser at:  
👉 (http://localhost:8080)

- You’ll see a UI where you can browse topics, consume messages, and publish new ones.

++++
