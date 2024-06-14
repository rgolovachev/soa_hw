package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"

	postspb "posts/proto"
	statpb "stat/proto"
	stat "stat/stat_grpc"

	_ "github.com/ClickHouse/clickhouse-go/v2"

	"github.com/IBM/sarama"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
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
	testAuthor  string = os.Getenv("TEST_AUTHOR")
)

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

	if testAuthor != "test" {
		grpc_conn, err = grpc.Dial("dns:///posts:11777", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatal(err)
		}

		grpc_client = postspb.NewPostsServiceClient(grpc_conn)
	}

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

						var author string

						if testAuthor != "test" {
							resp, err := grpc_client.CheckIfPostExists(context.Background(), &postspb.CheckIfPostExistsReq{PostId: decodedMsg.PostID})
							if err != nil || !resp.Exists {
								continue
							}

							author = resp.Author
						} else {
							author = testAuthor
						}

						if topic == "likes" {
							_, err = conn.Exec("INSERT INTO likes (post_id, author, username) VALUES (?, ?, ?)", decodedMsg.PostID, author, decodedMsg.Username)
						} else {
							_, err = conn.Exec("INSERT INTO views (post_id, author, username) VALUES (?, ?, ?)", decodedMsg.PostID, author, decodedMsg.Username)
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

func handleGrpc() {
	lis, err := net.Listen("tcp", "0.0.0.0:11888")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	stat_service, err := stat.NewServer()
	if err != nil {
		log.Fatalf("failed to open DB: %v", err)
	}

	statpb.RegisterStatServiceServer(grpcServer, stat_service)

	log.Println("stat grpc service started")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("stat grpc service failed")
	}
}

func main() {
	var wg sync.WaitGroup

	env := os.Getenv("STAT_ENV")
	fmt.Fprintf(os.Stderr, "env: %v\n", env)

	if env != "service_testing" {
		wg.Add(1)
		httpServer := func() {
			defer wg.Done()

			handleHTTP()
		}
		go httpServer()
	}

	if env != "service_testing" {
		wg.Add(1)
		kafkaConsumer := func() {
			defer wg.Done()

			consumeKafka()
		}

		go kafkaConsumer()
	}

	wg.Add(1)
	grpcServer := func() {
		defer wg.Done()

		handleGrpc()
	}

	go grpcServer()

	wg.Wait()
}
