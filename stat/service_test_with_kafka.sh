#!/bin/bash

docker-compose -f service_test_with_kafka.yml build > /dev/null 2>&1
echo "build done"
docker-compose -f service_test_with_kafka.yml up -d > /dev/null 2>&1
echo "up done"

sleep 40

docker exec -it clickhouse clickhouse-client --query="TRUNCATE TABLE likes;" > /dev/null 2>&1
echo '{"post_id": "228", "username": "aboba"}' | docker exec -i kafka kafka-console-producer.sh --broker-list 0.0.0.0:9092 --topic likes
RAWPOST=$(docker exec -it clickhouse clickhouse-client --query="SELECT post_id FROM likes" | tail -n +3)
RAWAUTHOR=$(docker exec -it clickhouse clickhouse-client --query="SELECT author FROM likes" | tail -n +3)
RAWUSERNAME=$(docker exec -it clickhouse clickhouse-client --query="SELECT username FROM likes" | tail -n +3)

POST=$(echo "$RAWPOST" | sed 's/[[:space:]]*$//')
AUTHOR=$(echo "$RAWAUTHOR" | sed 's/[[:space:]]*$//')
USERNAME=$(echo "$RAWUSERNAME" | sed 's/[[:space:]]*$//')

if [[ "$POST" == "228" ]] && [[ "$USERNAME" == "aboba" ]] && [[ "$AUTHOR" == "test" ]]; then
    echo "Test passed"
else
    echo "Test failed"
    docker-compose -f service_test_with_kafka.yml down > /dev/null 2>&1
    exit 1
fi

docker-compose -f service_test_with_kafka.yml down > /dev/null 2>&1
