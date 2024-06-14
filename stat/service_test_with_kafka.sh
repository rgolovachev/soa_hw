#!/bin/bash

docker-compose -f service_test_with_kafka.yml build > /dev/null 2>&1
echo "build done"
docker-compose -f service_test_with_kafka.yml up -d > /dev/null 2>&1
echo "up done"

sleep 35

docker exec -it clickhouse clickhouse-client --query="TRUNCATE TABLE likes;" > /dev/null 2>&1
echo '{"post_id": "228", "username": "aboba"}' | docker exec -i kafka kafka-console-producer.sh --broker-list 0.0.0.0:9092 --topic likes
OUTPUT=$(docker exec -it clickhouse clickhouse-client --query="SELECT * FROM likes" | tail -n +3)

trimmed_string=$(echo "$OUTPUT" | sed 's/[[:space:]]*$//')

if [ "$OUTPUT" = "228     test    aboba" ]; then
    echo "Test passed"
else
    echo "Test failed"
    echo "$OUTPUT"
    echo "228     test    aboba"
    docker-compose -f service_test_with_kafka.yml down > /dev/null 2>&1
    exit 1
fi

docker-compose -f service_test_with_kafka.yml down > /dev/null 2>&1
