#!/bin/bash

docker-compose -f service_test.yml build
echo "build done"
docker-compose -f service_test.yml up -d
echo "up done"

sleep 40

docker exec -it clickhouse clickhouse-client --query="INSERT INTO likes (*) VALUES ('some_uid0', 'vasya', 'petya')" > /dev/null 2>&1
docker exec -it clickhouse clickhouse-client --query="INSERT INTO likes (*) VALUES ('some_uid0', 'vasya', 'petya2')" > /dev/null 2>&1
docker exec -it clickhouse clickhouse-client --query="INSERT INTO likes (*) VALUES ('some_uid0', 'vasya', 'petya3')" > /dev/null 2>&1
docker exec -it clickhouse clickhouse-client --query="INSERT INTO views (*) VALUES ('some_uid0', 'vasya', 'petya3')" > /dev/null 2>&1

output1=$(grpcurl -d '{"post_id": "some_uid0"}' -plaintext 0.0.0.0:11888 stat.StatService/GetPostStats)
if [ $? -ne 0 ]; then
  docker-compose -f service_test.yml down > /dev/null 2>&1
  exit 1
fi

LIKES=$(echo $output1 | jq -r '.likes')
VIEWS=$(echo $output1 | jq -r '.views')

if [ $LIKES -ne "3" ] || [ $VIEWS -ne "1" ]; then
  echo "Test failed"
  docker-compose -f service_test.yml down > /dev/null 2>&1
  exit 1
else
  echo "Test passed"
fi

docker-compose -f service_test.yml down > /dev/null 2>&1
