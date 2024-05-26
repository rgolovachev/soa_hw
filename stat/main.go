package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"

	postspb "posts/proto"

	_ "github.com/ClickHouse/clickhouse-go/v2"

	"github.com/IBM/sarama"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// just following 2.1 from hw statement
func dummyHandler(w http.ResponseWriter, r *http.Request) {
}

func handleHTTP() {
	http.HandleFunc("/dummy", dummyHandler)

	log.Println("stat service started")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

type Message struct {
	PostID   string `json:"post_id"`
	Username string `json:"username"`
}

var (
	grpc_conn   *grpc.ClientConn
	grpc_client postspb.PostsServiceClient
)

func checkIfPostExists(postID string) (bool, error) {
	grpc_resp, err := grpc_client.CheckIfPostExists(context.Background(), &postspb.CheckIfPostExistsReq{PostId: postID})
	if err != nil {
		return false, err
	}

	return grpc_resp.Exists, nil
}

func consumeKafka() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	brokers := []string{os.Getenv("KAFKA_BROKER")}
	topics := []string{"likes", "views"}

	client, err := sarama.NewClient(brokers, config)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	doneCh := make(chan struct{})

	conn, err := sql.Open("clickhouse", "tcp://clickhouse:9000?username=default&password=")

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	grpc_conn, err = grpc.Dial("dns:///posts:11777", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	grpc_client = postspb.NewPostsServiceClient(grpc_conn)

	for _, topic := range topics {
		partitions, err := consumer.Partitions(topic)
		if err != nil {
			log.Fatal(err)
		}

		for _, partition := range partitions {
			go func(topic string, partition int32) {
				partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
				if err != nil {
					log.Fatal(err)
				}
				defer partitionConsumer.Close()

				for {
					select {
					case msg := <-partitionConsumer.Messages():
						log.Printf("Received message from topic %s: %s", topic, string(msg.Value))

						var decodedMsg Message
						err := json.Unmarshal(msg.Value, &decodedMsg)
						if err != nil {
							log.Printf("Received unexpected error while unmarshalling: %v\n", err)
							continue
						}

						if exists, err := checkIfPostExists(decodedMsg.PostID); err != nil || !exists {
							continue
						}

						if topic == "likes" {
							_, err = conn.Exec("INSERT INTO likes (post_id, username) VALUES (?, ?)", decodedMsg.PostID, decodedMsg.Username)
						} else {
							_, err = conn.Exec("INSERT INTO views (post_id, username) VALUES (?, ?)", decodedMsg.PostID, decodedMsg.Username)
						}

						if err != nil {
							log.Printf("Recevied unexpected error while inserting to CH: %v\n", err)
							continue
						}

						log.Println("data was inserted successfully")
					case err := <-partitionConsumer.Errors():
						log.Printf("Error consuming messages: %v", err)
					case <-doneCh:
						return
					}
				}
			}(topic, partition)
		}
	}

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, os.Interrupt)
	<-sigterm

	close(doneCh)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	httpServer := func() {
		defer wg.Done()

		handleHTTP()
	}
	kafkaConsumer := func() {
		defer wg.Done()

		consumeKafka()
	}

	go httpServer()
	go kafkaConsumer()

	wg.Wait()
}
